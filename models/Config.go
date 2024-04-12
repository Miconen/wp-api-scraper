package models

type Config struct {
	Target struct {
		URL        string `yaml:"url"`
		Path       string `yaml:"path"`
		Categories string `yaml:"categories"`
		PerPage    string `yaml:"per_page"`
	} `yaml:"target"`
	Keywords []string `yaml:"keywords"`
}
