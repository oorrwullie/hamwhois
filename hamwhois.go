package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/mitchellh/mapstructure"
)

type OpInfo struct {
	Call    string `json:"call"`
	Class   string `json:"class"`
	Expires string `json:"expires"`
	Status  string `json:"status"`
	Grid    string `json:"grid"`
	Lat     string `json:"lat"`
	Lon     string `json:"lon"`
	Fname   string `json:"fname"`
	Mi      string `json:"mi"`
	Name    string `json:"name"`
	Suffix  string `json:"suffix"`
	Addr1   string `json:"addr1"`
	Addr2   string `json:"addr2"`
	State   string `json:"state"`
	Zip     string `json:"zip"`
	Country string `json:"country"`
}

// GetOpInfo fetch operator info from hamdb.org
func main() {
	if len(os.Args) == 1 {
		fmt.Println("\nUSAGE: hamwhois {callsign}\n")
		os.Exit(0)
	}

	callsign := os.Args[1]
	resp, err := http.Get(fmt.Sprintf("https://api.hamdb.org/v1/%s/json/hamdb", callsign))
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var res map[string]interface{}
	var opInfo OpInfo
	json.Unmarshal(body, &res)

	status := res["hamdb"].(map[string]interface{})["messages"].(map[string]interface{})["status"].(string)

	if status != "OK" {
		fmt.Println(status)
		os.Exit(0)
	}

	mapstructure.Decode(res["hamdb"].(map[string]interface{})["callsign"].(map[string]interface{}), &opInfo)

	opJSON, _ := json.MarshalIndent(opInfo, "", "  ")
	fmt.Println(string(opJSON))
}
