package rss

import (
	"encoding/xml"
	"fmt"
	"time"

	"github.com/devmegablaster/jatt/internal/config"
	"github.com/devmegablaster/jatt/pkg/reader"
	"github.com/devmegablaster/jatt/pkg/renderer"
	"github.com/devmegablaster/jatt/pkg/styles"
)

type RssSvc struct {
	cfg config.JattConfig
}

func New(cfg config.JattConfig) *RssSvc {
	return &RssSvc{
		cfg: cfg,
	}
}

type Rss struct {
	XMLName      xml.Name `xml:"rss"`
	Version      string   `xml:"version,attr"`
	XMLNSContent string   `xml:"xmlns:content,attr"`
	Channel      Channel  `xml:"channel"`
}

type Channel struct {
	Title       string    `xml:"title"`
	Description string    `xml:"description"`
	Link        string    `xml:"link"`
	PubDate     time.Time `xml:"pubDate"`
	Items       []Item    `xml:"item"`
}

type Item struct {
	Title       string    `xml:"title"`
	Link        string    `xml:"link"`
	Description string    `xml:"description"`
	PubDate     time.Time `xml:"pubDate"`
	Content     string    `xml:"content:encoded"`
}

func (r *RssSvc) GenerateFeed(files []reader.File, renderedFiles []renderer.RenderedFile) []byte {
	posts := []Item{}

	for _, file := range files {
		if file.FrontMatter.Draft {
			continue
		}

		if file.FrontMatter.Layout == "listing" {
			for _, item := range file.Listing {
				t, err := time.Parse("2006-01-02", item.Date)
				if err != nil {
					t = time.Now()
					fmt.Println(styles.ErrorStyle.Render("Error parsing date"))
				}

				content := ""

				for _, renderedFile := range renderedFiles {
					if renderedFile.ID == item.ID {
						content = renderedFile.HtmlContent
					}
				}

				posts = append(posts, Item{
					Title:       item.Title,
					Link:        item.URL,
					Description: item.Description,
					PubDate:     t,
					Content:     content,
				})
			}
		}
	}

	channel := Channel{
		Title:       r.cfg.SiteConfig.Title,
		Description: r.cfg.SiteConfig.Description,
		Link:        r.cfg.SiteConfig.BaseURL,
		PubDate:     time.Now(),
		Items:       posts,
	}

	feed := Rss{
		Version:      "2.0",
		XMLNSContent: "http://purl.org/rss/1.0/modules/content/",
		Channel:      channel,
	}

	enc, err := xml.MarshalIndent(&feed, "", "  ")
	if err != nil {
		fmt.Println(styles.ErrorStyle.Render("Error generating RSS feed"))
	}

	enc = []byte(xml.Header + string(enc))

	fmt.Println(styles.DebugStyle.Render("Generated RSS feed"))

	return enc
}
