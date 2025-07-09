// This is for testing purposes only

package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var logCmd = &cobra.Command{
	Use:   "log",
	Short: "Logs a word",
	Long: `Logs a word. 

Long description test`,
	Run: runLog,
}

func init() {
	rootCmd.AddCommand(logCmd)
	logCmd.Flags().StringP("word", "W", "", "word to be logged")
}

func runLog(cmd *cobra.Command, _ []string) {
	word, err := cmd.Flags().GetString("word")
	if err != nil {
		fmt.Println(err, "Command failed to execute")
	}

	fmt.Println("The word you entered is: " + word)

}
