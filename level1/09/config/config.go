package config

import (
	"encoding/json"
	"fmt"
	"net/url"
	"os"
)

type Config struct {
	Port         string `json:"port"`
	Db_url       string `json:"db_url"`
	Jaeger_url   string `json:"jaeger_url"`
	Sentry_url   string `json:"sentry_url"`
	Kafka_broker string `json:"kafka_broker"`
	Some_app_id  string `json:"some_app_id"`
	Some_app_key string `json:"some_app_key"`
}

func validate(port, db_url, jaeger_url, sentry_url, kafka_broker, some_app_id, some_app_key string) {
	urls := []string{db_url, jaeger_url, sentry_url}

	for _, checkUrl := range urls {
		_, err := url.ParseRequestURI(checkUrl)
		if err != nil {
			panic(err)
		}
		break
	}

	strings := []string{port, kafka_broker, some_app_id, some_app_key}

	for i, checkString := range strings {
		if len(checkString) == 0 {
			panic(i)
		}
		break
	}
}

func ConnectToDb() {
	data, err := os.Open("data.json")
	if err != nil {
		panic(err)
	}

	defer data.Close()

	var cnfg Config
	if err = json.NewDecoder(data).Decode(&cnfg); err != nil {
		panic(err)
	}

	fmt.Println(cnfg)

	validate(cnfg.Port, cnfg.Db_url, cnfg.Jaeger_url, cnfg.Sentry_url, cnfg.Kafka_broker, cnfg.Some_app_id, cnfg.Some_app_key)

	fmt.Println(cnfg.Port, cnfg.Db_url, cnfg.Jaeger_url, cnfg.Sentry_url, cnfg.Kafka_broker, cnfg.Some_app_id, cnfg.Some_app_key)
}
