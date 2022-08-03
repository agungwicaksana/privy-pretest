package cake

import (
	"context"

	"github.com/agungwicaksana/privy-pretest/pkg/response"
)

type CakeService interface {
	Find(ctx context.Context, page, limit int, sortBy []Sort) (resp response.Response)
	// FindOne(ctx context.Context, id string) (resp response.Response, err error)
	// Update(ctx context.Context, req UpdateFaqReq) (resp RequestOtpResponse, err error)
}
