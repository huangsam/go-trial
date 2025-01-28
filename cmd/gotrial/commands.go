package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/gocolly/colly/v2"
	"github.com/huangsam/go-trial/pkg/abstraction"
	"github.com/huangsam/go-trial/pkg/basicintro"
	"github.com/huangsam/go-trial/pkg/concurrency"
	"github.com/urfave/cli/v3"
)

// demoCommand is a command to run a demo.
var demoCommand *cli.Command = &cli.Command{
	Name:        "demo",
	Usage:       "Run demo with some pkg functions",
	Description: "This command runs functions from multiple packages.",
	Action: func(ctx context.Context, c *cli.Command) error {
		slog.Debug(basicintro.GreetWorld())

		slog.Info(basicintro.GreetName("Peter"))

		circle := abstraction.Circle{Radius: 6}
		size := abstraction.Classify(circle)
		slog.Warn(fmt.Sprintf("Circle size is %v", size))

		answers := concurrency.GetAnswersWithChannels()
		slog.Error("Retrieved answers with channels", "answers", answers)

		return nil
	},
}

// serverCommand is a command to run an HTTP server.
var serverCommand *cli.Command = &cli.Command{
	Name:        "server",
	Usage:       "Run simple HTTP server",
	Description: "This command runs an HTTP server with one endpoint.",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "port",
			Value: ":8080",
			Usage: "HTTP server port",
		},
	},
	Action: func(ctx context.Context, c *cli.Command) error {
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			slog.Info(
				fmt.Sprintf("Got %s request from %s", r.Method, r.Host),
				"path", r.URL.Path,
			)
			fmt.Fprintf(w, "Hello, World!")
		})
		if err := http.ListenAndServe(c.String("port"), nil); err != nil {
			panic(err)
		}
		return nil
	},
}

var scrapeCommand *cli.Command = &cli.Command{
	Name:        "scrape",
	Usage:       "Run colly scraping on a website",
	Description: "This command scrapes hackerspaces.org",
	Action: func(ctx context.Context, c *cli.Command) error {
		collector := colly.NewCollector(
			colly.AllowedDomains("hackerspaces.org", "wiki.hackerspaces.org"),
		)

		collector.OnHTML("a[href]", func(e *colly.HTMLElement) {
			link := e.Attr("href")
			slog.Info("Link found", "link", link)
		})

		collector.OnRequest(func(r *colly.Request) {
			slog.Debug("Visit site", "link", r.URL.String())
		})

		collector.Visit("https://hackerspaces.org/")

		return nil
	},
}
