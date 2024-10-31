package main

import (
	"example/cddownload/internal/championlist"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	names := championlist.GetChampions()

	for _, n := range names {
		formattedName := strings.ToLower((n.Name))
		fileName := fmt.Sprintf("https://raw.communitydragon.org/14.21/game/data/characters/%v/%v.bin.json", formattedName, formattedName)
		downloadFile(fmt.Sprintf("champions/%v.json", formattedName), fileName)
	}
}

func downloadFile(ouputFileName string, filePath string) (err error) {
	out, err := os.Create(ouputFileName)
	if err != nil {
		return err
	}

	defer out.Close()

	resp, err := http.Get(filePath)
	defer resp.Body.Close()

	o, err := io.Copy(out, resp.Body)
	_ = o

	fmt.Println("This is my first program in Go")

	return nil
}
