package delivery

import (
	"fmt"

	"github.com/albar2305/enigma-laundry-apps/config"
	"github.com/albar2305/enigma-laundry-apps/delivery/controller/api"
	"github.com/albar2305/enigma-laundry-apps/repository"
	"github.com/albar2305/enigma-laundry-apps/usecase"
	"github.com/albar2305/enigma-laundry-apps/utils/exceptions"
	"github.com/gin-gonic/gin"
)

type Server struct {
	uomUC  usecase.UomUseCase
	engine *gin.Engine
	host   string
}

func (s *Server) RUn() {
	s.initController()
	err := s.engine.Run(s.host)
	if err != nil {
		panic(err)
	}
}

func (s *Server) initController() {
	api.NewUomController(s.uomUC, s.engine)
}
func NewServer() *Server {
	cfg, err := config.NewConfig()
	exceptions.CheckErr(err)
	dbConn, _ := config.NewDbConnection(cfg)
	db := dbConn.Conn()
	uomRepo := repository.NewUomRepository(db)
	uomUseCase := usecase.NewUomUseCase(uomRepo)

	engine := gin.Default()
	host := fmt.Sprintf("%s:%s", cfg.ApiHost, cfg.ApiPort)
	return &Server{uomUC: uomUseCase, engine: engine, host: host}
}
