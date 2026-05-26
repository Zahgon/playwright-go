//go:build ignore
// +build ignore

package main

import (
	"context"
	"fmt"
	"log"
	"math"
	"os"
	"path/filepath"

	"github.com/playwright-community/playwright-go"
)

func assertErrorToNilf(message string, err error) { _ = "STUB: not implemented"; return }

func worker(id int, jobs chan Job, results chan<- Job, browser playwright.Browser) {
	_ = "STUB: not implemented"
	return
}

func processJob(browser playwright.Browser, job Job, ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

type Job struct {
	URL     string
	Try     int
	err     error
	Success bool
}

func main() {
	log.Println("Downloading Alexa top domains")
	topDomains, err := getAlexaTopDomains()
	assertErrorToNilf("could not get alexa top domains: %w", err)
	log.Println("Downloaded Alexa top domains successfully")
	cwd, err := os.Getwd()
	if err != nil {
		assertErrorToNilf("could not get cwd %w", err)
	}
	if err := os.Mkdir(filepath.Join(cwd, "out"), 0o777); err != nil && !os.IsExist(err) {
		assertErrorToNilf("could not create output directory %w", err)
	}

	pw, err := playwright.Run()
	assertErrorToNilf("could not launch playwright: %w", err)
	browser, err := pw.Chromium.Launch(playwright.BrowserTypeLaunchOptions{
		Headless: playwright.Bool(false),
	})
	assertErrorToNilf("could not launch Chromium: %w", err)

	numberOfJobs := int(math.Min(30, float64(len(topDomains))))

	jobs := make(chan Job, numberOfJobs)
	results := make(chan Job, numberOfJobs)

	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results, browser)
	}

	for _, url := range topDomains[:numberOfJobs] {
		jobs <- Job{
			URL: url,
		}
	}

	for a := 0; a < numberOfJobs; a++ {
		job := <-results
		if job.Success {
			fmt.Println("success:", job.URL)
		} else {
			fmt.Println("error:", job.URL, job.err)
		}
	}

	close(jobs)
	close(results)

	assertErrorToNilf("could not close browser: %w", browser.Close())
	assertErrorToNilf("could not stop Playwright: %w", pw.Stop())
}

func getAlexaTopDomains() ([]string, error) { _ = "STUB: not implemented"; return nil, nil }
