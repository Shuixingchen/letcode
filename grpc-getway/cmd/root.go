package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

var rootCmd = &cobra.Command{
	Use: "grpc",
	Short: "Run the gRPC hello-world server",
}
var cfgFile string

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}


func init() {
	rootCmd.Flags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.ethx.yaml)")

	cobra.OnInitialize(initConfig, initLogger)
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	}
}
func initLogger() {

}