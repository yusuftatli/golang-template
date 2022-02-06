package repositories

import (
	"github.com/yusuftatli/go-template-with-redis-gorm-api/entities"
	"gorm.io/gorm"
)

type UserRepositoryInterface interface {
	FindFirst(i *entities.User) (*entities.User, error)
}

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepositoryInterface {
	return &UserRepository{db}
}

func (repository *UserRepository) FindFirst(i *entities.User) (*entities.User, error) {
	err := repository.db.Take(&i).Error
	if err != nil {
		return nil, err
	}

	return i, nil
}
