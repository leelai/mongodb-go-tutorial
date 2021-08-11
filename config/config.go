package config

import (
	"log"
	"os"
	"strings"

	"example.com/hello-mongodb/db"
	"github.com/caarlos0/env/v6"
)

type config struct {
	MongoUrl string `env:"MONGO_URL" envDefault:"mongodb+srv://leelai:james67210@cluster0.ch5ms.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"`
}

func (conf config) Validate() bool {

	if len(conf.MongoUrl) == 0 {
		log.Fatalf("ERR: mongo url is empty!")
		return false
	}

	if url := strings.ToLower(conf.MongoUrl); !strings.HasPrefix(url, "mongodb+srv://") {
		log.Fatalf("ERR: MONGO_URL format is not correct!")
		return false
	}

	return true
}

var Config config

func Initial() {

	// parse config
	if err := env.Parse(&Config); err != nil {
		log.Fatalf("ERR: %v\n", err)
	}

	// validate config
	if !Config.Validate() {
		os.Exit(1)
	}

	db.Dbconnect(Config.MongoUrl)
}
