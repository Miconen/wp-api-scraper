package main

import (
	"os"
	"sync"
	"wp-api-scraper/models"
	"wp-api-scraper/utils"
)

var cfg models.Config

func main() {
	if _, err := os.Stat("./results/"); os.IsNotExist(err) {
		os.Mkdir("./results/", 0700)
	}

	cfg = utils.ConfigParser()

	var wg sync.WaitGroup

	channel := make(chan models.Post)

	wg.Add(1)
	go utils.PostProducer(channel, &wg, &cfg)

	wg.Add(1)
	go utils.PostConsumer(channel, &wg, &cfg)

	wg.Wait()

	close(channel)
}
