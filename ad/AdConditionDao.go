package ad

import (
	"database/sql"
	"strings"
)

func addAdConditions(tx *sql.Tx, adConditions []AdCondition) error {
	// 准备 SQL 插入语句
	stmt, err := tx.Prepare("INSERT INTO ad_condition (age_start, age_end, gender, country, platform, ad_id) VALUES ($1, $2, $3, $4, $5, $6)")
	if err != nil {
		return err
	}

	// 逐个插入 AdCondition 数据
	for _, adCondition := range adConditions {
		// 将数组转换为逗号分隔的字符串
		genderStr := "{" + strings.Join(adCondition.Gender, ",") + "}"
		countryStr := "{" + strings.Join(adCondition.Country, ",") + "}"
		platformStr := "{" + strings.Join(adCondition.Platform, ",") + "}"

		// 执行 SQL 插入语句
		_, err := stmt.Exec(adCondition.AgeStart, adCondition.AgeEnd, genderStr, countryStr, platformStr, adCondition.AdID)
		if err != nil {
			return err
		}
	}

	return nil
}

func convertToInterfaceSlice(slice []interface{}) []interface{} {
	newSlice := make([]interface{}, len(slice))
	for i, v := range slice {
		newSlice[i] = v
	}
	return newSlice
}
