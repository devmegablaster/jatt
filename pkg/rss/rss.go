package rss

import (
	"encoding/xml"
	"fmt"
	"time"

	"github.com/devmegablaster/jatt/internal/config"
	"github.com/devmegablaster/jatt/pkg/reader"
	"github.com/devmegablaster/jatt/pkg/styles"
)

type Rss struct {
	cfg config.JattConfig
}

func New(cfg config.JattConfig) *Rss {
	return &Rss{
		cfg: cfg,
	}
}

type Feed struct {
	Title       string               `xml:"title"`
	Description string               `xml:"description"`
	Link        string               `xml:"link"`
	PubDate     time.Time            `xml:"pubDate"`
	Items       []reader.ListingItem `xml:"item"`
}

func (r *Rss) GenerateFeed(files []reader.File) []byte {
	posts := []reader.ListingItem{}

	for _, file := range files {
		if file.FrontMatter.Draft {
			continue
		}

		if file.FrontMatter.Layout == "listing" {
			for _, item := range file.Listing {
				posts = append(posts, item)
			}
		}
	}

	feed := Feed{
		Title:       r.cfg.SiteConfig.Title,
		Description: r.cfg.SiteConfig.Description,
		Link:        r.cfg.SiteConfig.BaseURL,
		PubDate:     time.Now(),
		Items:       posts,
	}

	enc, err := xml.MarshalIndent(&feed, "", "  ")
	if err != nil {
		fmt.Println(styles.ErrorStyle.Render("Error generating RSS feed"))
	}

	fmt.Println(styles.DebugStyle.Render("Generated RSS feed"))

	return enc
}
