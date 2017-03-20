// Copyright © 2016 Jose Carlos Ustra Junior <dev@ustrajunior.com>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/boltdb/bolt"
	"github.com/cloudflare/cloudflare-go"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	api       *cloudflare.API
	cfgFile   string
	zone      string
	dnsRecord string
	ip        string
	force     bool

	db *bolt.DB
)

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "keepup",
	Short: "Update your cloudflare Records",
	Long: `A command line to update your cloudflare.com dns records
	with your current ip or with other given ip`,
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	db = openBoltDB()
	defer db.Close()

	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

// Domain a single domain on the dns
type Domain struct {
	Zone string
	DNS  string
	IP   string
}

func init() {
	cobra.OnInitialize(initConfig)

	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.keepup.yaml)")
	RootCmd.PersistentFlags().StringVar(&zone, "zone", "", "the dns zone to work on (domain.com)")
	RootCmd.PersistentFlags().StringVar(&dnsRecord, "dns", "", "the dns record to be updated (my.domain.com)")
	RootCmd.PersistentFlags().StringVar(&ip, "ip", "", "the ip that will be used to update the dns record (127.0.0.1)")
	RootCmd.PersistentFlags().BoolVar(&force, "force", false, "forces the dns update even if the ip is the same")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" { // enable ability to specify config file via flag
		viper.SetConfigFile(cfgFile)
	} else {
		viper.SetConfigName(".keepup") // name of config file (without extension)
		viper.AddConfigPath("$HOME")   // adding home directory as first search path
		viper.AutomaticEnv()           // read in environment variables that match
	}

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}

func openBoltDB() *bolt.DB {
	keepupFolder := fmt.Sprintf("%s/.keepup/", os.Getenv("HOME"))

	if _, err := os.Stat(keepupFolder); os.IsNotExist(err) {
		os.Mkdir(keepupFolder, 0777)
	}

	db, err := bolt.Open(keepupFolder+"keepup.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}

	tx, err := db.Begin(true)
	if err != nil {
		return nil
	}
	defer tx.Rollback()

	_, err = tx.CreateBucketIfNotExists([]byte("domains"))
	if err != nil {
		log.Fatal(err)
	}

	if err = tx.Commit(); err != nil {
		log.Fatal(err)
	}

	return db
}
