package example

import (
	"fmt"

	"github.com/zhtfirst/go-packages/config"
)

func GetConfig() {
	config.Setup("")

	brokerUrl := config.GetString("worker", "broker_url")
	fmt.Println(brokerUrl)

	concurrency := config.GetInt64("worker", "concurrency")
	fmt.Println(concurrency)

	boole := config.GetBoole("worker", "boole")
	fmt.Println(boole)

	float := config.GetFloat("worker", "float")
	fmt.Println(float)
}
