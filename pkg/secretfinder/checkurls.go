package secretfinder

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
	"sync"

	"github.com/schollz/progressbar/v3"
)

func CheckURLs(urls []string, secrets []Secret, concurrency int, verbose bool, bar *progressbar.ProgressBar) []SecretMatched {
	var allMatches []SecretMatched
	matchesCh := make(chan []SecretMatched)
	var wg sync.WaitGroup
	semaphore := make(chan struct{}, concurrency)

	for _, url := range urls {
		wg.Add(1)
		go func(u string) {
			defer func() { <-semaphore }()
			defer wg.Done()

			semaphore <- struct{}{}

			jsCode, err := fetchJavaScriptCode(u, &http.Client{}, verbose)

			if err != nil {
				if verbose {
					log.Printf("Error fetching URL %s: %v\n", u, err)
				}
				return
			}

			// Only proceed to check for secrets if fetch operation was successful
			success := jsCode != ""
			if !success {
				log.Printf("Failed to fetch JavaScript code for URL %s", u)
				return
			}

			var matches []SecretMatched
			for _, secret := range secrets {
				re := regexp.MustCompile(secret.Regex)
				match := re.FindString(jsCode)
				if match != "" {
					matches = append(matches, SecretMatched{
						Secret: secret,
						URL:    u,
						Match:  match,
					})
				}
			}

			matchesCh <- matches

			if bar != nil {
				bar.Add(1)
			}
		}(url)
	}

	go func() {
		wg.Wait()
		close(matchesCh)
	}()

	seen := make(map[string]bool)
	for matches := range matchesCh {
		for _, match := range matches {
			key := fmt.Sprintf("%s:%s", match.Secret.Name, match.Match)
			if !seen[key] {
				seen[key] = true
				allMatches = append(allMatches, match)
			}
		}
	}

	return allMatches
}

func CheckURLForSecrets(url string, secret Secret) []SecretMatched {
	re := regexp.MustCompile(secret.Regex)
	matches := re.FindAllString(url, -1)

	var matchedSecrets []SecretMatched
	for _, match := range matches {
		matchedSecrets = append(matchedSecrets, SecretMatched{
			Secret: secret,
			URL:    url,
			Match:  match,
		})
	}

	return matchedSecrets
}
