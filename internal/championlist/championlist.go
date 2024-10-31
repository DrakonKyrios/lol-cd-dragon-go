package championlist

import (
	"encoding/json"
	"io"
	"net/http"
)

type championName struct {
	Name string `json:"name"`
}
type championList struct {
	Ahri championName `json:"Characters/Ahri"`
}

func GetChampions() (names map[string]championName) {
	resp, err := http.Get("https://raw.communitydragon.org/14.21/game/global/champions/champions.bin.json")

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
