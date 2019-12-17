package tree

// Item represent 
type Item interface {
	String() string
	Less(comp Item) bool
}
