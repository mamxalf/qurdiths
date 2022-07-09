package config

import (
	"github.com/joho/godotenv"
	"hadithgo/app/exceptions"
	"os"
)

type Config interface {
	Get(key string) string
}

type configImpl struct {
}

func (config *configImpl) Get(key string) string {
	return os.Getenv(key)
}

func New(filenames ...string) Config {
	err := godotenv.Load(filenames...)
	exceptions.PanicIfNeeded(err)
	return &configImpl{}
}
