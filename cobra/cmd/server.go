package cmd

import (
	"letcode/cobra/handlers"

	"github.com/spf13/cobra"
)

var serverCmd = &cobra.Command{
	Use:   "run",
	Short: "server",
	Long:  `Demo is a test appcation for print things`,
	Run:   FixData,
}

func init() {
	//把serverCmd加入到rootCmd
	rootCmd.AddCommand(serverCmd)
}

//run 入口函数
func FixData(cmd *cobra.Command, args []string) {
	handler := handlers.NewHandler()
	handler.Init()
	handler.Serve()
}
