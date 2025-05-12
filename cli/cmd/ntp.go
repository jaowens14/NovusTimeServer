/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/jowens25/axi"
	"github.com/spf13/cobra"
)

// ntpCmd represents the ntp command
var ntpCmd = &cobra.Command{
	Use:   "ntp",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		listMode, _ := cmd.Flags().GetBool("list")

		if listMode {
			axi.ListNtpProperties()
		}

		fmt.Println("ntp called")
	},
}

func init() {
	rootCmd.AddCommand(ntpCmd)
	ntpCmd.Flags().BoolP("list", "l", false, "list ntp properties and values")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// ntpCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// ntpCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
