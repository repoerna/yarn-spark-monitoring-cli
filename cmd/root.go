package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.yarn-spark-monitoring.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "yarn-spark-monitoring",
	Short: "call SPARK(with YARN Resource Manager) API",
	Long:  `Just use it or leave it`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	// rootCmd.AddCommand(sparkjob.Cmd())
	// rootCmd.AddCommand(SparkApplicationJobs)

	YarnClusterApps.AddCommand(SparkApplicationJobs)
	rootCmd.AddCommand(YarnClusterApps)

	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
