package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

type JattConfig struct {
	SiteConfig SiteConfig `mapstructure:",squash"`
	NavConfig  NavConfig  `mapstructure:"nav"`
	HomeConfig HomeConfig `mapstructure:"home"`
}

type SiteConfig struct {
	Title      string `mapstructure:"title"`
	Author     string `mapstructure:"author"`
	Theme      string `mapstructure:"theme"`
	ContentDir string `mapstructure:"contentDir"`
	OutputDir  string `mapstructure:"outputDir"`
}

type NavConfig struct {
	Title    string    `mapstructure:"title"`
	NavItems []NavItem `mapstructure:"items"`
}

type NavItem struct {
	Label  string `mapstructure:"label"`
	Url    string `mapstructure:"url"`
	Weight int    `mapstructure:"weight"`
}

type HomeConfig struct {
	Name        string   `mapstructure:"name"`
	Image       string   `mapstructure:"image"`
	Description string   `mapstructure:"description"`
	Social      []Social `mapstructure:"social"`
	Buttons     []Button `mapstructure:"buttons"`
}

type Social struct {
	Icon string `mapstructure:"icon"`
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
