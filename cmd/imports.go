package cmd

import (
	"path/filepath"
	"slices"
	"strings"

	"github.com/daanv2/go-code-grapher/pkg/ast"
	"github.com/daanv2/go-code-grapher/pkg/extensions/xflags"
	"github.com/daanv2/go-code-grapher/pkg/extensions/xos"
	"github.com/daanv2/go-code-grapher/pkg/extensions/xregexp"
	"github.com/daanv2/go-code-grapher/pkg/extensions/xslices"
	"github.com/daanv2/go-code-grapher/pkg/grapher"
	"github.com/spf13/cobra"
)

// importsCmd represents the imports command
var importsCmd = &cobra.Command{
	Use:   "imports",
	Short: "generate graphs of imports between packages",
	Long:  `generate graphs of imports between packages`,
	RunE:  GraphImports,
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

	importsCmd.Flags().BoolP("recursive", "r", true, "Recursively scan directories for Go files")
	importsCmd.Flags().StringArrayP("dir", "d", []string{"."}, "The directory to parse and consume")
	importsCmd.Flags().String("mod-file", "go.mod", "The go.mod file to read the module path from")

	// filters
	importsCmd.Flags().StringArray("filter-packages", []string{}, "The regex pattern to filter packages by, if empty all packages are allowed")
	importsCmd.Flags().StringArray("filter-imports", []string{}, "The regex pattern to filter imports by, if empty all imports are allowed")
	importsCmd.Flags().Bool("filter-dirs", true, "Filters out any package that was not in the provided directories")

	grapher.ImportsGraphers.AddFlags(importsCmd.Flags())
}

func GraphImports(cmd *cobra.Command, args []string) error {
	recursive, _ := cmd.Flags().GetBool("recursive")
	dirs, _ := cmd.Flags().GetStringArray("dir")
	modFile, _ := cmd.Flags().GetString("mod-file")

	// Collect
	col, err := ast.NewImportCollector(modFile)
	if err != nil {
		return err
	}
	
	for file := range xos.AllGoFiles(dirs, recursive) {
		err = col.Collect(file)
		if err != nil {
			return err
		}
	}

	// Clean
	err = cleanImports(col, dirs, cmd)
	if err != nil {
		return err
	}

	err = xflags.SetIfUnchanged(cmd.Flags(), "annotations", "title=" + col.ModuleName())
	if err != nil {
		return err
	}

	// Graph with fully qualified package names
	return grapher.ImportsGraphers.Process(col, cmd.Flags())
}

func cleanImports(col *ast.ImportCollector, dirs []string, cmd *cobra.Command) error {
	// Filter Packages
	pks, _ := cmd.Flags().GetStringArray("filter-packages")
	pkgfilter, err := xregexp.FromPatterns(pks)
	if err != nil {
		return err
	}

	if pkgfilter.Len() > 0 {
		for pack := range col.Imports() {
			if !pkgfilter.Match(pack) {
				delete(col.Imports(), pack)
			}
		}
	}

	// Filter Imports
	imps, _ := cmd.Flags().GetStringArray("filter-imports")
	impfilter, err := xregexp.FromPatterns(imps)
	if err != nil {
		return err
	}
	if impfilter.Len() > 0 {
		for pack, imports := range col.Imports() {
			col.Imports()[pack] = impfilter.Filter(imports)
		}
	}

	// Filter Dirs
	filterDirs, _ := cmd.Flags().GetBool("filter-dirs")
	if filterDirs {
		keep := make(map[string]struct{})
		dirs, err = xslices.MapE(dirs, filepath.Abs)
		if err != nil {
			return err
		}

		for dir, packages := range col.DirPackages() {
			dir, err = filepath.Abs(dir)
			if err != nil {
				return err
			}

			if slices.ContainsFunc(dirs, func(d string) bool {
				return strings.HasPrefix(dir, d)
			}) {
				for _, p := range packages {
					keep[p] = struct{}{}
				}
			}
		}

		for pack := range col.Imports() {
			if _, ok := keep[pack]; !ok {
				delete(col.Imports(), pack)
			}
		}
	}

	return nil
}
