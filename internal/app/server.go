package app

import (
	"github.com/gin-gonic/gin"
	"github.com/ruancaetano/go-gin-fruits/internal/domain/usecase"
	"github.com/ruancaetano/go-gin-fruits/internal/infra/repository"
	"github.com/ruancaetano/go-gin-fruits/internal/presentation/handler"

	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
)

type Server struct {
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) Start() {
	r := gin.Default()

	s.setupRoutes(r)

	if r.Run() != nil {
		panic("fail to start server")
	}
}

func (s *Server) setupRoutes(r *gin.Engine) {
	mrepository := repository.NewFruitMemoryRepository()

	searchFruitUseCase := usecase.NewSearchFruitUseCase(mrepository)
	createFruitUseCase := usecase.NewCreateFruitUseCase(mrepository)
	getFruitUseCase := usecase.NewGetFruitUseCase(mrepository)
	updateFruitUseCase := usecase.NewUpdateFruitUseCase(mrepository)
	deleteFruitUseCase := usecase.NewDeleteFruitUseCase(mrepository)

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.GET("/fruits/search", handler.MakeSearchFruitHandler(searchFruitUseCase))
	r.GET("/fruits/:id", handler.MakeGetFruitHandler(getFruitUseCase))
	r.POST("/fruits", handler.MakeCreateFruitHandler(createFruitUseCase))
	r.PUT("/fruits/:id", handler.MakeUpdateFruitHandler(updateFruitUseCase))
	r.DELETE("/fruits/:id", handler.MakeDeleteFruitHandler(deleteFruitUseCase))
}
