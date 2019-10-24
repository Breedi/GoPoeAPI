package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

)

type Category struct {
	Accessories []string `json:"accessories"`
	Armour []string `json:"armour"`
	Jewels []string `json:"jewels"`
	Weapons []string `json:"weapons"`
}

type Sockets struct {
	Group int `json:"group"`
	Attr string `json:"attr"`
	sColour string `json:"s_colour"`
}

type ValueTypes struct {
	key []int `json:"key"`
	ValueDescr string `json:"value_descr"`
}

type FrameType struct {
	key int `json:"key"`
	Value string `json:"value"`
}

type Properties struct {
	Name string `json:"name"`
	Values [][]ValueTypes `json:"valuesTypes"`
	DisplayMode int `json:"displayMode"`
	Type int `json:"type"`
	Progress int `json:"progress"`
}

type Requirements struct {
	Name string `json:"name"`
	Values [][]ValueTypes `json:"valuesTypes"`
	DisplayMode int `json:"display_mode"`
	Type int `json:"type"`
	Progress int `json:"progress"`
}

type Items struct {
	AbyssJewel bool `json:"abyssJewel"`
	AdditionalProperties []Properties `json:"additional_properties"`
	ArtFilename string `json:"artFilename"`
	Category []Category `json:"category"`
	Corrupted bool `json:"corrupted"`
	CosmeticMods []string `json:"cosmeticMods"`
	CraftedMods []string `json:"craftedMods"`
	DescText string `json:"descText"`
	Duplicated bool `json:"duplicated"`
	Elder bool `json:"elder"`
	EnchantMods []string `json:"enchantMods"`
	ExplicitMods []string `json:"explicitMods"`
	FlavourText []string `json:"flavourText"`
	FrameType int `json:"frameType"`
	H int `json:"h"`
	Icon string `json:"icon"`
	Id string `json:"id"`
	Identified bool `json:"identified"`
	Ilvl int `json:"ilvl"`
	implicitMods []string `json:"implicitMods"`
	InventoryID string `json:"inventoryId"`
	IsRelic bool `json:"isRelic"`
	League string `json:"league"`
	LockedToCharacter bool `json:"lockedToCharacter"`
	MaxStackSize int `json:"maxStackSize"`
	Name string `json:"name"`
	NextLevelRequirements []Requirements `json:"nextLevelRequirements"`
	Note string `json:"note"`
	Properties []Properties `json:"properties"`
	ProphecyDiffText string `json:"prophecyDiffText"`
	ProphecyText string `json:"prophecyText"`
	Requirements []Requirements `json:"requirements"`
	SecDescrText string `json:"secDescrText"`
	Shaper bool `json:"shaper"`
	SocketedItems []Items `json:"socketedItems"`
	Sockets []Sockets `json:"sockets"`
	StackSize int `json:"stackSize"`
	Support bool `json:"support"`
	TalismanTier int `json:"talisman_tier"`
	TypeLine string `json:"typeine"`
	UtilityMods []string `json:"utilityMods"`
	Verified bool `json:"verified"`
	W int `json:"w"`
	X int `json:"x"`
	Y int `json:"y"`
}

type Stashes struct {
	AccountName string `json:"accountName"`
	LastCharacterName string `json:"lastCharacterName"`
	Id string `json:"id"`
	Stash string `json:"stash"`
	StashType string `json:"stashType"`
	Items []Items `json:"items"`
	Public bool `json:"public"`
}

type TradeData struct {
	NextChangeID string `json:"next_change_id"`
	Stashes []Stashes `json:"stashes"`
}

const  (
	fetchUrl = "http://api.pathofexile.com/public-stash-tabs/0"
)


func main() {
	fmt.Printf("Starting Application ...\n")
	fmt.Printf("Trying to fetch data from: %s \n", fetchUrl)
	response, err := http.Get(fetchUrl)
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		fmt.Printf("Successful...\n")
		data, _ := ioutil.ReadAll(response.Body)
		//fmt.Println(string(data))

		fmt.Println(json.Valid(data))


		var currentData TradeData
		err := json.Unmarshal(data, &currentData)
		if err != nil {
			fmt.Println("errors: ", err)
		}

		fmt.Printf("NextChange-ID: %+v \n", currentData.NextChangeID)
		fmt.Printf("%+v \n", currentData.Stashes[0].Id)
		fmt.Printf("%+v \n", currentData.Stashes[0].LastCharacterName)
		fmt.Println(len(currentData.Stashes))


	}


}