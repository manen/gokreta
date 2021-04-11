package kreta

type ListSchool struct {
	KretaID      int    `json:"instituteId"`
	ID           string `json:"instituteCode"`
	Name         string `json:"name"`
	City         string `json:"city"`
	URL          string `json:"url"`
	AdURL        string `json:"advertisingUrl"`
	InfoImageURL string `json:"informationImageUrl"`
	InfoURL      string `json:"informationUrl"`
	// Features <any> `json:"featureToggleSet"`
}
