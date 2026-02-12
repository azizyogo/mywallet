package user

import "mywallet/model"

func (rsc UserResource) create(user *model.User) error {
	return rsc.DB.Create(user).Error
}

func (rsc UserResource) findByEmail(email string) (*model.User, error) {
	var user model.User
	err := rsc.DB.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (rsc UserResource) findByID(id uint) (*model.User, error) {
	var user model.User
	err := rsc.DB.Where("id = ?", id).First(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}
