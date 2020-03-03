package main

import (
	"github.com/mpuzanov/calendar/cmd/calendar/grpc"
	"github.com/mpuzanov/calendar/cmd/calendar/web"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "calendar",
	Short: "Calendar is a calendar microservice demo",
}

func init() {
	rootCmd.AddCommand(grpc.ServerCmd, web.ServerCmd)
}
