package config

import (
	"fmt"
	"os"
	"strconv"
)

type Config struct {
	PORT          uint
	POSTGRES_USER string
	POSTGRES_PW   string
	POSTGRES_DB   string
	POSTGRES_PORT uint
	POSTGRES_HOST string
	SECRET        string
	GO_ENV        string
	API_DOMAIN    string
	FE_URL        string
}

var Conf *Config

func Init() error {
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		return fmt.Errorf("failed to read PORT")
	}

	postgresPort, err := strconv.Atoi(os.Getenv("POSTGRES_PORT"))
	if err != nil {
		return fmt.Errorf("failed to read POSTGRES_PORT")
	}

	Conf = &Config{
		PORT:          uint(port),
		POSTGRES_USER: os.Getenv("POSTGRES_USER"),
		POSTGRES_PW:   os.Getenv("POSTGRES_PW"),
		POSTGRES_DB:   os.Getenv("POSTGRES_DB"),
		POSTGRES_PORT: uint(postgresPort),
		POSTGRES_HOST: os.Getenv("POSTGRES_HOST"),
		SECRET:        os.Getenv("SECRET"),
		GO_ENV:        os.Getenv("GO_ENV"),
		API_DOMAIN:    os.Getenv("API_DOMAIN"),
		FE_URL:        os.Getenv("FE_URL"),
	}

	return nil
}
