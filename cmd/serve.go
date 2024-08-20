/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"go-echo/internal/app"

	"github.com/spf13/cobra"
)

func startServer(cmd *cobra.Command, args []string) {

	appServer := app.InitializeServer()
	app.LoadEnv(appServer)

	db := app.OpenConnection()
	app.LoadRouter(appServer, db)
	app.StartServer()
}

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "start app server",
	Run:   startServer,
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
