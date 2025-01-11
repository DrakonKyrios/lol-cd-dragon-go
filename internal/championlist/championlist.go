package championlist

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type championName struct {
	Name string `json:"name"`
}

type championList struct {
	Ahri championName `json:"Characters/Ahri"`
}

type itemList struct {

}

func GetChampions(version string) (names map[string]championName) {
	filePath := fmt.Sprintf("https://raw.communitydragon.org/%v/game/global/champions/champions.bin.json", version)

	resp, err := http.Get(filePath)

	if err != nil {
		panic(err.Error())
	}

	defer resp.Body.Close()
	_ = err

	container := make(map[string]championName)

	b, _ := io.ReadAll(resp.Body)

	e := json.Unmarshal(b, &container)
	_ = e

	return container
}

func GetCommunityItemsByVersion(version_directory string) (date []byte) {
	filePath := fmt.Sprintf("https://raw.communitydragon.org/%v/game/items.cdtb.bin.json", version_directory)
	response, err := http.Get(filePath)

	if err != nil {
		panic(err.Error())
	}

	defer response.Body.Close()		
	
	b, _ := io.ReadAll(response.Body)

	return b
}