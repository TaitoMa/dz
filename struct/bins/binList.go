package bins

import (
	"encoding/json"
	"github.com/fatih/color"
	"struct/files"
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
	data, err := files.ReadFile("binList.json")
	if err != nil {
		color.Red(err.Error(), "ReadBinListFromFile")
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
	files.WriteFile(byteBinList, "binList.json")
}
