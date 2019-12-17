package tree

// Item represent ...
type Item interface {
	String() string
	Less(comp Item) bool
}

// Items is the list of Item.
type Items []Item

func (ii Items) Len() int {
	return len(ii)
}

func (ii Items) Less(i, j int) bool {
	return ii[i].Less(ii[j])
}

func (ii Items) Swap(i, j int) {
	tmp := ii[j]
	ii[j] = ii[i]
	ii[i] = tmp
	return
}
