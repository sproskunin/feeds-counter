package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
)

type Offer struct {
	Id        string `xml:"id,attr"`
	Available string `xml:"available,attr"`
}

type XmlOffers struct {
	XMLName xml.Name `xml:"offers"`
	Offers  []Offer  `xml:"offer"`
}

type XmlShop struct {
	XMLName xml.Name  `xml:"shop"`
	Offers  XmlOffers `xml:"offers"`
}

type XmlCatalog struct {
	XMLName xml.Name `xml:"yml_catalog"`
	Shop    XmlShop  `xml:"shop"`
}

func exampleReadDir() {
	entries, err := ioutil.ReadDir("/home/proskunin/data/remote")
	if err != nil {
		log.Panicf("failed reading directory: %s", err)
	}
	fmt.Printf("\nNumber of files in current directory: %d", len(entries))
}

func exampleReadFile() {
	data, err := ioutil.ReadFile("/home/proskunin/data/remote/feed-auto-goods-yandex.xml")
	if err != nil {
		log.Panicf("failed reading data from file: %s", err)
	}
	catalog := XmlCatalog{}
	err = xml.Unmarshal(data, &catalog)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println("\nOffers count: ", len(catalog.Shop.Offers.Offers))

	fmt.Printf("\nLength: %d bytes", len(data))
	fmt.Printf("\nError: %v", err)
}

func main() {

	fmt.Printf("\nКоличество файлов в дирректори")
	exampleReadDir()

	fmt.Printf("\nРазмер файла в байтах и количество итемов в фиде")
	exampleReadFile()
}
