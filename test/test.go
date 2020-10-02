package test

type Servicer interface {
	GetAll(offset, limit int, userId uuid.UUID) ([]models.Test, error)
	GetById(id string) (*models.Test, error)
	Add(req AddRequest, userId uuid.UUID) (*AddResponse, error)
	Update(id string, req UpdateRequest) error
	Remove(id string) error
}

type Service struct {
	Repository Repositorier
}

func NewService(db *gorm.DB) *Service {
	return &Service{
		Repository: NewRepository(db),
	}
}

func (s *Service) GetAll(offset, limit int, userId uuid.UUID) ([]models.Test, error) {
	tests, err := s.Repository.GetAll(offset, limit, userId)
	if err != nil {
		logrus.Errorf("cannot get all tests")
		return nil, err
	}

	return tests, nil
}

func (s *Service) GetById(id string) (*models.Test, error) {
	uuid := uuid.MustParse(id)
	test, err := s.Repository.GetById(uuid)
	if err != nil {
		logrus.Errorf("cannot get test by id: %s, %s", id, err.Error())
		return nil, err
	}

	return test, nil
}

// @TODO when create test should create default category
func (s *Service) Add(req AddRequest, userId uuid.UUID) (*AddResponse, error) {
	testId := uuid.New()
	test := testRequestToModel(req.Request)
	test.ID = testId
	test.UserID = userId

	err := s.Repository.Add(test)
	if err != nil {
		logrus.Errorf("cannot create test, %s", err.Error())
		return nil, err
	}

	response := AddResponse{ID: testId.String()}
	return &response, nil
}

func (s *Service) Update(id string, req UpdateRequest) error {
	test, err := s.GetById(id)
	if err != nil {
		logrus.Errorf("cannot get test by id, %s", err.Error())
		return err
	}

	update := testRequestToModel(req.Request)
	update.ID = test.ID
	update.UserID = test.UserID
	err = s.Repository.Update(update)
	if err != nil {
		logrus.Errorf("cannot update test by id, %s", err.Error())
		return err
	}

	return nil
}

func (s *Service) Remove(id string) error {
	err := s.Repository.Remove(id)
	if err != nil {
		logrus.Errorf("cannot remove test by id, %s", err.Error())
		return err
	}

	return nil
}
