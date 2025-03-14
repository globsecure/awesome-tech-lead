package main

import (
	"fmt"
	"net/http"
	"os"
	"regexp"
	"time"

	catalog "github.com/tech-leads-club/awesome-tech-lead/internal"
)

var (
	allowedTypes       = []string{"article", "video", "course", "book", "podcast", "feed"}
	allowedLevels      = []string{"beginner", "intermediate", "advanced"}
	allowedLanguages   = []string{"en", "pt_br", "es", "en_us"}
	allowedCareerBands = []string{"mid", "senior", "principal", "staff", "tl", "junior"}
	urlRegex           = regexp.MustCompile(`^(http|https)://[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}(/\S*)?$`)
)

func main() {
	data, err := os.ReadFile("catalog.yml")
	if err != nil {
		fmt.Println("Error reading catalog.yml:", err)
		os.Exit(1)
	}

	items, err := catalog.ParseCatalog(data)
	if err != nil {
		fmt.Println("Error parsing catalog.yml:", err)
		os.Exit(1)
	}

	for _, item := range items {
		if err := validateItem(item); err != nil {
			fmt.Println("Validation error:", err)
			os.Exit(1)
		}
	}

	fmt.Println("catalog.yml validated successfully!")
}

func validateItem(item catalog.CatalogItem) error {
	if item.URL == "" {
		return fmt.Errorf("URL cannot be empty: %v", item)
	}
	if !urlRegex.MatchString(item.URL) {
		return fmt.Errorf("Invalid URL: %s", item.URL)
	}
	// if err := validateLink(item.URL); err != nil {
	// 	return fmt.Errorf("Broken link: %s (%v)", item.URL, err)
	// }
	if item.Title == "" {
		return fmt.Errorf("Title cannot be empty: %v", item)
	}
	if !contains(allowedTypes, item.Type) {
		return fmt.Errorf("Invalid type: %s", item.Type)
	}
	if !contains(allowedLevels, item.Level) {
		return fmt.Errorf("Invalid level: %s", item.Level)
	}
	if !contains(allowedLanguages, item.Language) {
		return fmt.Errorf("Invalid language: %s", item.Language)
	}
	for _, band := range item.CareerBands {
		if !contains(allowedCareerBands, band) {
			return fmt.Errorf("Invalid career band: %s", band)
		}
	}

	// Add more validations as needed

	return nil
}

func validateLink(url string) error {
	client := http.Client{Timeout: 10 * time.Second}
	resp, err := client.Head(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("status code %d", resp.StatusCode)
	}
	return nil
}

func contains(list []string, item string) bool {
	for _, value := range list {
		if value == item {
			return true
		}
	}
	return false
}
