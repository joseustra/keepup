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
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/boltdb/bolt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/ustrajunior/keepup/cfgo"
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Updates the dns record with current ip",
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		var currentIP string
		domain := &Domain{}

		if len(ip) > 7 {
			currentIP = ip
		} else {
			currentIP, err = cfgo.GetIPV4IP()
			if err != nil {
				log.Fatal(err)
			}
		}

		if !force {
			db.View(func(tx *bolt.Tx) error {
				b := tx.Bucket([]byte("domains"))
				v := b.Get([]byte(fmt.Sprintf("%s-%s", zone, dnsRecord)))

				if len(v) > 0 {
					json.Unmarshal(v, domain)
				}
				return nil
			})

			if len(domain.IP) >= 7 {
				if domain.IP == currentIP {
					fmt.Println("current and old ip are the same, no need to update")
					return
				}
			}
		}

		client, err := cfgo.NewClient(viper.GetString("cfKey"), viper.GetString("cfEmail"))
		if err != nil {
			log.Fatal(err)
		}

		record, err := client.GetDNSRecord(zone, dnsRecord)
		if err != nil {
			log.Fatal(err)
		}

		if record.Content != currentIP {
			record.Content = strings.TrimSpace(currentIP)
			err = client.UpdateDNSRecord(record)
			if err != nil {
				log.Fatal(err)
			}
		}

		domain = &Domain{
			zone,
			record.Name,
			record.Content,
		}

		fmt.Printf("DNS updated to: %s %s\n", domain.DNS, domain.IP)

		err = db.Update(func(tx *bolt.Tx) error {
			b := tx.Bucket([]byte("domains"))
			j, _ := json.Marshal(domain)
			err := b.Put([]byte(fmt.Sprintf("%s-%s", zone, domain.DNS)), j)
			return err
		})

		if err != nil {
			panic(err)
		}
	},
}

func init() {
	RootCmd.AddCommand(updateCmd)
}
