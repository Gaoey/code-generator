package goldornament

type Repositorier interface {
	GetAll(offset, limit int, userId uuid.UUID) ([]models.GoldOrnamentAsset, error)
	GetById(id uuid.UUID) (*models.GoldOrnamentAsset, error)
	Add(goldornament models.GoldOrnamentAsset) error
	Update(goldornament models.GoldOrnamentAsset) error
	Remove(id string) error
}

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (repo *Repository) GetAll(offset, limit int, userId uuid.UUID) ([]models.GoldOrnamentAsset, error) {
	var goldornament []models.GoldOrnamentAsset
	err := repo.db.Table("").Offset(offset).Limit(limit).Where("is_deleted = ? and user_id = ?", 0, userId).Find(&goldornament).Error
	if err != nil {
		return nil, err
	}
	return goldornament, nil
}

func (repo *Repository) GetById(id uuid.UUID) (*models.GoldOrnamentAsset, error) {
	var goldornament models.GoldOrnamentAsset
	err := repo.db.Table("").Where("id = ? and is_deleted = ?", id, 0).Find(&goldornament).Error
	if err != nil {
		return nil, err
	}

	return &goldornament, nil
}

func (repo *Repository) Add(goldornament models.GoldOrnamentAsset) error {
	err := repo.db.Table("").Create(&goldornament).Error
	if err != nil {
		return err
	}

	return nil
}

func (repo *Repository) Update(goldornament models.GoldOrnamentAsset) error {
	err := repo.db.Table("").Save(&goldornament).Error
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
