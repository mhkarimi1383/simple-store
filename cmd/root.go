/*
Copyright © 2022 Muhammed Hussein Karimi <info@karimi.dev>
*/
package cmd

import (
	"github.com/mhkarimi1383/simple-store/api"
	"github.com/mhkarimi1383/simple-store/internal/config"
	"github.com/mhkarimi1383/simple-store/internal/flagloader"
	"github.com/mhkarimi1383/simple-store/internal/pathhelper"
	"github.com/mhkarimi1383/simple-store/types"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "simple-store",
	Short: "Simply Store your file",
	Long:  `Made to make your work easier as for saving files.`,
	Run:   start,
}

func Execute() {
	if err := flagloader.SetFlagsFromEnv(rootCmd.PersistentFlags(), "SS"); err != nil {
		log.Fatalln(err)
	}

	if err := rootCmd.Execute(); err != nil {
		log.Fatalln(err)
	}
}

var cfg types.Config

func init() {
	rootCmd.PersistentFlags().StringVar(&cfg.ListenAddress, "listen-address", "127.0.0.1:8080", "Address:Port combination used to serve API")
	rootCmd.PersistentFlags().StringVar(&cfg.BasePath, "base-path", "/data", "Path to use for saving files")
}

func start(_ *cobra.Command, _ []string) {
	log.Printf("%+v", cfg)
	log.Infoln("Making sure given base path is present...")
	pathhelper.CreatePath(cfg.BasePath)
	log.Infoln("Base path is ready, Let's GoOoOoOo")
	config.SetConfig(&cfg)
	api.Serve(cfg.ListenAddress)
}
