package tree

type item string

func (i item) Less(comp Item) bool {
	return string(i) < string(comp.(item))
}

func (i item) String() string {
	return string(i)
}
