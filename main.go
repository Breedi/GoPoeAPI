package main

import (
	"encoding/json"
	"fmt"
	"github.com/Breedi/GoPoeTrade/models"
	"github.com/Breedi/GoPoeTrade/views"
	"github.com/labstack/echo"
	"io/ioutil"
	"net/http"
	"strconv"
)

const (
	fetchUrl = "http://api.pathofexile.com/public-stash-tabs/"
)

func ShowStashInfo(stashes models.TradeData) {
	for index, data := range stashes.Stashes {
		if len(data.Items) != 0 {
			fmt.Println("===================| STASH-NR: " + strconv.Itoa(index) + " |===================")
			fmt.Println(data.LastCharacterName + "\n" + data.Id + "n" + data.AccountName + "\n")
			fmt.Println("Stash-Name: \t + " + data.Stash + "\t Stash-Type: " + data.StashType)
			if data.Public == true {
				fmt.Println("This Stash is public!")
			} else {
				fmt.Println("This Stash isn't public!")
			}
			fmt.Println("Contains " + strconv.Itoa(len(data.Items)) + " Items inside this Stash")
		} else {
			continue
		}

	}
}

func FetchDataFromAPI(stashes []models.TradeData, nextChangeID string, fetchCounter *int) []models.TradeData {

	var currentData models.TradeData
	fmt.Printf("Starting Application ...\n")
	fmt.Printf("Trying to fetch data from: %s \n", fetchUrl+nextChangeID)
	response, err := http.Get(fetchUrl + nextChangeID)
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		fmt.Printf("Successful...\n")
		data, _ := ioutil.ReadAll(response.Body)
		//fmt.Println(string(data))
		fmt.Println(json.Valid(data))

		err := json.Unmarshal(data, &currentData)
		if err != nil {
			fmt.Println("errors: ", err)
		}
		ShowStashInfo(currentData)

		stashes = append(stashes, currentData)
		*fetchCounter += 1

		fmt.Printf("NextChange-ID: %+v \n", currentData.NextChangeID)
		fmt.Printf("%+v \n", currentData.Stashes[0].Id)
		fmt.Printf("%+v \n", currentData.Stashes[0].LastCharacterName)
		fmt.Println(len(currentData.Stashes))

		if *fetchCounter >= 1 {
			return stashes
		} else {
			FetchDataFromAPI(stashes, currentData.NextChangeID, fetchCounter)
		}

	}
	return stashes
}

func main() {

	var PoEAPIData []models.TradeData
	var fetchCounter int

	PoEAPIData = FetchDataFromAPI(PoEAPIData, "", &fetchCounter)
	fmt.Println("Successfully fetched from API Endpoint: " + strconv.Itoa(fetchCounter) + " times")

	e := echo.New()

	e.GET("/", views.Index)

	e.GET("/api/stashes", func(context echo.Context) error {
		b, err := json.Marshal(PoEAPIData)
		if err != nil {
			fmt.Println(err)
		}
		return context.String(http.StatusOK, string(b))
	})

	e.Logger.Fatal(e.Start(":1234"))

	//beego.Run()
}
