package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/chromedp/chromedp"
)

func main() {
	// create a new context and cancel it after 10 seconds
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// set the URL of the website to capture
	url := "https://www.linkedin.com/in/romit-tajale-108b96242/"

	// create a new chrome instance and navigate to the website
	chromeCtx, cancel := chromedp.NewContext(ctx)
	defer cancel()
	err := chromedp.Run(chromeCtx, chromedp.Navigate(url))
	if err != nil {
		log.Fatal(err)
	}

	// set the size of the viewport and capture a screenshot of the website
	var buf []byte
	err = chromedp.Run(chromeCtx, chromedp.EmulateViewport(1920, 1080), chromedp.CaptureScreenshot(&buf))
	if err != nil {
		log.Fatal(err)
	}

	// write the screenshot to a file
	outputFile := "screenshot.png"
	err = os.WriteFile(outputFile, buf, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
