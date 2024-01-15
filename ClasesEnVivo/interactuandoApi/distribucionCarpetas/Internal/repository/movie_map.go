package repository

type MovieMap struct {
	db map[int]internal.Movie
}

func (m *MovieMap) Save(movie internal.Movie) error {

	for _, m := range m.db {
		if m.Title == movie.Title {
			return ErrorMovieTitleAlreadyExists
		}
	}
	
	m.db[movie.ID] = movie
	return nil
}
