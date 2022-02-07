package repositories

import (
	"github.com/yusuftatli/go-template-with-redis-gorm-api/internal/entities"
	"gorm.io/gorm"
)

type IRepositoryExample interface {
	Insert(i *entities.Example) (*entities.Example, error)
	Update(i *entities.Example) (*entities.Example, error)
	Delete(i *entities.Example) error
	FindAll(i []*entities.Example) ([]*entities.Example, error)
}

type repositoryExample struct {
	db *gorm.DB
}

func NewRepositoryExample(db *gorm.DB) IRepositoryExample {
	return &repositoryExample{db}
}

func (repository *repositoryExample) Insert(i *entities.Example) (*entities.Example, error) {
	err := repository.db.Create(&i).Error
	if err != nil {
		return nil, err
	}
	return i, nil
}

func (repository *repositoryExample) Update(i *entities.Example) (*entities.Example, error) {
	err := repository.db.Save(&i).Error
	if err != nil {
		return nil, err
	}
	return i, nil
}

func (repository *repositoryExample) Delete(i *entities.Example) error {
	return repository.db.Delete(&i).Error
}

func (repository *repositoryExample) FindFirst(i *entities.Example) (*entities.Example, error) {
	err := repository.db.Take(&i).Error
	if err != nil {
		return nil, err
	}

	return i, nil
}

func (repository *repositoryExample) FindAll(a []*entities.Example) ([]*entities.Example, error) {
	dbQuery := repository.db.Preload("Template")
	err := dbQuery.Find(&a).Error

	if err != nil {
		return nil, err
	}
	return a, nil
}
