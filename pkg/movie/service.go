package movie

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/andreastihor/stockbit/models"
)

type Service struct {
	logger     *log.Logger
	Repository *Repository
}

type ServiceInterface interface {
	GetMovies(params *models.Params) ([]*models.Movie, error)
}

// NewService will initialize the service for  user endpoint
func NewService(l *log.Logger, repo *Repository) ServiceInterface {
	return &Service{
		logger:     l,
		Repository: repo,
	}
}

func (s *Service) GetMovies(params *models.Params) ([]*models.Movie, error) {
	if params.APIKey == "" {
		return nil, errors.New("apikey must be provided")
	}

	if params.Search == "" {
		return nil, errors.New("movie name must be provided")
	}

	results := []*models.Movie{}
	datas, err := getDataFromIMDB(params)
	if err != nil {
		return nil, err
	}

	if datas.Error != "" {
		return nil, errors.New(datas.Error)
	}

	for _, data := range datas.Search {
		result := &models.Movie{}
		result.AsResponse(data)
		results = append(results, result)
	}

	err = s.Repository.InsertMovie(*results[0])
	if err != nil {
		return nil, err
	}

	return results, nil
}

func getDataFromIMDB(params *models.Params) (*models.IMDBMovieRequest, error) {
	result := &models.IMDBMovieRequest{}
	queryParams := fmt.Sprintf("?apikey=%s&s=%s", params.APIKey, params.Search)
	if params.Pagination != "" {
		queryParams += fmt.Sprintf("&page=%s", params.Pagination)
	}

	url := "http://www.omdbapi.com/" + queryParams
	response, err := http.Get(url)
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	err = json.Unmarshal(responseData, &result)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return result, nil

}
