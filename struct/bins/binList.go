package bins

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
