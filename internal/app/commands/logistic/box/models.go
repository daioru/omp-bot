package box

import "fmt"

type Box struct {
	Title string
}

func (b *Box) String() string {
	return fmt.Sprintf("Package{Title: %s}", b.Title)
}
