package ping

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	//Version will be set by build flags
	version = "dev"

	rootCmd = &cobra.Command{
		Use: "pingcli",
		Short: "A lightweight HTTP ping utility",
		Long: `PingCLI is a command-line tool for checking HTTP endpoint availability and response times.`,
		Version: version,
	}
)

//Execute runs the root command
func Execute() error {
	return rootCmd.Execute()
}

func init(){
	//Global flags
	rootCmd.PersistentFlags().BoolP("verbose","v",false,"verbose output")

	//Add subcommands
	rootCmd.AddCommand(pingCmd)
}
