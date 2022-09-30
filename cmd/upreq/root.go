package upreq

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"upreq/pkg/upreq"
)

var rootCmd = &cobra.Command{
	Use:   "upreq",
	Short: "upreq - a small CLI to help manage your requirements.txt file",
	Long:  `upreq is a small CLI to help manage your requirements.txt file.`,
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {

		// grab current requirements
		var file = cmd.Flag("file").Value.String()
		var oldReqs = upreq.GetReqs(file)
		fmt.Printf("Found (%[1]s) requirements in %[2]s\n", fmt.Sprint(len(oldReqs)), file)

		// wipe the file
		upreq.WipeFile(file)

		// get current env requirements
		var newReqs = upreq.GetEnvReqs()

		// check for differences
		var diff = upreq.DiffCheck(oldReqs, newReqs)
		upreq.DisplayDiff(diff, cmd.Flag("strip").Value.String())

		// write new requirements
		var writtenReqs = upreq.WriteReqs(file, newReqs)
		fmt.Printf("Wrote (%[1]s) requirements to %[2]s\n", fmt.Sprint(len(writtenReqs)), file)

	},
}

func Execute() {

	// add all the flags
	rootCmd.Flags().StringP("file", "f", "requirements.txt", "Specify the requirements file to use")
	rootCmd.Flags().BoolP("strip", "s", false, "Strips '+' and '-' from the output")
	rootCmd.Flags().BoolP("git", "g", false, "Automatically add the file to git")

	// run root command
	if err := rootCmd.Execute(); err != nil {
		_, err := fmt.Fprintf(os.Stderr, "Whoops. There was an error while executing your CLI '%s'", err)
		if err != nil {
			return
		}
		os.Exit(1)
	}
}
