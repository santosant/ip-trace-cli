package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var hostnameCmd = &cobra.Command{
	Use:   "hostname",
	Short: "IP Hostname",
	Long:  `IP Hostname`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			for _, ip := range args {
				showHostname(ip)
			}
		} else {
			fmt.Println("Please provide IP to trace.")
		}
	},
}

// {
//   "ip": "1.1.1.1",
//   "hostname": "one.one.one.one",
//   "anycast": true,
//   "city": "Los Angeles",
//   "region": "California",
//   "country": "US",
//   "loc": "34.0522,-118.2437",
//   "org": "AS13335 Cloudflare, Inc.",
//   "postal": "90076",
//   "timezone": "America/Los_Angeles",
//   "readme": "https://ipinfo.io/missingauth"
// }

type IpHostname struct {
	Hostname string `json:"hostname"`
}

func showHostname(ip string) {
	url := "http://ipinfo.io/" + ip + "/geo"
	responseByte := getData(url)

	data := IpHostname{}

	err := json.Unmarshal(responseByte, &data)
	if err != nil {
		log.Println("Unable to unmarshal the response")
	}

	c := color.New(color.FgRed).Add(color.Underline).Add(color.Bold)
	c.Println("DATA FOUND :")

	fmt.Printf("HOSTANAME :%s\n", data.Hostname)
}

func getHostname(url string) []byte {

	response, err := http.Get(url)
	if err != nil {
		log.Println("Unable to get the response")
	}

	responseByte, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println("Unable to read the response.")
	}

	return responseByte

}
func init() {
	rootCmd.AddCommand(hostnameCmd)
}
