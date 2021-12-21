package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile string
)

//1.创建一个cmd
var rootCmd = &cobra.Command{
	Use:   "",
	Short: "A test demo",
	Long:  `Demo is a test appcation for print things`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		StartProfiling()
	},
	PersistentPostRun: func(cmd *cobra.Command, args []string) {
		StopProfiling()
	},
}

func init() {
	rootCmd.Flags().StringVar(&cfgFile, "config", "./config.yaml", "config path")
	cobra.OnInitialize(initConfig)
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigType("yml")
		viper.SetConfigFile(cfgFile)
	}
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("read config is failed err:", err)
	}
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func StartProfiling() {

}

func StopProfiling() {

}
