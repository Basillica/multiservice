package appenv

type AppConfig struct {
	AllowedOrigin       string `mapstructure:"ALLOWED_ORIGIN"`
	CookieDomain        string `mapstructure:"COOKIE_DOMAIN"`
	CookieSecureEnabled bool   `mapstructure:"COOKIE_SECURE_ENABLE"`
	CookieHTTPOnly      bool   `mapstructure:"COOKIE_HTTPONLY"`
}
