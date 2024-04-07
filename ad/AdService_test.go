package ad

import (
	"testing"
	"time"
)

func TestCreateAd_ValidParameters(t *testing.T) {
	vo := CreateAdReqVo{
		Title:   "Test Ad",
		StartAt: time.Now(),
		EndAt:   time.Now().Add(time.Hour * 24),
		Conditions: []Condition{
			{
				AgeStart: 18,
				AgeEnd:   60,
				Gender:   []Gender{Male, Female},
				Country:  []string{"TW"},
				Platform: []Platform{Android, IOS},
			},
		},
	}

	code, errMsg, adID := createAd(vo)

	if code != 0 {
		t.Errorf("Expected code to be 0, got %d", code)
	}
	if errMsg != "" {
		t.Errorf("Expected errMsg to be empty, got %s", errMsg)
	}
	if adID == 0 {
		t.Errorf("Expected adID to be non-zero, got %d", adID)
	}
}

func TestCreateAd_GenderNotExists(t *testing.T) {
	vo := CreateAdReqVo{
		Title:   "Test Ad",
		StartAt: time.Now(),
		EndAt:   time.Now().Add(time.Hour * 24),
		Conditions: []Condition{
			{
				AgeStart: 18,
				AgeEnd:   60,
				Gender:   []Gender{"S"},
				Country:  []string{"TW"},
				Platform: []Platform{Android, IOS},
			},
		},
	}

	code, errMsg, _ := createAd(vo)

	if code != 400 {
		t.Errorf("Expected code to be 400, got %d", code)
	}
	if errMsg != "value of Gender does not exist" {
		t.Errorf("Expected errMsg to be value of Gender does not exist, got %s", errMsg)
	}
}

func TestCreateAd_CountryNotExists(t *testing.T) {
	vo := CreateAdReqVo{
		Title:   "Test Ad",
		StartAt: time.Now(),
		EndAt:   time.Now().Add(time.Hour * 24),
		Conditions: []Condition{
			{
				AgeStart: 18,
				AgeEnd:   60,
				Gender:   []Gender{Female, Male},
				Country:  []string{"TWS"},
				Platform: []Platform{Android, IOS},
			},
		},
	}

	code, errMsg, _ := createAd(vo)

	if code != 400 {
		t.Errorf("Expected code to be 400, got %d", code)
	}
	if errMsg != "value of Country does not exist" {
		t.Errorf("Expected errMsg to be value of Country does not exist, got %s", errMsg)
	}
}

func TestCreateAd_PlatformNotExists(t *testing.T) {
	vo := CreateAdReqVo{
		Title:   "Test Ad",
		StartAt: time.Now(),
		EndAt:   time.Now().Add(time.Hour * 24),
		Conditions: []Condition{
			{
				AgeStart: 18,
				AgeEnd:   60,
				Gender:   []Gender{Female, Male},
				Country:  []string{"TW"},
				Platform: []Platform{"SCREEN"},
			},
		},
	}

	code, errMsg, _ := createAd(vo)

	if code != 400 {
		t.Errorf("Expected code to be 400, got %d", code)
	}
	if errMsg != "value of Platform does not exist" {
		t.Errorf("Expected errMsg to be value of Platform does not exist, got %s", errMsg)
	}
}

func TestCreateAd_EndAtShouldNotBeEarlierThanStartAt(t *testing.T) {
	vo := CreateAdReqVo{
		Title:   "Test Ad",
		StartAt: time.Now().Add(time.Hour * 24),
		EndAt:   time.Now(),
		Conditions: []Condition{
			{
				AgeStart: 18,
				AgeEnd:   60,
				Gender:   []Gender{Male, Female},
				Country:  []string{"TW"},
				Platform: []Platform{Android, IOS},
			},
		},
	}

	code, errMsg, _ := createAd(vo)

	if code != 400 {
		t.Errorf("Expected code to be 400, got %d", code)
	}
	if errMsg != "EndAt should not be earlier than StartAt" {
		t.Errorf("Expected errMsg to be EndAt should not be earlier than StartAt, got %s", errMsg)
	}

}

func TestGetAd_ValidParameters(t *testing.T) {
	vo := GetAdReqVo{
		Age:      30,
		Gender:   "M",
		Country:  "TW",
		Platform: "web",
		Offset:   0,
		Limit:    10,
	}

	code, errMsg, _ := getAd(vo)

	if code != 0 {
		t.Errorf("Expected code to be 200, got %d", code)
	}
	if errMsg != "" {
		t.Errorf("Expected errMsg to be empty, got %s", errMsg)
	}
}
