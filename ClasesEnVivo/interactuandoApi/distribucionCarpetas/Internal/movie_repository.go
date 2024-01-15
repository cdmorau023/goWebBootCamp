package internal

import "errors"

var (
	ErrorMovieTitleAlreadyExists = errors.New("movie title already exists")
)

type MovieRepository interface {
	Save(movie *Movie) (err error)
}
