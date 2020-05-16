package main

import (
	"fmt"
	"log"

	// "github.com/gocolly/twocaptcha"
	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
	slog "github.com/tebeka/selenium/log"
)

func main() {

	// url := "https://www.bloomberg.com/news/articles/2020-03-11/augmented-reality-startup-magic-leap-is-said-to-explore-a-sale"
	url := "https://www.leboncoin.fr/voitures/1781740745.htm/"

	caps := selenium.Capabilities{"browserName": "chrome"}
	chromeCaps := chrome.Capabilities{
		Args: []string{
			"--headless",
			"--no-sandbox",
			"--start-maximized",
			"--window-size=1024,768",
			"--disable-crash-reporter",
			"--hide-scrollbars",
			"--disable-gpu",
			"--disable-setuid-sandbox",
			"--disable-infobars",
			"--window-position=0,0",
			"--ignore-certifcate-errors",
			"--ignore-certifcate-errors-spki-list",
			"--user-agent=Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_2) AppleWebKit/604.4.7 (KHTML, like Gecko) Version/11.0.2 Safari/604.4.7",
			// "--proxy-server=sock5://localhost:8119", // 5566 // 8119
			// "--host-resolver-rules=\"MAP * 0.0.0.0 , EXCLUDE localhost\"",
		},
	}
	caps.AddChrome(chromeCaps)

	caps.SetLogLevel(slog.Server, slog.Off)
	caps.SetLogLevel(slog.Browser, slog.Off)
	caps.SetLogLevel(slog.Client, slog.Off)
	caps.SetLogLevel(slog.Driver, slog.Off)
	caps.SetLogLevel(slog.Performance, slog.Off)
	caps.SetLogLevel(slog.Profiler, slog.Off)

	wd, err := selenium.NewRemote(caps, fmt.Sprintf("http://selenium:%d/wd/hub", 4444))
	if err != nil {
		log.Fatal(err)
	}
	defer wd.Quit()

	err = wd.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	src, err := wd.PageSource()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("source", src)

}
