package delivery

import (
	"fmt"

	"github.com/albar2305/enigma-laundry-apps/config"
	"github.com/albar2305/enigma-laundry-apps/delivery/controller/api"
	"github.com/albar2305/enigma-laundry-apps/delivery/middleware"
	"github.com/albar2305/enigma-laundry-apps/manager"
	"github.com/albar2305/enigma-laundry-apps/utils/exceptions"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Server struct {
	useCaseManager manager.UseCaseManager
	engine         *gin.Engine
	host           string
	log            *logrus.Logger
}

func (s *Server) Run() {
	s.initController()
	err := s.engine.Run(s.host)
	if err != nil {
		panic(err)
	}
}

func (s *Server) initController() {
	s.engine.Use(middleware.LogRequestMiddleware(s.log))

	api.NewUomController(s.useCaseManager.UomUseCase(), s.engine)
	api.NewProductController(s.engine, s.useCaseManager.ProductUseCase())
	api.NewEmployeeController(s.engine, s.useCaseManager.EmployeeUseCase())
	api.NewCustomerController(s.engine, s.useCaseManager.CustomerUseCase())
	api.NewBillController(s.engine, s.useCaseManager.BillUseCase())
}
func NewServer() *Server {
	cfg, err := config.NewConfig()
	exceptions.CheckErr(err)
	infraManager, _ := manager.NewInfraManager(cfg)
	repoManager := manager.NewRepoManager(infraManager)
	useCaseManager := manager.NewUseCaseManagr(repoManager)
	engine := gin.Default()
	host := fmt.Sprintf("%s:%s", cfg.ApiHost, cfg.ApiPort)
	return &Server{
		useCaseManager: useCaseManager,
		engine:         engine,
		host:           host,
		log:            logrus.New(),
	}
}
