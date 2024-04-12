package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"sync"
	"time"
	"wp-api-scraper/models"
)

func PostProducer(channel chan<- models.Post, wg *sync.WaitGroup, cfg *models.Config) {
	defer wg.Done()

	baseUrl, err := url.Parse(cfg.Target.URL)
	if err != nil {
		fmt.Println("Malformed URL: ", err.Error())
		return
	}

	baseUrl.Path += cfg.Target.Path

	params := url.Values{}
	params.Add("categories", cfg.Target.Categories)
	params.Add("per_page", cfg.Target.PerPage)

	page := 1
	for {
		url := baseUrl
		params.Set("page", strconv.Itoa(page))
		url.RawQuery = params.Encode()

		res, err := fetchUrl(url.String())
		if err != nil {
			fmt.Println("Error getting page: ", err.Error())
			return
		}

		for _, post := range res {
			channel <- post
		}

		page++
		if page <= 10 {
			time.Sleep(5 * time.Second)
		}
	}
}

func fetchUrl(url string) (models.Posts, error) {
	res, err := http.Get(url)
	if err != nil {
		fmt.Printf("Failed HTTP GET request to URL: %s", url)
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Failed to read HTTP response body")
		return nil, err
	}

	var result models.Posts
	if err := json.Unmarshal(body, &result); err != nil {
		fmt.Println("Can not unmarshal JSON")
		return nil, err
	}

	return result, nil
}
