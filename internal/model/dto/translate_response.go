package dto

type TranslateResponse struct {
	Alternatives     []interface{} `json:"alternatives"`
	DetectedLanguage struct {
		Confidence float64 `json:"confidence"`
		Language   string  `json:"language"`
	} `json:"detectedLanguage"`
	TranslatedText string `json:"translatedText"`
}
