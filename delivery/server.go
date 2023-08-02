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
	s.setupControllers()
	err := s.engine.Run(s.host)
	if err != nil {
		panic(err)
	}
}

func (s *Server) setupControllers() {
	s.engine.Use(middleware.LogRequestMiddleware(s.log))
	// semua controller disini
	api.NewUomController(s.useCaseManager.UomUseCase(), s.engine)
	api.NewProductController(s.engine, s.useCaseManager.ProductUseCase())
	api.NewCustomerController(s.engine, s.useCaseManager.CustomerUseCase())
	api.NewEmployeeController(s.engine, s.useCaseManager.EmployeeUseCase())
	api.NewBillController(s.engine, s.useCaseManager.BillUseCase())
	api.NewUserController(s.engine, s.useCaseManager.UserUseCase())
}

func NewServer() *Server {
	cfg, err := config.NewConfig()
	exceptions.CheckErr(err)
	infraManager, _ := manager.NewInfraManager(cfg)
	repoManager := manager.NewRepoManager(infraManager)
	useCaseManager := manager.NewUseCaseManager(repoManager)
	engine := gin.Default()
	host := fmt.Sprintf("%s:%s", cfg.ApiHost, cfg.ApiPort)
	return &Server{
		useCaseManager: useCaseManager,
		engine:         engine,
		host:           host,
		log:            logrus.New(),
	}
}
