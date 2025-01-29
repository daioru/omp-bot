package box

import "fmt"

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
	return &DummyBoxService{boxModel: *NewDummyBoxModel()}
}

func (s *DummyBoxService) List(cursor int, limit int) ([]Box, error) {
	boxSlice := []Box{}

	for key := range s.boxModel.Entities {
		boxSlice = append(boxSlice, s.boxModel.Entities[key])
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
	return true, nil
}

func (s *DummyBoxService) Create(box Box) (int, error) {
	_, ok := s.boxModel.Entities[box.ID]
	if ok {
		return 0, fmt.Errorf("Box with id %d already exists", box.ID)
	}
	s.boxModel.Entities[box.ID] = box
	return box.ID, nil
}
