package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gocolly/colly/v2"
	"github.com/huangsam/go-trial/pkg/abstraction"
	"github.com/huangsam/go-trial/pkg/basicintro"
	"github.com/huangsam/go-trial/pkg/concurrency"
	"github.com/rs/zerolog/log"
	"github.com/urfave/cli/v3"
)

// demoCommand is a command to run a demo.
var demoCommand *cli.Command = &cli.Command{
	Name:        "demo",
	Usage:       "Run demo with some pkg functions",
	Description: "This command runs functions from multiple packages.",
	Action: func(ctx context.Context, c *cli.Command) error {
		log.Debug().Msg(basicintro.GreetWorld())

		log.Info().Msg(basicintro.GreetName("Peter"))

		circle := abstraction.Circle{Radius: 6}
		size := abstraction.Classify(circle)
		log.Warn().Msgf("Circle size is %v", size)

		answers := concurrency.GetAnswersWithChannels()
		log.Error().Interface("answers", answers).Msg("Got answers with channels")

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
			log.Info().
				Str("method", r.Method).
				Str("host", r.Host).
				Str("path", r.URL.Path).
				Msg("Got request")

			fmt.Fprintf(w, "Hello, World!")
		})
		if err := http.ListenAndServe(c.String("port"), nil); err != nil {
			return err
		}
		return nil
	},
}

// scrapeCommand is a command to scrape hackerspaces.org for links.
var scrapeCommand *cli.Command = &cli.Command{
	Name:        "scrape",
	Usage:       "Run scraping on a website",
	Description: "This command uses Colly to scrape hackerspaces.org",
	Action: func(ctx context.Context, c *cli.Command) error {
		collector := colly.NewCollector(
			colly.AllowedDomains("hackerspaces.org", "wiki.hackerspaces.org"),
		)

		collector.OnHTML("a[href]", func(e *colly.HTMLElement) {
			link := e.Attr("href")
			log.Info().Str("link", link).Msg("Link found")
		})

		collector.OnRequest(func(r *colly.Request) {
			log.Debug().Str("link", r.URL.String()).Msg("Visit site")
		})

		if err := collector.Visit("https://hackerspaces.org/"); err != nil {
			return err
		}

		return nil
	},
}
