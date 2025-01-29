package box

import "fmt"

var Entities map[int]Box

type DummyBoxModel struct {
	Entities map[int]Box
	Slice []BoxStruct
}

type BoxStruct struct {
	ID int
	Box Box
}

type Box struct {
	ID        int
	Weight    float32
	Volume    float32
	IsFragile bool
}

func (b *Box) String() string {
	return fmt.Sprintf("Box[%d] weigth: %.2f, volume %.2f, fragile: %t}",
		b.ID,
		b.Weight,
		b.Volume,
		b.IsFragile)
}

func (d *DummyBoxModel) NewBox(id int, weight, volume float32, isFragile bool) *Box {
	return &Box{
		ID:        id,
		Weight:    weight,
		Volume:    volume,
		IsFragile: isFragile,
	}
}

func NewDummyBoxModel() *DummyBoxModel {
	return &DummyBoxModel{Entities: make(map[int]Box)}
}
