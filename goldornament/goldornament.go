package goldornament

type Servicer interface {
	GetAll(offset, limit int, userId uuid.UUID) ([]models.GoldOrnamentAsset, error)
	GetById(id string) (*models.GoldOrnamentAsset, error)
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

func (s *Service) GetAll(offset, limit int, userId uuid.UUID) ([]models.GoldOrnamentAsset, error) {
	goldornaments, err := s.Repository.GetAll(offset, limit, userId)
	if err != nil {
		logrus.Errorf("cannot get all goldornaments")
		return nil, err
	}

	return goldornaments, nil
}

func (s *Service) GetById(id string) (*models.GoldOrnamentAsset, error) {
	uuid := uuid.MustParse(id)
	goldornament, err := s.Repository.GetById(uuid)
	if err != nil {
		logrus.Errorf("cannot get goldornament by id: %s, %s", id, err.Error())
		return nil, err
	}

	return goldornament, nil
}

// @TODO when create goldornament should create default category
func (s *Service) Add(req AddRequest, userId uuid.UUID) (*AddResponse, error) {
	goldornamentId := uuid.New()
	goldornament := goldornamentRequestToModel(req.Request)
	goldornament.ID = goldornamentId
	goldornament.UserID = userId

	err := s.Repository.Add(goldornament)
	if err != nil {
		logrus.Errorf("cannot create goldornament, %s", err.Error())
		return nil, err
	}

	response := AddResponse{ID: goldornamentId.String()}
	return &response, nil
}

func (s *Service) Update(id string, req UpdateRequest) error {
	goldornament, err := s.GetById(id)
	if err != nil {
		logrus.Errorf("cannot get goldornament by id, %s", err.Error())
		return err
	}

	update := goldornamentRequestToModel(req.Request)
	update.ID = goldornament.ID
	update.UserID = goldornament.UserID
	err = s.Repository.Update(update)
	if err != nil {
		logrus.Errorf("cannot update goldornament by id, %s", err.Error())
		return err
	}

	return nil
}

func (s *Service) Remove(id string) error {
	err := s.Repository.Remove(id)
	if err != nil {
		logrus.Errorf("cannot remove goldornament by id, %s", err.Error())
		return err
	}

	return nil
}
