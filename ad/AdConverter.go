package ad

func ConvertAdReqVoToAd(vo CreateAdReqVo) Ad {
	return Ad{
		Title:   vo.Title,
		StartAt: vo.StartAt,
		EndAt:   vo.EndAt,
	}
}

func ConvertAdReqVoToAdConditions(vo CreateAdReqVo, adID int) []AdCondition {
	var adConditions []AdCondition
	for _, condition := range vo.Conditions {
		adCondition := AdCondition{
			AgeStart: condition.AgeStart,
			AgeEnd:   condition.AgeEnd,
			Gender:   convertGenderToStringSlice(condition.Gender),
			Country:  condition.Country,
			Platform: convertPlatformToStringSlice(condition.Platform),
			AdID:     adID, // 根据实际情况填充
		}
		adConditions = append(adConditions, adCondition)
	}
	return adConditions
}

func convertGenderToStringSlice(genders []Gender) []string {
	stringSlice := make([]string, len(genders))
	for i, gender := range genders {
		stringSlice[i] = string(gender)
	}
	return stringSlice
}

func convertPlatformToStringSlice(platforms []Platform) []string {
	stringSlice := make([]string, len(platforms))
	for i, gender := range platforms {
		stringSlice[i] = string(gender)
	}
	return stringSlice
}
