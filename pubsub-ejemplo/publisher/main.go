package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	dapr "github.com/dapr/go-sdk/client"
)

const (
	pubsubComponentName = "votacionpubsub"
	pubsubTopic         = "votacion"
)

var (
	r = rand.New(rand.NewSource(time.Now().UnixNano()))
)

func main() {

	client, err := dapr.NewClient()
	if err != nil {
		panic(err)
	}
	defer client.Close()

	for true {

		order := `{"votacion":` + votacionRandom() + `}`

		err := client.PublishEvent(context.Background(), pubsubComponentName, pubsubTopic, []byte(order))
		if err != nil {
			panic(err)
		}

		fmt.Println("publicacion:", order)

		time.Sleep(5 * time.Second)
	}

}

func votacionRandom() string {

	if r.Intn(2) == 0 {
		return "si sale el curso"
	}
	return "no sale el curso"
}
