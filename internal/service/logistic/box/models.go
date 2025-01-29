package box

import "fmt"

var allEntities = []Box{
	{Title: "one"},
	{Title: "two"},
	{Title: "three"},
	{Title: "four"},
	{Title: "five"},
}

type Box struct {
	Title string
}

func (b *Box) String() string {
	return fmt.Sprintf("Package{Title: %s}", b.Title)
}
