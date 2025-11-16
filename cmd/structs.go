package cmd

import (
	"github.com/spf13/cobra"
)

// structsCmd represents the structs command
var structsCmd = &cobra.Command{
	Use:   "structs",
	Short: "",
	Long: ``,
	Example: "structs ./path/to/project ./another/path",
	RunE: GraphStructs,
}

func init() {
	rootCmd.AddCommand(structsCmd)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// structsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:

	structsCmd.Flags().BoolP("recursive", "r", true, "Recursively scan directories for Go files")
	structsCmd.Flags().StringArrayP("dir", "p", []string{"."}, "The directory to parse and consume")
}

func GraphStructs(cmd *cobra.Command, args []string) error {
	// recursive, _ := cmd.Flags().GetBool("recursive")
	// dirs, _ := cmd.Flags().GetStringArray("dir")



	return nil
}