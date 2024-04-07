package ad

import (
	"context"
	"dcard/infra"
	"encoding/json"
	"errors"
	"github.com/go-redis/redis/v8"
	"log"
	"math/rand"
	"strconv"
	"time"
)

func createAd(vo CreateAdReqVo) (int, string, int) {
	if vo.EndAt.Before(vo.StartAt) {
		return 400, "EndAt should not be earlier than StartAt", 0
	}

	err2 := checkConditionEnum(vo.Conditions)
	if err2 != nil {
		return 400, err2.Error(), 0
	}

	db, err := infra.GetDB()
	if err != nil {
		return 500, "Failed to connect to database:" + err.Error() + err.Error(), 0
	}

	tx, err := db.Begin()
	if err != nil {
		return 500, "Failed to begin transaction:" + err.Error(), 0
	}

	defer func() {
		if err != nil {
			err := tx.Rollback()
			if err != nil {
				return
			}
			log.Fatalln(err, "Transaction rollback:")
		} else {
			err = tx.Commit()
			if err != nil {
				log.Fatalln(err, "Transaction commit failed:")
			}
		}
	}()

	ad := ConvertAdReqVoToAd(vo)
	adID := addAd(tx, ad)

	adConditions := ConvertAdReqVoToAdConditions(vo, adID)
	err = addAdConditions(tx, adConditions)
	if err != nil {
		return 500, "Failure to insert adCondtions:" + err.Error(), 0
	}

	return 0, "", adID
}

func checkConditionEnum(conditions []Condition) error {
	for _, condition := range conditions {
		for _, gender := range condition.Gender {
			if gender != Male && gender != Female {
				return errors.New("value of Gender does not exist")
			}
		}

		for _, platform := range condition.Platform {
			if platform != Android && platform != IOS && platform != Web {
				return errors.New("value of Platform does not exist")
			}
		}

		for _, country := range condition.Country {
			if !CountryMap[country] {
				return errors.New("value of Country does not exist")
			}
		}

	}
	return nil
}

func getAd(vo GetAdReqVo) (int, string, []map[string]interface{}) {
	hash := strconv.Itoa(vo.Age) +
		vo.Gender +
		vo.Country +
		vo.Platform +
		strconv.Itoa(vo.Offset) +
		strconv.Itoa(vo.Limit)

	localCacheResult, ok := adCache.Get(hash)

	if ok {
		return 0, "", localCacheResult.([]map[string]interface{})
	}

	client := infra.GetRedisClient()
	ctx := context.Background()
	adCacheResultJson, err2 := client.HGet(ctx, "ad", hash).Result()
	lockKey := hash
	if err2 != nil {
		if err2 == redis.Nil {
			//lockKey := hash + "_" + uuid.New().String() // 防止該執行緒TTL到期釋放鎖後，自己又再一次釋放鎖，影響到其它執行緒
			isGetLocked, err := client.SetNX(ctx, lockKey, "locked", 200*time.Millisecond).Result()

			if err != nil {
				return 500, "Failed to lock:" + err.Error(), nil
			}
			if isGetLocked {
				log.Println("Get Lock:" + lockKey)

				db, err := infra.GetDB()
				if err != nil {
					return 500, "Failed to connect to database:" + err.Error(), nil
				}

				ads, err := retrieveAd(vo, db)
				if err != nil {
					return 500, "Error querying database:" + err.Error(), nil
				}

				cacheValue := AdCacheDto{
					Ads: ads,
				}

				cacheValueJson, err := json.Marshal(cacheValue)
				if err != nil {
					return 500, "Error encoding user:" + err.Error(), nil
				}

				client.HSet(ctx, "ad", hash, cacheValueJson)

				_, err = client.Del(ctx, lockKey).Result()

				adCache.Set(hash, ads)

				return 0, "", ads
			} else {
				retry := 1
				retryDelay := 100 * time.Millisecond

				for retry <= 3 {
					time.Sleep(retryDelay + randomDuration(100*time.Millisecond))
					exists, err := client.Exists(ctx, hash).Result()
					if err != nil {
						return 500, "Failed to lock:" + err.Error(), nil
					}

					if exists == 1 {
						retry++
						retryDelay *= 2
						continue
					} else {
						return getAd(vo)
					}
				}
			}
		} else {
			return 500, "Failed to obtain redis data:" + err2.Error(), nil
		}
	}

	log.Println("Thread Read Cache:" + lockKey)
	var cacheValue AdCacheDto
	err := json.Unmarshal([]byte(adCacheResultJson), &cacheValue)
	if err != nil {
		return 500, "Error encoding user:" + err.Error(), nil
	}
	adCache.Set(hash, cacheValue.Ads)

	return 0, "", cacheValue.Ads
}

func randomDuration(max time.Duration) time.Duration {
	return time.Duration(rand.Intn(int(max))) * time.Millisecond
}
