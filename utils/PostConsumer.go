package utils

import (
	"fmt"
	"os"
	"strings"
	"sync"
	"wp-api-scraper/models"
)

func PostConsumer(channel <-chan models.Post, wg *sync.WaitGroup, cfg *models.Config) {
	m := make(map[int]string)

	defer wg.Done()
	var res models.Posts

	for data := range channel {
		res = append(res, data)
		for _, post := range res {
			content := strings.ToLower(post.Content.Rendered)

			for _, keyword := range cfg.Keywords {
				if !strings.Contains(content, keyword) {
					continue
				}

				// Check if the post has already been saved
				if _, ok := m[post.ID]; ok {
					continue
				}

				filename := strings.TrimSpace(keyword)
				filename = strings.ReplaceAll(filename, " ", "-")
				err := savePost(filename+".txt", post.Link)
				if err != nil {
					fmt.Println("Consumer: Error saving post")
				}
				fmt.Println("Found: ", keyword)
			}

			// Save post link under all keywords before setting
			m[post.ID] = post.Link
		}
	}
}

func savePost(filename string, link string) error {
	path := "./results/" + filename
	file, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	// Write data to the file
	_, err = fmt.Fprintf(file, link+"\n")
	return err
}
