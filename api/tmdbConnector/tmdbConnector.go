package tmdbConnector

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/BassicallyDrums/filmbase/api/models"
	"github.com/BassicallyDrums/filmbase/api/tmdbConnector/auth"
)

const Url string = "https://api.themoviedb.org/3/movie/popular?language=en-US&page=1"

func GetFilms() (models.TmdbResponse, error) {
	auth := auth.Authenticate()

	req, err := http.NewRequest("GET", Url, nil)
	if err != nil {
		return models.TmdbResponse{}, err
	}

	barerToken := "Bearer " + auth.APIToken

	req.Header.Add("accept", "application/json")
	req.Header.Add("Authorization", barerToken)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return models.TmdbResponse{}, err
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return models.TmdbResponse{}, err
	}
	// fmt.Println(string(body))

	var tmbdResponse = models.TmdbResponse{}
	jsonerr := json.Unmarshal(body, &tmbdResponse)
	if jsonerr != nil {
		return models.TmdbResponse{}, jsonerr
	}

	return tmbdResponse, nil
}
