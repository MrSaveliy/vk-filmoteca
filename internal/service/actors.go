package service

import (
	"github.com/MrSaveliy/vk-filmoteca/internal/repository"
	"github.com/MrSaveliy/vk-filmoteca/models"
)

type ActorsService struct {
	repo repository.Actors
}

func NewActorsService(repo repository.Actors) *ActorsService {
	return &ActorsService{repo: repo}
}

func (s *ActorsService) CreateActor(actor models.Actor) (int, error) {
	return s.repo.CreateActor(actor)
}

func (s *ActorsService) GetActorById(id int) (models.Actor, error) {
	return s.repo.GetActorById(id)
}
