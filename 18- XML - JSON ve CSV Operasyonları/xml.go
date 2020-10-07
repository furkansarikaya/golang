package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	file, err := os.Open("sites.xml")
	if err != nil {
		fmt.Printf("Error : %v", err)
		return
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Printf("Error : %v", err)
		return
	}

	v := ObjectSites{}

	err = xml.Unmarshal(data, &v)
	if err != nil {
		fmt.Printf("Error : %v", err)
		return
	}

	fmt.Println(v)
	fmt.Println(v.Description)

}

//ObjectSites Struct
type ObjectSites struct {
	XMLName     xml.Name `xml:"sites"`
	Version     string   `xml:"version,attr"`
	Sites       []site
	Description string `xml:",innerxml"`
}

type site struct {
	XMLName     xml.Name `xml:"site"`
	Text        string   `xml:",chardata"`
	Name        string   `xml:"Name"`
	Description string   `xml:"Description"`
	Category    string   `xml:"Category"`
}
