package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/MrBrooks89/BatSecrets/pkg/secretfinder"
	"github.com/schollz/progressbar/v3"
)

func main() {
	urlFlag := flag.String("u", "", "URL to check for secrets")
	listFlag := flag.String("l", "", "File containing a list of URLs to check for secrets")
	concurrencyFlag := flag.Int("c", 100, "Limit the concurrency of goroutines")
	verboseFlag := flag.Bool("v", false, "Print verbose output")
	flag.Parse()

	var urls []string
	if *urlFlag != "" {
		urls = append(urls, *urlFlag)
	} else if *listFlag != "" {
		urlsFromFile := secretfinder.ReadURLsFromFile(*listFlag)
		urls = append(urls, urlsFromFile...)
	} else {
		log.Println("Please specify either a single URL (-u) or a file containing a list of URLs (-l)")
		return
	}

	secrets := secretfinder.GetSecretRegexes()

	bar := progressbar.Default(int64(len(urls)))

	allMatches := secretfinder.CheckURLs(urls, secrets, *concurrencyFlag, *verboseFlag, bar)

	// Print all unique matches at the end
	seen := make(map[string]bool)
	for _, match := range allMatches {
		key := fmt.Sprintf("%s:%s", match.Secret.Name, match.Match)
		if !seen[key] {
			seen[key] = true
			fmt.Printf("Secret: %s\nDescription: %s\nURL: %s\nMatch: %s\n\n", match.Secret.Name, match.Secret.Description, match.URL, match.Match)
		}
	}
}
