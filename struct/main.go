package main

import (
	"github.com/joho/godotenv"
	"log"
	"struct/api"
	"struct/bins"
	"struct/files"
	"struct/storage"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	bin1 := bins.NewBin("QQQQQQ", true, "Gopher11")
	bin2 := bins.NewBin("bin2", false, "Gopher22")
	bin3 := bins.NewBin("bin3", true, "Gopher33333")

	db := files.NewJsonDb("binListDb.json")
	db1 := storage.NewStorage("storageDb.json")

	bin4 := bins.NewBin("db2", false, "Gopher22")
	bin5 := bins.NewBin("db22", true, "Gopher33333")

	binsList := bins.NewBinList(db, []*bins.Bin{bin1, bin2, bin3})
	binsList2 := bins.NewBinList(db1, []*bins.Bin{bin4, bin5})
	binsList.WriteBinListToFile()
	binsList2.WriteBinListToFile()
	api.SomeApiFunc()
	//dataBytes, _ := binsList.ReadBinListFromFile()
	//var data bins.BinList
	//json.Unmarshal(dataBytes, &data)
	//fmt.Println(data)
}
