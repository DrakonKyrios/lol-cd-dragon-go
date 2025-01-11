package main

import (
	"encoding/json"
	"example/cddownload/internal/championlist"
	"example/cddownload/internal/ddragon"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

var version_directory string

func main() {	
	if( len(os.Args[1:]) <= 0){
		panic("Missing Version Argument ")
	}

	version_directory := os.Args[1]

	os.Mkdir(version_directory, 0777)

	var part string
	if( len(os.Args[1:]) > 1){
		part = os.Args[2]
	}

	if len(part) == 0 || part == "a" {
		importCommunityItemFile(version_directory)
	}
	
	if len(part) == 0 || part == "b" {
		importChampionStats(version_directory)
	}

	if len(part) == 0 || part == "c" {
		importItemFile(version_directory)
	}

	if len(part) == 0 || part == "champions" {
		importChampions(version_directory)
	}
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
	AffectsTypeFlags int `json:"mAffectsTypeFlags"`
	EffectAmount json.RawMessage `json:"mEffectAmount"`
	SpellTags json.RawMessage `json:"mSpellTags"`
}

func importChampionStats(version_directory string) {
	data := ddragon.GetChampionsByVersion(version_directory)

	out, err := os.Create(fmt.Sprintf("%v/_championsStats.json", version_directory))
	if err != nil {
		panic(err)
	}

	defer out.Close()

	o, _ := out.Write(data)
	_ = o

	defer out.Close()	
}

func importCommunityItemFile(version_directory string){
	fmt.Println("Community Item Download Start")
	data := championlist.GetCommunityItemsByVersion(version_directory)

	out, err := os.Create(fmt.Sprintf("%v/_community_items.json", version_directory))
	if err != nil {
		panic(err)
	}

	defer out.Close()

	o, _ := out.Write(data)
	_ = o

	defer out.Close()
	fmt.Println("Community Item Download Finished")
}

func importItemFile(version_directory string) {
	fmt.Println("DDragon Item Download Start")
	data := ddragon.GetItemsByVersion(version_directory)

	out, err := os.Create(fmt.Sprintf("%v/_items.json", version_directory))
	if err != nil {
		panic(err)
	}

	defer out.Close()

	o, _ := out.Write(data)
	_ = o

	defer out.Close()
	fmt.Println("DDragon Item Download Finished")
}

func importChampions(version string) {	
	names := championlist.GetChampions(version)	
	out, err := os.Create(fmt.Sprintf("%v/_champions.json", version))
	if err != nil {
		panic(err)
	}

	defer out.Close()
	championFull := make(map[string]json.RawMessage)	
	for _, n := range names {		
		formattedName := strings.ToLower((n.Name))
		fileName := fmt.Sprintf("https://raw.communitydragon.org/%v/game/data/characters/%v/%v.bin.json", version, formattedName, formattedName)
		data := downloadFile(formattedName, fmt.Sprintf("%v/%v.json", version, n.Name),  fileName)
		
		championFull[formattedName] = data

		fmt.Printf("%#v has been added \n", n.Name)
		
		time.Sleep(4 * time.Second)
	}

	s, _ := json.MarshalIndent(championFull, "", "\t")

	o, _ := out.Write(s)
	_ = o
	
}

func downloadFile(championName string, outputFileName string, filePath string) (data json.RawMessage) {

	resp, championErr := http.Get(filePath)

	if championErr != nil {
		panic(championErr)
	}
	
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

	o, _ := out.Write(as_json)
	_ = o

	return as_json
}
