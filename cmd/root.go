package cmd

import (
	"context"
	"syscall"

	"github.com/charmbracelet/fang"
	"github.com/charmbracelet/log"
	"github.com/daanv2/go-code-grapher/pkg/logging"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "go-code-grapher",
	Short: "",
	Long: ``,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
	PersistentPreRunE: logging.ApplyLoggerFlags,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := fang.Execute(
		context.Background(),
		rootCmd,
		fang.WithNotifySignal(syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL),
	)

	if err != nil {
		log.Fatal("execution failed", "error", err)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.go-code-grapher.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	logging.Flags(rootCmd.PersistentFlags())
}
