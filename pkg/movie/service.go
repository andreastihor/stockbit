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
	Repository RepositoryInterface
}

type ServiceInterface interface {
	GetMovies(params *models.Params) ([]*models.Movie, error)
}

// NewService will initialize the service for  user endpoint
func NewService(repo RepositoryInterface) ServiceInterface {
	return &Service{
		Repository: repo,
	}
}

func (s *Service) GetMovies(params *models.Params) ([]*models.Movie, error) {
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

	for _, res := range results {
		err = s.Repository.InsertMovie(*res)
		if err != nil {
			return nil, err
		}
	}

	return results, nil
}

func getDataFromIMDB(params *models.Params) (*models.IMDBMovieRequest, error) {
	result := &models.IMDBMovieRequest{}
	req, err := http.NewRequest("GET", "http://www.omdbapi.com/", nil)
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}

	q := req.URL.Query()
	q.Add("apikey", "faf7e5bb")
	q.Add("s", params.Search)
	if params.Pagination != "" {
		q.Add("page", params.Pagination)
	}

	req.URL.RawQuery = q.Encode()
	response, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Print(err.Error())
		return nil, err
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	err = json.Unmarshal(responseData, &result)
	if err != nil {
		// this error is caused by bad request, we can solve it by implementing a better http call, but for this test, i chose to make it simple only
		if err.Error() == "invalid character '<' looking for beginning of value" {
			return nil, errors.New("PLEASE Delete Space in movie searching title")
		}

		fmt.Println(err)
		return nil, err
	}

	return result, nil
}
