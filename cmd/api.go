/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"os"
	"os/signal"

	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
	"gitlab.internal.b2w.io/team/a-tech/star-wars/src/server"
	"gitlab.internal.b2w.io/team/a-tech/star-wars/src/server/config"
	"gitlab.internal.b2w.io/team/a-tech/star-wars/src/shared/providers/logger"
)

// apiCmd represents the api command
var (
	port        int
	addr        string
	environment string
	apiCmd      = &cobra.Command{
		Use:   "api",
		Short: "API start server",
		Long:  ``,
		Run:   run,
	}
)

func run(cmd *cobra.Command, args []string) {
	if err := godotenv.Load("./environments/.env." + environment); err != nil {
		panic(err)
	}

	os.Getenv("ENV")

	cfg, err := config.ReadConfigFromEnv(environment)

	if err != nil {
		panic(err)
	}

	svr := server.New(addr, port)
	log := logger.New()
	svr.Run(cfg, log)

	chanExit := make(chan os.Signal, 1)
	signal.Notify(chanExit, os.Interrupt)
	<-chanExit
}

func init() {
	rootCmd.AddCommand(apiCmd)
	// Get start server options
	apiCmd.PersistentFlags().IntVarP(&port, "port", "p", 3000, "The -p option specified port HTTP server")
	apiCmd.PersistentFlags().StringVarP(&addr, "address", "b", "127.0.0.1", "The -b option binds specified IP, by default it is localhost")
	apiCmd.PersistentFlags().StringVarP(&environment, "environment", "e", "development", "The -e option specified the environment")
}
