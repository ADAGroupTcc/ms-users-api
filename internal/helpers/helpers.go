package helpers

import (
	"strings"

	"github.com/labstack/echo/v4"
)

func BindQueryParams(c echo.Context, queryParams *QueryParams) error {
	if err := c.Bind(queryParams); err != nil {
		return err
	}
	queryParams.normalize()
	return nil
}

type QueryParams struct {
	RawUserIds string `json:"user_ids" query:"user_ids"`
	UserIDs    []string
	CategoryID []string `json:"category_id,omitempty" query:"category_id"`
	Limit      int64    `json:"limit" query:"limit"`
	Offset     int64    `json:"next_page" query:"next_page"`
}

func (q *QueryParams) normalize() {
	q.UserIDs = strings.Split(q.RawUserIds, ",")
	q.RawUserIds = ""
	if q.Limit < 1 {
		q.Limit = 10
	}
	if q.Offset < 0 {
		q.Offset = 0
	}
}
