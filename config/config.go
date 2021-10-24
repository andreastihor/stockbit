package config

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

//Conf global environtment
var Env *GeneralConfig = nil

//GeneralConfig ...
type GeneralConfig struct {
	DSN string
}

//init Get env from .env file
func init() {
	logger := log.New(os.Stdout, "Getting Configuration From Environment", log.LstdFlags)

	if Env == nil {
		if err := godotenv.Load(); err != nil {
			logger.Print("No .env file found")
		}
		host := os.Getenv("DB_HOST")
		if host == "" {
			logger.Print("no DB_HOST in environment")
		}

		port := os.Getenv("DB_PORT")
		if port == "" {
			logger.Print("no DB_PORT in environment")
		}

		dbname := os.Getenv("DB_NAME")
		if dbname == "" {
			logger.Print("no DB_NAME in environment")

		}
		usname := os.Getenv("DB_USERNAME")
		if usname == "" {
			logger.Print("no DB_USERNAME in environment")

		}
		pwd := os.Getenv("DB_PASSWORD")
		if pwd == "" {
			logger.Print("no DB_PASSWORD in environment")

		}

		mode := os.Getenv("SSL_MODE")
		if mode == "" {
			logger.Print("no SSL_MODE in environment")
		}

		ssl_mode := "disable"
		if strings.ToUpper(mode) == "TRUE" {
			ssl_mode = "enable"
		}
		Env = &GeneralConfig{
			DSN: GenerateDSN(host, usname, pwd, port, dbname, ssl_mode),
		}
	}

}

//GenerateDSN generate connection string
func GenerateDSN(host string, user string, password string, port string, dbname string, ssl_mode string) string {
	constr := fmt.Sprintf("host=%s user=%s password=%s port=%s  dbname=%s sslmode=%s",
		host,
		user,
		password,
		port,
		dbname,
		ssl_mode)

	return constr
}
