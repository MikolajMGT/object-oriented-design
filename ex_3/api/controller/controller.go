package controller

import (
	"ex_3/api/manager"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
)

type Controller struct {
	srv *echo.Echo
}

func NewController() *Controller {
	return &Controller{srv: echo.New()}
}

func (c *Controller) Start(addr string) {

	manager.GetInstance().Start()

	c.srv.GET("/user/get", c.getUser)
	c.srv.POST("/user/create", c.createUser)
	c.srv.PUT("/user/update", c.updateUser)
	c.srv.DELETE("/user/delete", c.deleteUser)

	go func() {
		if err := c.srv.Start(addr); err != nil {
			log.Panic().Err(err).Msg("unable to start http server")
		}
	}()
}

func (c *Controller) Stop() {
	if err := c.srv.Close(); err != nil {
		log.Info().Err(err).Msg("error while closing http server...")
	}

	manager.GetInstance().Stop()
}