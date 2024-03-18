package repository

import (
	"fmt"

	"github.com/MrSaveliy/vk-filmoteca/models"
	"github.com/jmoiron/sqlx"
)

type MoviesPostgres struct {
	db *sqlx.DB
}

func NewMoviesPostgres(db *sqlx.DB) *MoviesPostgres {
	return &MoviesPostgres{db: db}
}

func (r *MoviesPostgres) CreateMovie(movie models.Movie) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (name, description, release_date, rating ) values ($1, $2, $3, $4) RETURNING id", moviesTable)
	row := r.db.QueryRow(query, movie.Name, movie.Description, movie.ReleaseDate, movie.Rating)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *MoviesPostgres) GetMovieById(id int) (models.Movie, error) {
	var movie models.Movie
	query := fmt.Sprintf("SELECT name, description, release_date, rating FROM %s WHERE id=$1", moviesTable)
	err := r.db.Get(&movie, query, id)
	return movie, err
}
