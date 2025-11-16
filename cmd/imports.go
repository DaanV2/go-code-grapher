/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/daanv2/go-code-grapher/pkg/ast"
	"github.com/daanv2/go-code-grapher/pkg/extensions/xos"
	"github.com/spf13/cobra"
)

// importsCmd represents the imports command
var importsCmd = &cobra.Command{
	Use:   "imports",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: GraphImports,
}

func init() {
	rootCmd.AddCommand(importsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// importsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// importsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	structsCmd.Flags().BoolP("recursive", "r", true, "Recursively scan directories for Go files")
	structsCmd.Flags().StringArrayP("dir", "p", []string{"."}, "The directory to parse and consume")
}

func GraphImports(cmd *cobra.Command, args []string) error {
	recursive, _ := cmd.Flags().GetBool("recursive")
	dirs, _ := cmd.Flags().GetStringArray("dir")

	col := ast.NewImportCollector()
	for file := range xos.AllGoFiles(dirs, recursive) {
		err := col.Collect(file)
		if err != nil {
			return err
		}
	}

	return nil
}