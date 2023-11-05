package cmd

import (
	"github.com/H15Z/gptsummary/domain/actors"
	"github.com/spf13/cobra"
)

var syncjobCmd = &cobra.Command{
	Use:   "stream",
	Short: "Streams and Enriches Articles ",
	Long:  `Streams article data from CSV file uses `,
	Run: func(cmd *cobra.Command, args []string) {
		// runJob(args[0], args[1])

		count, _ := cmd.Flags().GetInt("count")
		threads, _ := cmd.Flags().GetInt("threads")

		// if dry_run {
		// 	log.Println("--------------------------------------")
		// 	log.Println("====== DRY RUN - NO API UPDATES ======")
		// 	log.Println("--------------------------------------")
		// }

		actors.StartStream(count, threads)
	},
}

func init() {
	rootCmd.AddCommand(syncjobCmd)
	syncjobCmd.Flags().IntP("count", "c", 5, "Article Count Limit")   // this is set low not to use up all GPT credits during testing
	syncjobCmd.Flags().IntP("threads", "t", 5, "Article Count Limit") // this is set low not to use up all GPT credits during testing
	// syncjobCmd.Flags().BoolP("dry", "d", false, "Dry Run - Skips API Enrychment But Simulates API Delay for performance testing")
}
