package subdomain

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (s *Service) List() []Subdomain {
	return allEntities
}

func (s *Service) Get(idx int) (*Subdomain, error) {
	return &allEntities[idx], nil
}

func (s *Service) Edit(idx int, title string) (*Subdomain, error) {
	// Поменять на правильную логику
	allEntities[idx] = Subdomain{
		Title: title,
	}

	return &allEntities[idx], nil
}

func (s *Service) Delete(idx int) error {
	// Подумать об обработке ошибок
	allEntities[idx] = allEntities[len(allEntities)-1]
	allEntities[len(allEntities)-1] = Subdomain{}
	allEntities = allEntities[:len(allEntities)-1]

	return nil
}

func (s *Service) New(args string) (*Subdomain, int, error) {
	allEntities = append(allEntities, Subdomain{
		Title: args,
	})

	return &allEntities[len(allEntities)-1], len(allEntities) - 1, nil
}
