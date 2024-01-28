package cmd

import (
	"github.com/spf13/cobra"
	"github.com/taylormonacelli/limesleep/extend2"
)

// extend2Cmd represents the extend2 command
var extend2Cmd = &cobra.Command{
	Use:   "extend2",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		extend2.Main()
	},
}

func init() {
	rootCmd.AddCommand(extend2Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// extend2Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// extend2Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
