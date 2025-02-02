package sub

import (
	"context"

	"github.com/gocolly/colly/v2"
	"github.com/rs/zerolog/log"
	"github.com/urfave/cli/v3"
)

// ScrapeCommand is a command to scrape hackerspaces.org for links.
var ScrapeCommand *cli.Command = &cli.Command{
	Name:        "scrape",
	Usage:       "Run scraping on a website",
	Description: "This command uses Colly to scrape hackerspaces.org",
	Action: func(ctx context.Context, c *cli.Command) error {
		collector := colly.NewCollector(
			colly.AllowedDomains("hackerspaces.org", "wiki.hackerspaces.org"),
		)

		collector.OnHTML("a[href]", func(e *colly.HTMLElement) {
			log.Info().Str("link", e.Attr("href")).Msg("Link found")
		})

		collector.OnRequest(func(r *colly.Request) {
			log.Info().Str("link", r.URL.String()).Msg("Visit site")
		})

		if err := collector.Visit("https://hackerspaces.org/"); err != nil {
			return err
		}

		return nil
	},
}
