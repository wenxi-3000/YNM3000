/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"YNM3000/libs"
	"os"

	"github.com/spf13/cobra"
)

var options = libs.Options{}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "YNM3000",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.YNM3000.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.

	//输入
	rootCmd.PersistentFlags().StringVarP(&options.CmdInput.ResultPath, "result", "o", "./results", "结果保存目录")
	rootCmd.PersistentFlags().StringVarP(&options.CmdInput.Input, "input", "i", "", "目标")
	rootCmd.PersistentFlags().StringSliceVarP(&options.CmdInput.Inputs, "inputs", "I", []string{}, "多个目标")
	rootCmd.PersistentFlags().StringVarP(&options.CmdInput.InputFile, "inputFile", "f", "", "读入的文件")
	//rootCmd.PersistentFlags().StringVar(&options.Wordlist.Subdomain, "subdict", "./libs/dict/Top_Subdomains.txt", "子域名字典")
	rootCmd.PersistentFlags().StringVar(&options.Org, "org", "", "org")
	rootCmd.PersistentFlags().BoolVar(&options.Clean, "clean", true, "是否清理文件")
	cobra.OnInitialize(initConfig)
}

func initConfig() {
	libs.InitOptions(&options)
}
