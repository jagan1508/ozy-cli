/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var dataFile string

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "ozy",
	Short: "Ozy is a simple todo application on your cli",
	Long: `Ozy is a simple todo application on your cli which helps
you organize your tasks and projects in a simple manner`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func initConfig() {
	viper.SetConfigName(".ozy")
	viper.AddConfigPath("$PWD")
	viper.AutomaticEnv()
	viper.SetEnvPrefix("ozy")

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("using config file:", viper.ConfigFileUsed())
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.ozy.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	cobra.OnInitialize(initConfig)
	home, err := homedir.Dir()
	if err != nil {
		log.Println("no Home address . Set it first using --datafile")
	}

	rootCmd.PersistentFlags().StringVar(&dataFile,
		"datafile",
		home+string(os.PathSeparator)+".ozytodo.json",
		"data file to store todo list")

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file(default is $PWD/.ozy.yaml)")

	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
