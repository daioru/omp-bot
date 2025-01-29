package box

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (s *Service) List() []Box {
	return allEntities
}

func (s *Service) Get(idx int) (*Box, error) {
	return &allEntities[idx], nil
}

func (s *Service) Edit(idx int, title string) (*Box, error) {
	// Поменять на правильную логику
	allEntities[idx] = Box{
		Title: title,
	}

	return &allEntities[idx], nil
}

func (s *Service) Delete(idx int) error {
	// Подумать об обработке ошибок
	allEntities[idx] = allEntities[len(allEntities)-1]
	allEntities[len(allEntities)-1] = Box{}
	allEntities = allEntities[:len(allEntities)-1]

	return nil
}

func (s *Service) New(args string) (*Box, int, error) {
	allEntities = append(allEntities, Box{
		Title: args,
	})

	return &allEntities[len(allEntities)-1], len(allEntities) - 1, nil
}
