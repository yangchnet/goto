/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	info   string
	passwd string
	secret string
)

// serverCmd represents the server command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "add a remote server",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(info)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serverCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serverCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	addCmd.Flags().StringVarP(&info, "des", "i", "", "server_name:user_name:server_ip")
	addCmd.Flags().StringVar(&passwd, "passwd", "", "server password")
	addCmd.Flags().StringVar(&secret, "secret", "", "secret file")
}
