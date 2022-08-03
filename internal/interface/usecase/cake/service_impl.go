package cake

import (
	"context"
	"fmt"

	"github.com/agungwicaksana/privy-pretest/internal/domain"
	"github.com/agungwicaksana/privy-pretest/pkg/response"
	"github.com/labstack/gommon/log"
)

type service struct {
	cakeRepo domain.CakeRepository
}

func NewService() *service {
	return &service{}
}

func (s *service) SetCakeRepo(repo domain.CakeRepository) *service {
	s.cakeRepo = repo
	return s
}

func (s *service) Validate() CakeService {
	if s.cakeRepo == nil {
		panic("cakeRepo is nil")
	}

	return s
}

func (s *service) Find(ctx context.Context, page, limit int, sortBy []Sort) (resp response.Response) {
	start := limit * (page - 1)

	sortByStr := ""
	for _, s := range sortBy {
		sorter := fmt.Sprintf("%s %s,", s.Column, s.Type)
		sortByStr = fmt.Sprintf("%s %s", sortByStr, sorter)
	}
	sortByStr = sortByStr[:len(sortByStr)-1]
	data, err := s.cakeRepo.Find(ctx, start, limit, sortByStr)
	if err != nil {
		log.Error("SQL Error", err)
		resp = response.CreateErrorResponse(response.STATUS_SQL_ERROR, "")
		return
	}
	resp = response.CreateResponse(response.STATUS_READ_SUCCESS, response.MESSAGE_SUCCESS, data)
	return
}

func (s *service) Save(ctx context.Context, req CakeRequest) (resp response.Response) {
	trx, err := s.cakeRepo.BeginTrx(ctx)
	if err != nil {
		log.Error("SQL Error", err)
		resp = response.CreateErrorResponse(response.STATUS_SQL_ERROR, "")
		return
	}

	data := domain.CakeEntity{
		Title:       req.Title,
		Description: req.Description,
		Rating:      req.Rating,
		Image:       req.Image,
	}

	result, err := s.cakeRepo.Insert(ctx, trx, data)
	if err != nil {
		log.Error("SQL Error", err)
		resp = response.CreateErrorResponse(response.STATUS_SQL_ERROR, "")
		return
	}
	trx.Commit()
	if err != nil {
		log.Error("SQL Error", err)
		resp = response.CreateErrorResponse(response.STATUS_SQL_ERROR, "")
		return
	}

	id, err := result.LastInsertId()
	if err != nil {
		log.Error("Create Error", err)
		resp = response.CreateErrorResponse(response.STATUS_CREATE_ERROR, "")
		return
	}

	resp = response.CreateResponse(response.STATUS_CREATE_SUCCESS, response.MESSAGE_SUCCESS, id)
	return
}
