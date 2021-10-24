package models

type IMBDMovie struct {
	Title  string `json:"Title"`
	Year   string `json:"Year"`
	ImdbID string `json:"imdbID"`
	Type   string `json:"Type"`
	Poster string `json:"Poster"`
}

type IMDBMovieRequest struct {
	Search   []*IMBDMovie `json:"Search"`
	Response string       `json:"Response"`
	Error    string       `json:"Error"`
}

type Params struct {
	Pagination string `schema:"page"`
	Search     string `schema:"s"`
}

type Movie struct {
	Title  string `json:"title"`
	Year   string `json:"year"`
	ImdbID string `json:"imdb_id"`
	Type   string `json:"type"`
	Poster string `json:"poster"`
}

func (data *Movie) AsResponse(imdbData *IMBDMovie) {
	data.ImdbID = imdbData.ImdbID
	data.Title = imdbData.Title
	data.Year = imdbData.Year
	data.Type = imdbData.Type
	data.Poster = imdbData.Poster
}
