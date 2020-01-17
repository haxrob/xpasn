/*
xpasn - Expand an AS number to prefixes or IP addresses
Author: @x1sec / robert@x1sec.com
*/

package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
)

type data struct {
	Status        string `json:"status"`
	StatusMessage string `json:"status_message"`
	Data          struct {
		Ipv4Prefixes []struct {
			Prefix string `json:"prefix"`
		} `json:"ipv4_prefixes"`
	} `json:"data"`
}

func usage() {
	fmt.Fprintf(os.Stderr, "Usage of %s <asn>:\n", os.Args[0])
	flag.PrintDefaults()
}

func main() {

	var expand bool
	flag.BoolVar(&expand, "e", false, "Expand subnets to IPs (optional)")

	flag.Usage = usage
	flag.Parse()

	var asn string

	lo := log.New(os.Stderr, "", 0)

	if len(flag.Args()) == 1 {
		asn = flag.Args()[0]
	} else {
		lo.Println("You must specify an AS number")
		os.Exit(1)
	}

	url := fmt.Sprintf("https://api.bgpview.io/asn/%s/prefixes", asn)

	resp, err := http.Get(url)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		lo.Println("Error accessing bgpview API", 0)
		os.Exit(1)
	}

	var apiResult data
	if err := json.Unmarshal(body, &apiResult); err != nil {
		lo.Println("Error parsing bgpview response")
	}

	for _, prefix := range apiResult.Data.Ipv4Prefixes {
		if expand == true {
			for _, ip := range netExpand(prefix.Prefix) {
				fmt.Println(ip)
			}
		} else {
			fmt.Println(prefix.Prefix)
		}
	}
}

func netExpand(network string) []string {
	var ips []string
	ip, ipnet, err := net.ParseCIDR(network)
	if err != nil {
		panic(err)
	}
	for ip := ip.Mask(ipnet.Mask); ipnet.Contains(ip); incIP(ip) {
		ips = append(ips, ip.String())
	}

	return ips[1 : len(ips)-1]
}

func incIP(ip net.IP) {
	for j := len(ip) - 1; j >= 0; j-- {
		ip[j]++
		if ip[j] > 0 {
			break
		}
	}
}
