package movie

import (
	"github.com/andreastihor/stockbit/models"
	"gorm.io/gorm"
)

//Repository ...
type Repository struct {
	DB *gorm.DB
}

type RepositoryInterface interface {
	InsertMovie(movie models.Movie) error
}

//NewRepository ...
func NewRepository(db *gorm.DB) RepositoryInterface {
	return &Repository{db}
}

func (repo *Repository) InsertMovie(movie models.Movie) error {
	err := repo.DB.Create(movie).Error
	return err
}
