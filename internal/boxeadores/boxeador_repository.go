package boxeadores

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

var Boxeadores []Boxeador

func CargaBoxeadores() {
	jsonFile, err := os.Open("boxeadores.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	byteValue, _ := io.ReadAll(jsonFile)

	json.Unmarshal(byteValue, &boxeadores)
}
