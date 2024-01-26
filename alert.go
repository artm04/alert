package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

var (
	// API
	api = "https://siren.pp.ua/api/v3/alerts"

	// Information
	help_eng = ` 
Usage: alert [EMPRY or OPTION]

--english or -en: translation of output into English

--help or -h:     help 
--version or -v:  version

Air raid alert in Ukraine CLI: <https://github.com/rendick/alert/>
`
	version = "0.1v Alpha"

	// Color
	Red   = "\033[31m"
	Bold  = "\033[1m"
	Reset = "\033[0m"
)

type Alert struct {
	RegionName    string `json:"regionName"`
	LastUpdate    string `json:"lastUpdate"`
	RegionEngName string `json:"regionEngName"`
	RegionId      string `json:"regionId"`
}

func main() {
	res, err := http.Get(api)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	if len(os.Args) != 2 {
		body, err := io.ReadAll(res.Body)
		if err != nil {
			log.Fatal("Помилка читання JSON файлу: ", err)
		}

		var alerts []Alert
		err = json.Unmarshal(body, &alerts)
		if err != nil {
			log.Fatal("Помилка перетворення даних у структуру у JSON файлі: ", err)
			return
		}

		for num, alert := range alerts {
			fmt.Printf("%d."+Red+Bold+" Повітряна тривога:"+Reset+" %s [%s] %s \n",
				num+1,
				alert.RegionName,
				alert.RegionId,
				strings.ReplaceAll(strings.ReplaceAll(alert.LastUpdate, "T", " "), "Z", ""))
		}
		os.Exit(0)
	} else {
		switch os.Args[1] {
		case "--help", "-help", "-h", "--h":
			fmt.Println(help_eng)
		case "--version", "-version", "-v", "--v":
			fmt.Println(version)
		case "--english", "-english", "-eng", "--eng":
			body, err := io.ReadAll(res.Body)
			if err != nil {
				log.Fatal("Error reading JSON file: ", err)
			}

			var alerts []Alert
			err = json.Unmarshal(body, &alerts)
			if err != nil {
				log.Fatal("Error unmarshaling JSON file: ", err)
				return
			}

			for num, alert := range alerts {
				if alert.RegionEngName == "" {
					alert.RegionEngName = "Autonomous Republic of Crimea"
				}
				fmt.Printf("%d."+Red+Bold+" Air raid alert:"+Reset+" %s %s \n",
					num+1,
					alert.RegionEngName,
					strings.ReplaceAll(strings.ReplaceAll(alert.LastUpdate, "T", " "), "Z", ""))
			}
		}
	}
}
