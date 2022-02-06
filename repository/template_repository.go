package repositories

import (
	"github.com/yusuftatli/go-template-with-redis-gorm-api/entities"
	"gorm.io/gorm"
)

type TemplateRepositoryInterface interface {
	Insert(i *entities.Template) (*entities.Template, error)
	Update(i *entities.Template) (*entities.Template, error)
	Delete(i *entities.Template) error
	FindAll(i []*entities.Template) ([]*entities.Template, error)
}

type TemplateRepository struct {
	db *gorm.DB
}

func NewTemplateRepository(db *gorm.DB) TemplateRepositoryInterface {
	return &TemplateRepository{db}
}

func (repository *TemplateRepository) Insert(i *entities.Template) (*entities.Template, error) {
	err := repository.db.Create(&i).Error
	if err != nil {
		return nil, err
	}
	return i, nil
}

func (repository *TemplateRepository) Update(i *entities.Template) (*entities.Template, error) {
	err := repository.db.Save(&i).Error
	if err != nil {
		return nil, err
	}
	return i, nil
}

func (repository *TemplateRepository) Delete(i *entities.Template) error {
	return repository.db.Delete(&i).Error
}

func (repository *TemplateRepository) FindFirst(i *entities.Template) (*entities.Template, error) {
	err := repository.db.Take(&i).Error
	if err != nil {
		return nil, err
	}

	return i, nil
}

func (repository *TemplateRepository) FindAll(a []*entities.Template) ([]*entities.Template, error) {
	dbQuery := repository.db.Preload("Template")
	err := dbQuery.Find(&a).Error

	if err != nil {
		return nil, err
	}
	return a, nil
}