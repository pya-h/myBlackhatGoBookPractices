package main

import (
	"shodan"
	"fmt"
	"log"
	"os"
)

func main() {
	defer func() {
		err := recover()
		if err != nil {
			log.Fatalln(err)
		}
	} ()
	if(len(os.Args) != 2) {
		panic("Insufficient arguments!")
	}
	apiKey := os.Getenv("SHODAN_API_KEY")  // must define it in your env variables
	//const apiKey string = "NVExoHk2fNJo37lJiHguFc8KHMTSAE2f"

	shodan_client := shodan.New(apiKey)

	if info, err := shodan_client.GetApiInfo(); err == nil {
		log.Printf("Query Credits: %d\nScan Credits: %d\n\n", info.QueryCredits, info.ScanCredits)

		if hostSearch, err := shodan_client.HostSearch(os.Args[1]); err == nil {
			for _, host := range hostSearch.Matches {
				log.Printf("%18s%8d\n", host.IPString, host.Port)
			}
		} else {
			panic(fmt.Sprint("Error while doing host search:", err))
		}
	} else {
		panic(fmt.Sprint("Error while retreiving Api Info:", err))
	}

}