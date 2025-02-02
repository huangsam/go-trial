package sub

import (
	"github.com/gocolly/colly/v2"
	"github.com/rs/zerolog/log"
	"github.com/urfave/cli/v2"
)

// ScrapeCommand is a command to scrape hackerspaces.org for links.
var ScrapeCommand *cli.Command = &cli.Command{
	Name:        "scrape",
	Usage:       "Run scraping on a website",
	Description: "This command uses Colly to scrape hackerspaces.org",
	Action: func(c *cli.Context) error {
		collector := colly.NewCollector(
			colly.AllowedDomains("hackerspaces.org", "wiki.hackerspaces.org"),
		)

		collector.OnHTML("a[href]", func(e *colly.HTMLElement) {
			log.Info().Str("link", e.Attr("href")).Msg("Link found")
		})

		collector.OnRequest(func(r *colly.Request) {
			log.Info().Str("link", r.URL.String()).Msg("Visit site")
		})

		return collector.Visit("https://hackerspaces.org/")
	},
}
