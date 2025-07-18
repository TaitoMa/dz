package bins

import (
	"encoding/json"
	"github.com/fatih/color"
)

type Db interface {
	Read() ([]byte, error)
	Write([]byte) error
}
type BinList struct {
	Bins []*Bin `json:"bins"`
}

type BinListWithDb struct {
	BinList
	db Db
}

func NewBinList(db Db, bins []*Bin) *BinListWithDb {
	list := make([]*Bin, len(bins))
	for i, bin := range bins {
		list[i] = bin
	}
	return &BinListWithDb{
		BinList: BinList{
			Bins: list,
		},
		db: db}
}

func (list *BinList) BinListToBytes() ([]byte, error) {
	data, err := json.Marshal(list)
	if err != nil {
		color.Red(err.Error(), "BinListToBytes")
		return nil, err
	}
	return data, nil
}

func (list *BinListWithDb) ReadBinListFromFile() ([]byte, error) {
	data, err := list.db.Read()
	//data, err := storage.ReadFile()
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (list *BinListWithDb) WriteBinListToFile() {
	byteBinList, err := list.BinListToBytes()
	if err != nil {
		color.Red(err.Error(), "WriteBinListToFile")
		return
	}
	err = list.db.Write(byteBinList)
	if err != nil {
		color.Red(err.Error(), "WriteBinListToFile")
		return
	}
	//storage.WriteFile(byteBinList, list.db.Filename)
}
