package handlers

import (
	"golang-template/modules/common/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

type getExampleResponse struct {
	Id     string `json:"id"`
	Title  string `json:"title"`
	Detail string `json:"detail"`
}

func (h *handler) GetItemExampleHandler(c echo.Context) error {
	id := c.Param("id")

	// Get data
	p, err := h.store.GetById(id)
	if err != nil {
		res := models.DataNotFound()
		return c.JSON(http.StatusNotFound, res)
	}

	// Save data to cachev
	// data, _ := json.Marshal(book)
	// t.cache.SetShortCache(cacheKey, data)

	// Response
	res := getExampleResponse{}
	res.Id = p.Id.Hex()
	res.Title = p.Title
	res.Detail = p.Detail

	resp := models.Success(res)

	return c.JSON(http.StatusOK, resp)
}
