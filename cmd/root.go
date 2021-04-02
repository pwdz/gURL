// Copyright © 2021 NAME HERE me.adibzadeh@gmail.com
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"os"
	"log"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/pwdz/gurl/app"
)

var cfgFile string

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "gurl",
	Short: "go http client",
	Long: `gurl is a cli http client like curl`,

	Run: func(cmd *cobra.Command, args []string) {
		
		method, err := cmd.Flags().GetString("method")
		if err != nil{
			log.Fatal(err)
			return
		}
		rawHeaders, err := cmd.Flags().GetStringSlice("header")
		if err != nil{
			log.Fatal(err)
			return
		}
		rawQuerries, err := cmd.Flags().GetStringSlice("query")
		if err != nil{
			log.Fatal(err)
			return
		}
		data, err := cmd.Flags().GetString("data")
		if err != nil{
			log.Fatal(err)
			return
		}
		json, err := cmd.Flags().GetString("json")
		if err != nil{
			log.Fatal(err)
			return
		}
		timeout, err := cmd.Flags().GetInt("timeout")
		if err != nil{
			log.Fatal(err)
			return
		}

		fmt.Println("kir:" , rawHeaders)
		if err := app.Send(args[0], method, rawHeaders, rawQuerries, data, json, timeout); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() { 
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.gurl.yaml)")

	RootCmd.Flags().StringP("method", "M", app.DefaultMethod , "pass method(GET/POST/PATCH/DELETE/PUT). default value is GET.")
	RootCmd.Flags().StringSliceP("header", "H", nil, "pass headers")
	RootCmd.Flags().StringSliceP("query", "Q", nil, "pass querries")
	RootCmd.Flags().StringP("data", "D", "", "pass body data")
	RootCmd.Flags().StringP("json", "J", "", "pass body in json format")
	RootCmd.Flags().IntP("timeout", "T", -1, "request timeout")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".gurl" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".gurl")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
