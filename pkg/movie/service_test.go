package movie

import (
	"testing"

	"github.com/andreastihor/stockbit/models"
	"github.com/andreastihor/stockbit/pkg/movie/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetMovies(t *testing.T) {
	repo := mocks.RepositoryInterface{}
	service := NewService(&repo)

	repo.On("InsertMovie", mock.Anything).Return(nil)
	mockParams := models.Params{
		Pagination: "1",
		Search:     "batman",
	}

	movies, err := service.GetMovies(&mockParams)
	assert.NoError(t, err)
	assert.NotNil(t, len(movies))
}
