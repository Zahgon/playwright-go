//go:build ignore
// +build ignore

package main

import (
	"fmt"

	"github.com/playwright-community/playwright-go"
)

func assertErrorToNilf(message string, err error) { _ = "STUB: not implemented"; return }

func main() {
	startHttpServer()

	pw, err := playwright.Run()
	assertErrorToNilf("could not launch playwright: %w", err)
	browser, err := pw.Chromium.Launch(playwright.BrowserTypeLaunchOptions{
		Headless: playwright.Bool(true),
	})
	assertErrorToNilf("could not launch Chromium: %w", err)
	context, err := browser.NewContext()
	assertErrorToNilf("could not create context: %w", err)
	page, err := context.NewPage()
	assertErrorToNilf("could not create page: %w", err)
	_, err = page.Goto("http://localhost:1234")
	assertErrorToNilf("could not goto: %w", err)
	assertErrorToNilf("could not set content: %w", page.SetContent(`<a href="/download" download>download</a>`))
	download, err := page.ExpectDownload(func() error {
		return page.Locator("text=download").Click()
	})
	assertErrorToNilf("could not download: %w", err)
	fmt.Println(download.SuggestedFilename())
	assertErrorToNilf("could not close browser: %w", browser.Close())
	assertErrorToNilf("could not stop Playwright: %w", pw.Stop())
}

func startHttpServer() { _ = "STUB: not implemented"; return }
