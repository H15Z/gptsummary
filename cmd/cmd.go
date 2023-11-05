package cmd

import (
	"github.com/H15Z/gptsummary/configs"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "cmd",
	Short: "GPT Data Enrichment CLI Application",
	Long:  `GPT Data Enrichment CLI Application`,
}

func Execute() {
	configs.InitConfig()
	cobra.CheckErr(rootCmd.Execute())
}
