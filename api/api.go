package api

import (
	"database/sql"

	"github.com/barnbridge/smartbackend/config"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type API struct {
	engine *gin.Engine
	db     *sql.DB
	logger *logrus.Entry
}

func New(db *sql.DB) *API {
	return &API{
		db:     db,
		logger: logrus.WithField("module", "api"),
	}
}

func (a *API) Run() {
	a.engine = gin.Default()

	if config.Store.API.DevCors {
		a.engine.Use(cors.New(cors.Config{
			AllowOrigins:     []string{config.Store.API.DevCorsHost},
			AllowMethods:     []string{"PUT", "PATCH", "GET", "POST"},
			AllowHeaders:     []string{"Origin"},
			ExposeHeaders:    []string{"Content-Length"},
			AllowCredentials: true,
		}))
	}

	a.setRoutes()

	err := a.engine.Run(":" + config.Store.API.Port)
	if err != nil {
		a.logger.Fatal(err)
	}
}

func (a *API) Close() {
}
