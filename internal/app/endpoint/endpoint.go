package endpoint

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

type Service interface {
	DaysLeft() int64
}
type Endpoint struct {
	s Service
}

func New(service Service) *Endpoint {
	return &Endpoint{
		s: service,
	}
}

func (e *Endpoint) Status(ctx echo.Context) error {

	s := fmt.Sprintf("Количество дней: %d", e.s.DaysLeft()/24)

	err := ctx.String(http.StatusOK, s)
	if err != nil {
		return err
	}
	return nil
}
