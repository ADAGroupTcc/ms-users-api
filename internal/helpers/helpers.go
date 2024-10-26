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
	RawUserIds     string `query:"user_ids"`
	UserIDs        []string
	ShowCategories bool  `query:"show_categories"`
	Limit          int64 `query:"limit"`
	Offset         int64 `query:"next_page"`
}

func (q *QueryParams) normalize() {
	q.UserIDs = strings.Split(q.RawUserIds, ",")
	q.RawUserIds = ""
	if q.Limit < 1 {
		q.Limit = 10
	}
	if q.Offset < 1 {
		q.Offset = 1
	}
}
