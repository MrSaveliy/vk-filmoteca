package repository

import (
	"fmt"

	"github.com/MrSaveliy/vk-filmoteca/models"
	"github.com/jmoiron/sqlx"
)

type ActorsPostgres struct {
	db *sqlx.DB
}

func NewActorsPostgres(db *sqlx.DB) *ActorsPostgres {
	return &ActorsPostgres{db: db}
}

func (r *ActorsPostgres) CreateActor(actor models.Actor) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (name, gender, date_birth, movie_id) values ($1, $2, $3, $4) RETURNING id", actorsTable)
	row := r.db.QueryRow(query, actor.Name, actor.Gender, actor.Date_birth, actor.Movie_id)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *ActorsPostgres) GetActorById(id int) (models.Actor, error) {
	var actor models.Actor
	query := fmt.Sprintf("SELECT name, gender, date_birth, movie_id FROM %s WHERE id=$1", actorsTable)
	err := r.db.Get(&actor, query, id)
	return actor, err
}
