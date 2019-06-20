package cmd

import (
	"github.com/spf13/cobra"
	"iads/server"
	"iads/server/routers/api/v1/user"
)

func init() {
	rootCmd.AddCommand(serverCmd)
}

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "restful api server",
	Run: func(cmd *cobra.Command, args []string) {
		user.Init()
		user.CreateTable()
		server.ServerStart()
	},
}
