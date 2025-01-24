package dto

type TranslateRequest struct {
	Text         string `json:"q"`
	Source       string `json:"source"`
	Target       string `json:"target"`
	Format       string `json:"format"`
	Alternatives int    `json:"alternatives"`
	ApiKey       string `json:"api_key"`
}
