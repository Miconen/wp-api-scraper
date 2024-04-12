# WordPress API Scraper

Designed to scrape a sites public WordPress API for a list of various infromation. I personally only needed it to scrape posts belonging to certain categories, but it can be easily modified to scrape other information.

## Installation

You will need Go installed and then just run

```bash
go run main.go
```

You will also need to add a config.yml file and modify it to fit your needs.

```yml
target:
  url: "https://www.example.com"
  path: "wp-json/wp/v2/posts"
  categories: category_number # Currently does not support optional url parameters
  per_page: your_number_here # Max is 100
  cooldown: 5 # Time in seconds to wait between requests
keywords: ["list of strings to save into files"]
```

## Disclaimer

Do not be weird and use this for malicious purposes. This is a tool I made for myself to help scrape information responsibly, do not abuse it. I am not responsible for any misuse of this tool.
