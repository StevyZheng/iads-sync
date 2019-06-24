package cmd

import (
	"github.com/spf13/cobra"
	"iads/server"
	"iads/server/models"
)

func init() {
	rootCmd.AddCommand(serverCmd)
}

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "restful api server",
	Run: func(cmd *cobra.Command, args []string) {
		models.CreateTable()
		server.ServerStart()
	},
}
