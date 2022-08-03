package cake

import (
	"context"

	"github.com/agungwicaksana/privy-pretest/pkg/response"
)

type CakeService interface {
	Find(ctx context.Context, page, limit int, sortBy []Sort) (resp response.Response)
	Save(ctx context.Context, req CakeRequest) (resp response.Response)
	FindOne(ctx context.Context, id string) (resp response.Response)
}
