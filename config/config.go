package config

import (
	"log"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	values map[string]string
}

func InitConfig() (*Config, error) {

	var c Config
	var err error
	c.values, err = godotenv.Read(".env")
	if err != nil {
		return nil, err
	}

	return &c, nil
}

func (c *Config) Port() string {
	return c.values["PORT"]
}

func (c *Config) Difficulty() int {

	difficulty, err := strconv.Atoi(c.values["DIFFICULTY"])
	if err != nil {
		log.Fatal(err)
	}

	return difficulty
}
