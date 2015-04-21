package config

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
)

var configFile = flag.String("config", "config.json", "Address of config file")

func Init(v interface{}) {
	flag.Parse()

	file, err := ioutil.ReadFile(*configFile)

	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(file, v)
	if err != nil {
		log.Fatal(err)
	}
}
