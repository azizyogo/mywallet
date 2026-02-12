package user

import (
	"mywallet/model"

	"gorm.io/gorm"
)

type (
	UserRepositoryItf interface {
		Create(user *model.User) error
		FindByEmail(email string) (*model.User, error)
		FindByID(id uint) (*model.User, error)
	}

	UserRepository struct {
		resource UserResourceItf
	}

	UserResourceItf interface {
		create(user *model.User) error
		findByEmail(email string) (*model.User, error)
		findByID(id uint) (*model.User, error)
	}

	UserResource struct {
		DB *gorm.DB
	}
)

func InitRepository(rsc UserResourceItf) UserRepository {
	return UserRepository{
		resource: rsc,
	}
}

func (d UserRepository) Create(user *model.User) error {
	return d.resource.create(user)
}

func (d UserRepository) FindByEmail(email string) (*model.User, error) {
	user, err := d.resource.findByEmail(email)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (d UserRepository) FindByID(id uint) (*model.User, error) {
	user, err := d.resource.findByID(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}
