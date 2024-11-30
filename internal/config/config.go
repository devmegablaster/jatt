package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

type JattConfig struct {
	SiteConfig SiteConfig `mapstructure:",squash" yaml:",inline"`
	NavConfig  NavConfig  `mapstructure:"nav" yaml:"nav"`
	HomeConfig HomeConfig `mapstructure:"home" yaml:"home"`
}

type SiteConfig struct {
	BaseURL     string `mapstructure:"baseURL"`
	Title       string `mapstructure:"title"`
	Description string `mapstructure:"description"`
	Author      string `mapstructure:"author"`
	Theme       string `mapstructure:"theme"`
	ContentDir  string `mapstructure:"contentDir"`
	OutputDir   string `mapstructure:"outputDir"`
}

type NavConfig struct {
	Title    string    `mapstructure:"title"`
	NavItems []NavItem `mapstructure:"items" yaml:"items"`
}

type NavItem struct {
	Label string `mapstructure:"label"`
	Url   string `mapstructure:"url"`
}

type HomeConfig struct {
	Name        string   `mapstructure:"name"`
	Image       string   `mapstructure:"image"`
	Description string   `mapstructure:"description"`
	Social      []Social `mapstructure:"social" yaml:"social"`
	Buttons     []Button `mapstructure:"buttons" yaml:"buttons"`
}

type Social struct {
	Name string `mapstructure:"name"`
	URL  string `mapstructure:"url"`
}

type Button struct {
	Name string
	URL  string
}

func NewConfig() JattConfig {
	viper.SetConfigFile("jatt.yaml")
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Error reading config", err)
		os.Exit(1)
	}

	cfg := JattConfig{}
	if err := viper.Unmarshal(&cfg); err != nil {
		fmt.Println("Unable to unmarshall config", err)
		os.Exit(1)
	}

	return cfg
}

func NewConfigTemplate() JattConfig {
	return JattConfig{
		SiteConfig: SiteConfig{
			Title:      "Jatt Site",
			Author:     "Author",
			Theme:      "dark",
			ContentDir: "content",
			OutputDir:  "public",
		},
		NavConfig: NavConfig{
			Title: "Jatt",
			NavItems: []NavItem{
				{
					Label: "Home",
					Url:   "/",
				},
				{
					Label: "About",
					Url:   "/about",
				},
			},
		},
		HomeConfig: HomeConfig{
			Name:        "Jatt",
			Image:       "https://via.placeholder.com/150",
			Description: "Jatt Site",
			Social: []Social{
				{
					Name: "github",
					URL:  "https://github.com",
				},
			},
			Buttons: []Button{
				{
					Name: "About",
					URL:  "/about",
				},
				{
					Name: "Jatt",
					URL:  "https://github.com/devmegablaster/jatt",
				},
			},
		},
	}
}
