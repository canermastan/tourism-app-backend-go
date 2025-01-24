package dto

/*
	{
	  "q": "...",
	  "source": "auto",
	  "target": "en",
	  "format": "text",
	  "alternatives": 3,
	  "api_key": ""
	}
*/
type TranslateRequest struct {
	Text         string `json:"q"`
	Source       string `json:"source"`
	Target       string `json:"target"`
	Format       string `json:"format"`
	Alternatives int    `json:"alternatives"`
	ApiKey       string `json:"api_key"`
}
