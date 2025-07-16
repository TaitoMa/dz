package bins

import (
	"encoding/json"
	"github.com/fatih/color"
	"struct/storage"
)

type BinList struct {
	bins []*Bin
}

func NewBinList(bins []*Bin) *BinList {
	list := make([]*Bin, len(bins))
	for i, bin := range bins {
		list[i] = bin
	}
	return &BinList{bins: list}
}

func (list *BinList) BinListToBytes() ([]byte, error) {
	data, err := json.Marshal(list)
	if err != nil {
		color.Red(err.Error(), "BinListToBytes")
		return nil, err
	}
	return data, nil
}

func (list *BinList) ReadBinListFromFile() ([]byte, error) {
	data, err := storage.ReadFile("binList.json")
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (list *BinList) WriteBinListToFile() {
	byteBinList, err := list.BinListToBytes()
	if err != nil {
		color.Red(err.Error(), "WriteBinListToFile")
		return
	}
	storage.SaveFile(byteBinList, "binList.json")
}
