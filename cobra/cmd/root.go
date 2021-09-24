package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	name    string
	age     int
	cfgFile string
)

//1.创建一个cmd
var rootCmd = &cobra.Command{
	Use:   "testCobra",
	Short: "A test demo",
	Long:  `Demo is a test appcation for print things`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(name) == 0 {
			fmt.Println("no name")
			return
		}
		Show(name, age)
	},
}

func init() {
	rootCmd.Flags().StringVarP(&name, "name", "n", "", "person's name")
	rootCmd.Flags().IntVarP(&age, "age", "a", 0, "person's age")
	rootCmd.Flags().StringVar(&cfgFile, "conf", "", "config path")
	cobra.OnInitialize(initConfig, initDb)
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	}
}
func initDb() {

}

func Show(name string, age int) {
	fmt.Printf("My Name is %s, My age is %d\n", name, age)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
