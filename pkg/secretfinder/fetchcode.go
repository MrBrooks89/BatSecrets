package secretfinder

import (
	"context"
	"io"
	"log"
	"net/http"
	"time"
)

func fetchJavaScriptCode(url string, c *http.Client, verbose bool) (string, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return "", err
	}

	// Set a 5 second timeout for the request
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Use the context for the request
	req = req.WithContext(ctx)

	resp, err := c.Do(req)
	if err != nil {
		if verbose {
			log.Printf("Error fetching JavaScript code for URL %s: %v\n", url, err)
		}
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}
