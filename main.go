package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

var anomalyMapType = map[string]string{
	"P": "Pothole",
	"S": "Speed Bump",
	"R": "Rough Road",
	"p": "Pothole",
	"s": "Speed Bump",
	"r": "Rough Road",
}

var dataFile *os.File

func init() {
	file, err := os.OpenFile("data.csv", os.O_APPEND|os.O_RDONLY|os.O_WRONLY, 0777)
	if err != nil {
		log.Fatalln("Could not open file, error - ", err.Error())
	}

	dataFile = file

}

func main() {

	for {
		var anomaly string
		fmt.Println("Choose the upcoming anomaly: (P for Pothole, S for Speed Bump, R for Rough Road)")
		_, err := fmt.Scan(&anomaly)
		if err != nil {
			fmt.Println("error at anomaly type: ", err.Error())
			continue
		}

		if _, ok := anomalyMapType[anomaly]; !ok {
			fmt.Println("invalid entry, you'll have to start again")
			continue
		}

		var startOption int
		fmt.Println("Enter 0 to Abort or 1 to Start:")
		_, err = fmt.Scan(&startOption)
		if err != nil {
			fmt.Println("error at option: ", err.Error())
			continue
		}

		if startOption != 1 {
			fmt.Println("aborting, you'll have to start again")
			continue
		}

		startTime := time.Now().Format("2006-01-02 15:04:05")

		var endOption int
		fmt.Println("Enter 0 to Abort or 1 to End:")
		_, err = fmt.Scan(&endOption)
		if err != nil {
			fmt.Println("error at option: ", err.Error())
			continue
		}

		if endOption != 1 {
			fmt.Println("aborting, you'll have to start again")
			continue
		}

		endTime := time.Now().Format("2006-01-02 15:04:05")

		_, err = dataFile.Write([]byte(startTime + "," + anomalyMapType[anomaly] + "," + endTime))
		if err != nil {
			fmt.Println("Could not log to file, error - ", err.Error())
		}
		fmt.Println("Successful!")
	}
}
