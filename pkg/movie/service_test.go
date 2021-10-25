package movie

import (
	"testing"

	"github.com/andreastihor/stockbit/models"
	"github.com/andreastihor/stockbit/pkg/movie/mocks"
	"github.com/stretchr/testify/assert"
)

func TestGetMovies(t *testing.T) {
	repo := mocks.RepositoryInterface{}
	service := NewService(&Repository{})
	mockMovie := models.Movie{
		Title:  "testing title",
		Year:   "1998",
		ImdbID: "123dasd",
		Type:   "movie",
		Poster: "https://m.media-amazon.com/images/M/MV5BYTEzMmE0ZDYtYWNmYi00ZWM4LWJjOTUtYTE0ZmQyYWM3ZjA0XkEyXkFqcGdeQXVyNTA4NzY1MzY@._V1_SX300.jpg",
	}
	repo.On("InsertMovie", mockMovie).Return(mockMovie).Once()

	mockParams := models.Params{
		Pagination: "1",
		Search:     "batman",
	}

	movies, err := service.GetMovies(&mockParams)
	assert.NoError(t, err)
	assert.NotNil(t, len(movies))
}
