package kittiesbundle

import (
	"github.com/jinzhu/gorm"
)

// KittiesMapperSQL define a SQL mapper
type KittiesMapperSQL struct {
	db *gorm.DB
}

// NewKittiesSQLMapper instance
func NewKittiesSQLMapper(db *gorm.DB) *KittiesMapperSQL {
	return &KittiesMapperSQL{
		db: db,
	}
}

// FindAll kitties in database
func (m *KittiesMapperSQL) FindAll() ([]Kitty, error) {
	var kitties []Kitty

	m.db.Find(&kitties)

	return kitties, nil
}

// Insert implement KittiesMapper interface
func (m *KittiesMapperSQL) Insert(k *Kitty) error {
	return m.db.Create(k).Error
}

// Delete implement KittiesMapper interface
func (m *KittiesMapperSQL) Delete(id int) error {
	return m.db.Delete(&Kitty{ID: id}).Error
}
