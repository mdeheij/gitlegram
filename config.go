package main

import (
	"encoding/json"
	"errors"
	"log"
	"os"
	"os/signal"
	"syscall"
)

//Config struct represents the config filé
type Config struct {
	Bottoken     string
	Address      string
	Port         int64
	Repositories []ConfigRepository
}

func NewConfig() *Config {
	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc, syscall.SIGHUP)
	go func() {
		<-sigc
		config = loadConfig("config.json")
		log.Println("config reloaded")
	}()

	config = loadConfig("config.json")
	return &config
}

//ConfigRepository represents a repository from the config file
type ConfigRepository struct {
	Clone_url      string
	Telegramtarget string
	Commands       []string
}

func PanicIf(err error, what ...string) {
	if err != nil {
		if len(what) == 0 {
			panic(err)
		}

		panic(errors.New(err.Error() + what[0]))
	}
}

var config Config
var configFile string

//TODO: toml instead of json
func loadConfig(configFile string) (c Config) {
	var file, err = os.Open(configFile)
	PanicIf(err)

	defer func() {
		err := file.Close()
		PanicIf(err)
	}()

	buffer := make([]byte, 1024)
	count := 0

	count, err = file.Read(buffer)
	PanicIf(err)

	err = json.Unmarshal(buffer[:count], &c)
	PanicIf(err)

	return c
}
