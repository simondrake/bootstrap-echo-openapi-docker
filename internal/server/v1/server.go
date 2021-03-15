package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Server struct {

}

func New() *Server {
	return &Server{}
}

func (s *Server) TestV1(c echo.Context) error {
	//ctx := c.Request().Context()
	return c.JSON(http.StatusOK, map[string]string{"Hello": "World"})
}

func (s *Server) RootV1(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{"Message": "Hello Root"})
}
