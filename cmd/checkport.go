/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
	"go-cli/utils"
	"strings"
)

// checkportCmd represents the checkport command
var checkportCmd = &cobra.Command{
	Use:   "checkport",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		host, _ := cmd.Flags().GetString("h")
		if host == "" {
			cmd.Println("host must be informed --h")
			return
		}
		ports, _ := cmd.Flags().GetString("p")
		if ports == "" {
			cmd.Println("ports must be informed Ex: 80,433,8080")
			return
		}
		portList := strings.Split(ports, ",")
		utils.CheckPort(host, portList)
	},
}

func init() {
	rootCmd.AddCommand(checkportCmd)
	checkportCmd.PersistentFlags().String("h", "", "url to be validate")
	checkportCmd.PersistentFlags().String("p", "", "list por split by comma. Ex: 80,443,8080")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// checkportCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
