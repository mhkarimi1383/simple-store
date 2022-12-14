// Package cmd the magic starts from here
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

var (
	rootCmd = &cobra.Command{
		Use:     "simple-store",
		Aliases: []string{"ss"},
		Short:   "Simply Store your file",
		Long:    `Made to make your work easier as for saving files.`,
		Run:     start,
	}
	cfg types.Config
)

// Execute this is the main function to start the main process
func Execute() {
	if err := flagloader.SetFlagsFromEnv(rootCmd.PersistentFlags(), "SS"); err != nil {
		log.Fatalln(err)
	}

	if err := rootCmd.Execute(); err != nil {
		log.Fatalln(err)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&cfg.ListenAddress, "listen-address", "l", "127.0.0.1:8080", "Address:Port combination used to serve API")
	rootCmd.PersistentFlags().StringVarP(&cfg.BasePath, "base-path", "d", "/data", "Path to use for saving files")
	rootCmd.PersistentFlags().BoolVar(&cfg.EnableSwagger, "enable-swagger", false, "Enable swagger or not? (Do not use it in production.!)")
	rootCmd.PersistentFlags().Int64VarP(&cfg.ChunkSize, "chunk-size", "c", 50, "Size to use for each part for chucking files (In bytes)")
}

func start(_ *cobra.Command, _ []string) {
	log.Infoln("Making sure given base path is present...")
	pathhelper.CreatePath(cfg.BasePath)
	log.Infoln("Base path is ready, Let's GoOoOoOo")
	config.SetConfig(&cfg)
	api.Serve(cfg.ListenAddress, cfg.EnableSwagger)
}
