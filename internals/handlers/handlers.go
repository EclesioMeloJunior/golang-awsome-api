package handlers

import (
	"strconv"

	"github.com/labstack/echo/v4"
)

// Response will define the api contract
// the api reponse
type Response struct {
	Message string      `json:"message,omitempty"`
	Body    interface{} `json:"data,omitempty"`
	Success bool        `json:"success"`
}

// getPagination is a internal function to extract
// the page and size from querystring
func getPagination(c echo.Context) (int, int) {
	pageQ := c.QueryParam("page")
	sizeQ := c.QueryParam("size")

	var page int
	var size int

	var err error
	if page, err = strconv.Atoi(pageQ); err != nil {
		page = 0
	}

	if size, err = strconv.Atoi(sizeQ); err != nil {
		size = 10
	}

	return page, size
}

func errorResponse(err error) *Response {
	return &Response{
		Body:    nil,
		Success: false,
		Message: err.Error(),
	}
}

func successResponse(data interface{}) *Response {
	return &Response{
		Body:    data,
		Success: true,
	}
}
