package main

import (
	"fmt"
	"time"
)

type BinList struct {
	bins []*Bin
}
type Bin struct {
	id        string
	private   bool
	createdAt time.Time
	name      string
}

func newBinList(bins []*Bin) *BinList {
	list := make([]*Bin, len(bins))
	for i, bin := range bins {
		list[i] = bin
	}
	return &BinList{bins: list}
}

func newBin(id string, private bool, name string) *Bin {
	return &Bin{
		id:        id,
		private:   private,
		createdAt: time.Now(),
		name:      name,
	}
}

func main() {
	bin1 := newBin("bin1", true, "Gopher1")
	bin2 := newBin("bin2", false, "Gopher2")
	bin3 := newBin("bin3", true, "Gopher3")

	bins := newBinList([]*Bin{bin1, bin2, bin3})

	fmt.Println(bins)
}
