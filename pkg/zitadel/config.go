package zitadel

type Config struct {
	URL       string   `mapstructure:"ZITADEL_URL"`
	Issuer    string   `mapstructure:"ZITADEL_ISSUER"`
	Scopes    []string `mapstructure:"ZITADEL_SCOPES"`
	JWTToken  string   `mapstructure:"ZITADEL_JWT_KEY"`
	ProjectID string   `mapstructure:"ZITADEL_PROJECT_ID"`
}
