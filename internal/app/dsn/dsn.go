package dsn

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"os"
)

func FromEnv(logger *logrus.Logger) string {
	if err := godotenv.Load(); err != nil {
		logger.Fatalln(err)
	}

	host, existHost := os.LookupEnv("DB_HOST")
	port, existPort := os.LookupEnv("DB_PORT")
	user, existUser := os.LookupEnv("DB_USER")
	pass, existPass := os.LookupEnv("DB_PASS")
	dbname, existName := os.LookupEnv("DB_NAME")
	if !existHost || !existPort || !existUser || !existPass || !existName {
		return ""
	}

	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, pass, dbname)
}
