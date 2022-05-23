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

var traceCmd = &cobra.Command{
	Use:   "trace",
	Short: "This was IP traced 🫠",
	Long:  `This was IP traced. by  🤯`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			for _, ip := range args {
				fmt.Println(ip)
				showData()
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

type Ip struct {
	Ip       string `json:"id"`
	City     string `json:"city"`
	Region   string `json:"region"`
	Country  string `json:"counter"`
	Loc      string `json:"loc"`
	Timezone string `json:"timezone"`
	Postal   string `json:"postal"`
}

func showData() {
	url := "http://ipinfo.io/1.1.1.1/geo"
	responseByte := getData(url)

	data := Ip{}

	err := json.Unmarshal(responseByte, &data)
	if err != nil {
		log.Println("Unable to unmarshal the response")
	}

	c := color.New(color.FgRed).Add(color.Underline).Add(color.Bold)
	c.Println("DATA FOUND :")

	fmt.Printf("IP :%s\nCITY :%s\nREGION :%s\nCOUNTRY :%s\nLOC :%s\nTIMEZONE :%s\nPOSTAL :%s\n", data.Ip, data.City, data.Region, data.Country, data.Loc, data.Timezone, data.Postal)
}

func getData(url string) []byte {

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
	rootCmd.AddCommand(traceCmd)
}
