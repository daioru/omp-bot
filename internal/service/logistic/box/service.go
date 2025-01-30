package box

import (
	"errors"
	"fmt"
)

type BoxService interface {
	Describe(boxID int) (Box, error)
	List(cursor int, limit int) ([]Box, error)
	Create(Box) (int, error)
	Update(boxID int, box Box) error
	Remove(boxID int) (bool, error)
}

type DummyBoxService struct {
	boxModel DummyBoxModel
}

func NewDummyBoxService() *DummyBoxService {
	dummy := DummyBoxService{boxModel: *NewDummyBoxModel()}
	dummy.Create(*NewBox(1, 1, 1, false))
	dummy.Create(*NewBox(2, 1, 1, false))
	dummy.Create(*NewBox(3, 1, 1, false))
	dummy.Create(*NewBox(4, 1, 1, false))
	dummy.Create(*NewBox(5, 1, 1, false))
	dummy.Create(*NewBox(6, 1, 1, false))
	dummy.Create(*NewBox(7, 1, 1, false))
	dummy.Create(*NewBox(8, 1, 1, false))
	dummy.Create(*NewBox(9, 1, 1, false))
	dummy.Create(*NewBox(10, 1, 1, false))
	dummy.Create(*NewBox(11, 1, 1, false))
	dummy.Create(*NewBox(12, 1, 1, false))
	dummy.Create(*NewBox(13, 1, 1, false))
	dummy.Create(*NewBox(14, 1, 1, false))
	dummy.Create(*NewBox(15, 1, 1, false))
	dummy.Create(*NewBox(16, 1, 1, false))
	dummy.Create(*NewBox(17, 1, 1, false))
	return &dummy
}

func (s *DummyBoxService) List(cursor int, limit int) ([]Box, error) {
	boxSlice := []Box{}

	if cursor < 0 || cursor > len(s.boxModel.IDs) {
		return boxSlice, errors.New("invalid cursor value")
	}

	end := cursor + limit
	if end > len(s.boxModel.IDs) {
		end = len(s.boxModel.IDs)
	}

	for _, id := range s.boxModel.IDs[cursor:end] {
		boxSlice = append(boxSlice, s.boxModel.Entities[id])
	}

	return boxSlice, nil
}

func (s *DummyBoxService) Describe(boxID int) (Box, error) {
	box, ok := s.boxModel.Entities[boxID]
	if !ok {
		return Box{}, fmt.Errorf("Box with id %d doesn't exist", boxID)
	}
	return box, nil
}

func (s *DummyBoxService) Update(boxID int, box Box) error {
	_, ok := s.boxModel.Entities[boxID]
	if !ok {
		return fmt.Errorf("Box with id %d doesn't exist", boxID)
	}

	s.boxModel.Entities[boxID] = box
	return nil
}

func (s *DummyBoxService) Remove(boxID int) (bool, error) {
	_, ok := s.boxModel.Entities[boxID]
	if !ok {
		return false, nil
	}

	delete(s.boxModel.Entities, boxID)
	for i, v := range s.boxModel.IDs {
		if v == boxID {
			s.boxModel.IDs = append(s.boxModel.IDs[:i], s.boxModel.IDs[i+1:]...)
			break
		}
	}

	return true, nil
}

func (s *DummyBoxService) Create(box Box) (int, error) {
	_, ok := s.boxModel.Entities[box.ID]
	if ok {
		return 0, fmt.Errorf("Box with id %d already exists", box.ID)
	}
	s.boxModel.Entities[box.ID] = box
	s.boxModel.IDs = append(s.boxModel.IDs, box.ID)
	return box.ID, nil
}
