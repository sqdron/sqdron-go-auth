package aws

type FacebookConfig struct {
	ClientID     string `json:"key"`
	ClientSecret string `json:"secret"`
	Bucket       string `json:"bucket"`
}
