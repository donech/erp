package cmd

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/donech/tool/xlog"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "erp",
	Short: "A smart erp system of donech universe",
	Long:  `A smart erp system of donech universe`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("request at least one arg")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("Welcome for use donech erp system, see more information with -h flag")
	},
}

func init() {
	cobra.OnInitialize(initConfig, initLogger)
	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "app.yaml", "-c app.yaml")
	rootCmd.AddCommand(exampleCmd)
}

func initConfig() {
	log.Println("using config file: ", cfgFile)
	viper.SetConfigFile(cfgFile)
	viper.AutomaticEnv() // read in environment variables that match
	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err != nil {
		log.Panicln("Read config file failed:", viper.ConfigFileUsed(), err.Error())
	}
}

func initLogger() {
	cfg := xlog.Config{}
	err := viper.Sub("log").Unmarshal(&cfg)
	if err != nil {
		log.Fatalln("can't unmarshal viper to Config :", err)
	}
	fmt.Println(cfg)
}

//Execute the endpoint to expose
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
