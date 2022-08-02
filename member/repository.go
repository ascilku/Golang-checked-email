package member

import "gorm.io/gorm"

type Repository interface {
	SaveRepository(member Member) (Member, error)
	FindByEmail(nama string) (Member, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) SaveRepository(member Member) (Member, error) {
	err := r.db.Create(&member).Error
	if err != nil {
		return member, err
	} else {
		return member, nil
	}
}

func (r *repository) FindByEmail(nama string) (Member, error) {
	var member Member
	err := r.db.Where("nama", nama).Find(&member).Error
	if err != nil {
		return member, err
	} else {
		return member, nil
	}
}
