package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

var (
	serverHost string
	serverPort uint16
	username   string
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use: "metro",
	// Run: func(cmd *cobra.Command, args []string) {
	// },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringVar(&serverHost, "host", "0.0.0.0", "Host string or IP that the Metro server serving")
	rootCmd.Flags().Uint16Var(&serverPort, "port", 50051, "Port number that the Metro server exposing")
	rootCmd.Flags().StringVarP(&username, "user", "u", "metro", "username")
}

func getServerAddress() string {
	return serverHost + ":" + strconv.Itoa(int(serverPort))
}
