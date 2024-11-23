package cmd

import (
	"fmt"
	"os"

	"github.com/devmegablaster/jatt/internal/config"
	"github.com/devmegablaster/jatt/pkg"
	"github.com/devmegablaster/jatt/pkg/styles"
	"github.com/spf13/cobra"
)

const logo = `
       _       __  __ 
      (_)___ _/ /_/ /_
     / / __ ` + "`" + `/ __/ __/
    / / /_/ / /_/ /_ 
 __/ /\__,_/\__/\__/  
/___/                 
`

var rootCmd = &cobra.Command{
	Use:   "Jatt - Static Site Generator",
	Short: "Jatt is used to build static sites",
	Long:  `Jatt is a static site generator that is used to build static sites.`,
	Run: func(cmd *cobra.Command, args []string) {
		cfg := config.NewConfig()
		fmt.Println(styles.LogoStyle.Render(logo))

		jatt := pkg.NewJatt(cfg)
		jatt.Run()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(versionCommand)
	rootCmd.AddCommand(initJatt)
}
