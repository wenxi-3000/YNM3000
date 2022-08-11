/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"YNM3000/code/core"
	"YNM3000/code/logger"
	"YNM3000/code/utils"
	"os"
	"path"

	"github.com/spf13/cobra"
)

// scanCmd represents the scan command
var scanCmd = &cobra.Command{
	Use:   "scan",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: runScan,
}

func init() {
	//workflow
	rootCmd.PersistentFlags().StringVar(&options.Scan.Flow, "flow", "general", "指定workflow")
	//workflow
	rootCmd.PersistentFlags().StringVar(&options.Scan.FlowFolder, "flowPath", "", "指定workflow的目录")
	rootCmd.AddCommand(scanCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// scanCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// scanCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func runScan(_ *cobra.Command, _ []string) {

	initScan()

	//打印要跑的目标
	logger.Info("输入的目标如下: ")
	for input := range options.Inputs {
		logger.Println(input)
	}
	for input := range options.Inputs {
		logger.Info("开始扫描: ", input)
		core.Run(input, options)
	}
}

func initScan() {
	//设置workflow的folder
	if options.Scan.FlowFolder == "" {
		options.Scan.FlowFolder = path.Join(options.Paths.Root, "workflow")
		if !utils.FolderExists(options.Scan.FlowFolder) {
			logger.Info("workflow目录不存在")
			os.Exit(1)
		}
	}
}
