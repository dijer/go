package config

import (
	"flag"
	"fmt"
	"net/url"
)

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
	port := flag.String("port", "8080", "port to connect to db")
	db_url := flag.String("host", "postgres://db-user:db-password@petstore-db:5432/petstore?sslmode=disable", "url to connect to db")
	jaeger_url := flag.String("jaeger_url", "http://jaeger:16686", "jaeger")
	sentry_url := flag.String("sentry_url", "http://sentry:9000", "host to connect to db")
	kafka_broker := flag.String("kafka_broker", "kafka:9092", "kafka broker")
	some_app_id := flag.String("some_app_id", "testid", "app id")
	some_app_key := flag.String("some_app_key", "testkey", "app key")

	flag.Parse()

	validate(*port, *db_url, *jaeger_url, *sentry_url, *kafka_broker, *some_app_id, *some_app_key)

	fmt.Println(*port, *db_url, *jaeger_url, *sentry_url, *kafka_broker, *some_app_id, *some_app_key)
}
