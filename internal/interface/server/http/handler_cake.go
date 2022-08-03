package http

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/agungwicaksana/privy-pretest/internal/interface/usecase/cake"
	"github.com/agungwicaksana/privy-pretest/pkg/helpers"
	"github.com/agungwicaksana/privy-pretest/pkg/response"
	"github.com/labstack/echo/v4"
)

type cakeHandler struct {
	cakeService cake.CakeService
}

func NewCakeHandler() *cakeHandler {
	return &cakeHandler{}
}

func (h *cakeHandler) SetCakeService(service cake.CakeService) *cakeHandler {
	h.cakeService = service
	return h
}

func (h *cakeHandler) Validate() *cakeHandler {
	if h.cakeService == nil {
		panic("cakeService is nil")
	}

	return h
}

func (h *cakeHandler) Find(c echo.Context) (err error) {
	ctx := c.Request().Context()

	pageStr := c.QueryParam("page")
	limitStr := c.QueryParam("limit")

	page := 1
	limit := 5

	if pageStr != "" {
		page, err = strconv.Atoi(pageStr)
		if err != nil {
			c.Set("RC", response.STATUS_REQUEST_VALIDATION_FAILED)
			return
		}
	}

	if limitStr != "" {
		limit, err = strconv.Atoi(limitStr)
		if err != nil {
			c.Set("RC", response.STATUS_REQUEST_VALIDATION_FAILED)
			return
		}
	}

	sortStr := c.QueryParam("sort")

	var (
		sort    []string
		sortBy  []cake.Sort
		columns []string = []string{"title", "description", "rating", "image", "created_at", "updated_at"}
	)

	if sortStr != "" {
		sort = strings.Split(sortStr, ",")
		for _, s := range sort {
			types := "ASC"
			column := s
			if s[0:1] == "-" {
				types = "DESC"
				column = s[1:]
			}

			if isContain := helpers.IsSliceContainStr(columns, column); !isContain {
				fmt.Println(column)
				err = errors.New("Sorting column is not valid")
				c.Set("RC", response.STATUS_REQUEST_VALIDATION_FAILED)
				return
			}

			sortBy = append(sortBy, cake.Sort{
				Column: column,
				Type:   types,
			})
		}
	} else {
		sortBy = []cake.Sort{
			{
				Column: "rating",
				Type:   "DESC",
			},
			{
				Column: "title",
				Type:   "ASC",
			},
		}
	}

	resp := h.cakeService.Find(ctx, page, limit, sortBy)
	return c.JSON(http.StatusOK, resp)
}

func (h *cakeHandler) Save(c echo.Context) (err error) {
	ctx := c.Request().Context()

	req := cake.CakeRequest{}
	if err = helpers.Validate(c, &req); err != nil {
		return
	}

	resp := h.cakeService.Save(ctx, req)
	return c.JSON(http.StatusOK, resp)
}

func (h *cakeHandler) FindOne(c echo.Context) (err error) {
	ctx := c.Request().Context()

	id := c.Param("id")

	_, err = strconv.Atoi(id)
	if err != nil {
		c.Set("RC", response.STATUS_REQUEST_VALIDATION_FAILED)
		return
	}

	resp := h.cakeService.FindOne(ctx, id)
	return c.JSON(http.StatusOK, resp)
}

func (h *cakeHandler) Update(c echo.Context) (err error) {
	ctx := c.Request().Context()

	id := c.Param("id")

	_, err = strconv.Atoi(id)
	if err != nil {
		c.Set("RC", response.STATUS_REQUEST_VALIDATION_FAILED)
		return
	}

	req := cake.CakeRequest{}
	if err = helpers.Validate(c, &req); err != nil {
		return
	}

	resp := h.cakeService.Update(ctx, id, req)
	return c.JSON(http.StatusOK, resp)
}
