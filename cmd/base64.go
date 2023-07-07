/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"go-cli/utils"

	"github.com/spf13/cobra"
)

// base64Cmd represents the base64 command
var base64Cmd = &cobra.Command{
	Use:   "base64",
	Short: "encode and decode string",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		encodeStr, _ := cmd.Flags().GetString("e")
		decodeStr, _ := cmd.Flags().GetString("d")
		if encodeStr != "" {
			encoded := utils.EncodedString(encodeStr)
			cmd.Println(encoded)
		} else if decodeStr != "" {
			decoded := utils.DecodeString(decodeStr)
			cmd.Println(decoded)
		}
	},
}

func init() {
	rootCmd.AddCommand(base64Cmd)
	base64Cmd.PersistentFlags().String("e", "", "Encode string")
	base64Cmd.PersistentFlags().String("d", "", "Decode string")
}
