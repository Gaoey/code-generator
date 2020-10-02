package test

type Repositorier interface {
	GetAll(offset, limit int, userId uuid.UUID) ([]models.Test, error)
	GetById(id uuid.UUID) (*models.Test, error)
	Add(test models.Test) error
	Update(test models.Test) error
	Remove(id string) error
}

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (repo *Repository) GetAll(offset, limit int, userId uuid.UUID) ([]models.Test, error) {
	var test []models.Test
	err := repo.db.Table("").Offset(offset).Limit(limit).Where("is_deleted = ? and user_id = ?", 0, userId).Find(&test).Error
	if err != nil {
		return nil, err
	}
	return test, nil
}

func (repo *Repository) GetById(id uuid.UUID) (*models.Test, error) {
	var test models.Test
	err := repo.db.Table("").Where("id = ? and is_deleted = ?", id, 0).Find(&test).Error
	if err != nil {
		return nil, err
	}

	return &test, nil
}

func (repo *Repository) Add(test models.Test) error {
	err := repo.db.Table("").Create(&test).Error
	if err != nil {
		return err
	}

	return nil
}

func (repo *Repository) Update(test models.Test) error {
	err := repo.db.Table("").Save(&test).Error
	if err != nil {
		return err
	}

	return nil
}

func (repo *Repository) Remove(id string) error {
	err := repo.db.Table("").
		Where("id = ?", id).
		Update("is_deleted", 1).
		Update("deleted_at", time.Now()).
		Error
	if err != nil {
		return err
	}

	return nil
}
