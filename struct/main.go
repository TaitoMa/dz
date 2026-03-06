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
	//isCreate := flag.Bool("create", false, "create new bin")
	//isUpdate := flag.Bool("update", false, "update bin")
	//isDelete := flag.Bool("delete", false, "delete bin")
	//isGet := flag.Bool("get", false, "get bin")
	//isList := flag.Bool("list", false, "list bin")
	//file := flag.String("file", "", "file bin")
	//name := flag.String("name", "", "name bin")
	//id := flag.String("id", "", "id bin")
	api.SomeApiFunc()

	fillSomeBins()
}

func fillSomeBins() {
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
}
