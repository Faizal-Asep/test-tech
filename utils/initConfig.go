package util

import (
	"log"
	"time"

	"github.com/spf13/viper"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var (
	Config      *ConfigStruct
	OauthConfGl = &oauth2.Config{
		ClientID:     "",
		ClientSecret: "",
		RedirectURL:  "",
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
		Endpoint:     google.Endpoint,
	}
	// oauthStateStringGl = ""
)

func init() {
	viper.AddConfigPath(".")
	viper.SetConfigType("env")
	viper.SetConfigName("app")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}

	err = viper.Unmarshal(&Config)
	if err != nil {
		log.Fatal(err)
	}

	OauthConfGl.ClientID = Config.GoogleClientID
	OauthConfGl.ClientSecret = Config.GoogleClientSecret
	OauthConfGl.RedirectURL = Config.GoogleOAuthRedirectUrl
}

type ConfigStruct struct {
	DbHost         string `mapstructure:"DB_HOST"`
	DbUser         string `mapstructure:"DB_USER"`
	DbPassword     string `mapstructure:"DB_PASSWORD"`
	DbDatabase     string `mapstructure:"DB_DATABASE"`
	EncriptKey     string `mapstructure:"ENCRIPT_KEY"`
	EncriptIv      string `mapstructure:"ENCRIPT_IV"`
	FrontEndOrigin string `mapstructure:"FRONTEND_ORIGIN"`

	JWTTokenSecret string        `mapstructure:"JWT_SECRET"`
	TokenExpiresIn time.Duration `mapstructure:"TOKEN_EXPIRED_IN"`
	TokenMaxAge    int           `mapstructure:"TOKEN_MAXAGE"`

	GoogleClientID         string `mapstructure:"GOOGLE_OAUTH_CLIENT_ID"`
	GoogleClientSecret     string `mapstructure:"GOOGLE_OAUTH_CLIENT_SECRET"`
	GoogleOAuthRedirectUrl string `mapstructure:"GOOGLE_OAUTH_REDIRECT_URL"`
}
