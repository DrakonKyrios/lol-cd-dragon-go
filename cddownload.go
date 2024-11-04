package main

import (
	"encoding/json"
	"example/cddownload/internal/championlist"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	names := championlist.GetChampions()

	out, err := os.Create("champions/_champions.json")
	if err != nil {
		panic(err)
	}

	defer out.Close()
	championFull := make(map[string]json.RawMessage)	
	for _, n := range names {		
		formattedName := strings.ToLower((n.Name))
		fileName := fmt.Sprintf("https://raw.communitydragon.org/14.21/game/data/characters/%v/%v.bin.json", formattedName, formattedName)
		data := downloadFile(formattedName, fmt.Sprintf("champions/%v.json", n.Name), fileName)
		
	
		championFull[formattedName] = data

		fmt.Println(fmt.Sprintf("%#v has been added", n.Name))
		
		time.Sleep(4 * time.Second)
	}

	s, _ := json.MarshalIndent(championFull, "", "\t")

	o, _ := out.Write(s)
	_ = o
}

type Abilities struct {
	Name  string `json:"mScriptName"`
	Spell Spell  `json:"mSpell"`
}

type Spell struct {
	DataValues        json.RawMessage `json:"mDataValues"`
	SpellCalculations json.RawMessage `json:"mSpellCalculations"`
	CastTime          json.RawMessage `json:"mCastTime"`
	Cooldown          json.RawMessage `json:"cooldownTime"`
	Mana              json.RawMessage `json:"mana"`
}

func downloadFile(championName string, outputFileName string, filePath string) (data json.RawMessage) {

	resp, err := http.Get(filePath)
	defer resp.Body.Close()

	ca := make(map[string]Abilities)
	oc := make(map[string]Abilities)

	b, _ := io.ReadAll(resp.Body)

	e := json.Unmarshal(b, &oc)
	_ = e

	for k, info := range oc {
		if strings.Contains(strings.ToLower(k), championName) {
			ca[k] = info
		}
	}

	out, err := os.Create(outputFileName)
	if err != nil {
		panic(err)
	}

	defer out.Close()

	as_json, _ := json.MarshalIndent(ca, "", "\t")

	o, err := out.Write(as_json)
	_ = o

	return as_json
}
