package sitemap

import (
	"encoding/xml"
	"fmt"
	"strings"

	"github.com/devmegablaster/jatt/internal/config"
	"github.com/devmegablaster/jatt/pkg/reader"
	"github.com/devmegablaster/jatt/pkg/styles"
)

type SitemapSvc struct {
	cfg config.JattConfig
}

func New(cfg config.JattConfig) *SitemapSvc {
	return &SitemapSvc{
		cfg: cfg,
	}
}

type Sitemap struct {
	XMLName xml.Name `xml:"urlset"`
	XMLNS   string   `xml:"xmlns,attr"`
	URLs    []URL    `xml:"url"`
}

type URL struct {
	Loc string `xml:"loc"`
}

func (s *SitemapSvc) GenerateSitemap(files []reader.File) []byte {
	sitemap := Sitemap{
		XMLNS: "http://www.sitemaps.org/schemas/sitemap/0.9",
		URLs:  []URL{},
	}

	for _, file := range files {
		if file.FrontMatter.Draft {
			continue
		}

		name := strings.ReplaceAll(file.Name, "index", "")

		sitemap.URLs = append(sitemap.URLs, URL{
			Loc: s.cfg.SiteConfig.BaseURL + "/" + name,
		})
	}

	output, err := xml.MarshalIndent(sitemap, "", "    ")
	if err != nil {
		panic(err)
	}

	enc := []byte(xml.Header + string(output))

	fmt.Println(styles.DebugStyle.Render("Generated sitemap"))

	return enc
}

func (s *SitemapSvc) GenerateRobots() []byte {
	content := "User-agent: *\nDisallow:\nSitemap: " + s.cfg.SiteConfig.BaseURL + "/sitemap.xml\n"
	fmt.Println(styles.DebugStyle.Render("Generated robots.txt"))
	return []byte(content)
}
