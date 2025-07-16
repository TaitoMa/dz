package main

import (
	"encoding/json"
	"fmt"
	"struct/bins"
)

func main() {
	bin1 := bins.NewBin("bin1", true, "Gopher1")
	bin2 := bins.NewBin("bin2", false, "Gopher2")
	bin3 := bins.NewBin("bin3", true, "Gopher3")

	binsList := bins.NewBinList([]*bins.Bin{bin1, bin2, bin3})
	fmt.Println(binsList)
	binsList.WriteBinListToFile()
	dataBytes, _ := binsList.ReadBinListFromFile()
	var data bins.BinList
	json.Unmarshal(dataBytes, &data)
	fmt.Println(data)
}
