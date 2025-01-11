package ddragon

import (
	"fmt"
	"io"
	"net/http"
)

func GetChampionsByVersion(version_directory string) (data []byte) {
	filename := fmt.Sprintf("https://ddragon.leagueoflegends.com/cdn/%v.1/data/en_US/championFull.json", version_directory)
	response, err := http.Get(filename)

	if err != nil {
		panic(err.Error())
	}

	defer response.Body.Close()		
	
	b, _ := io.ReadAll(response.Body)

	return b
}

func GetItemsByVersion(version_directory string) (data []byte) {
	filename := fmt.Sprintf("https://ddragon.leagueoflegends.com/cdn/%v.1/data/en_US/item.json", version_directory)
	response, err := http.Get(filename)

	if err != nil {
		panic(err.Error())
	}

	defer response.Body.Close()		
	
	b, _ := io.ReadAll(response.Body)

	return b
}
