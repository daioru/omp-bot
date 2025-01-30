package box

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var Entities map[int]Box

type DummyBoxModel struct {
	Entities map[int]Box
	IDs      []int
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

func (d *DummyBoxModel) NewBoxFromArgs(args string) (*Box, error) {
	argsParts := strings.SplitN(args, " ", 4)
	if len(argsParts) != 4 {
		return &Box{}, errors.New("invalid amount of arguments")
	}

	id, err := strconv.Atoi(argsParts[0])
	if err != nil {
		return &Box{}, errors.New("invalid ID format")
	}

	weight, err := strconv.ParseFloat(argsParts[1], 32)
	if err != nil {
		return &Box{}, errors.New("invalid weight format")
	}

	volume, err := strconv.ParseFloat(argsParts[2], 32)
	if err != nil {
		return &Box{}, errors.New("invalid volume format")
	}

	isFragile, err := strconv.ParseBool(argsParts[3])
	if err != nil {
		return &Box{}, errors.New("invalid fragility format")
	}

	return NewBox(id, float32(weight), float32(volume), isFragile), nil
}

func NewBox(id int, weight, volume float32, isFragile bool) *Box {
	return &Box{
		ID:        id,
		Weight:    weight,
		Volume:    volume,
		IsFragile: isFragile,
	}
}

func NewDummyBoxModel() *DummyBoxModel {
	return &DummyBoxModel{
		Entities: make(map[int]Box),
		IDs:      []int{},
	}
}
