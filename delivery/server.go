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
	uomUC      usecase.UomUseCase
	productUC  usecase.ProductUseCase
	employeeUC usecase.EmployeeUseCase
	customerUC usecase.CustomerUseCase
	engine     *gin.Engine
	host       string
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
	api.NewProductController(s.engine, s.productUC)
	api.NewEmployeeController(s.engine, s.employeeUC)
	api.NewCustomerController(s.engine, s.customerUC)
}
func NewServer() *Server {
	cfg, err := config.NewConfig()
	exceptions.CheckErr(err)
	dbConn, _ := config.NewDbConnection(cfg)
	db := dbConn.Conn()
	uomRepo := repository.NewUomRepository(db)
	productRepo := repository.NewProductRepository(db)
	employeeRepo := repository.NewEmployeeRepository(db)
	customerRepo := repository.NewCustomerRepository(db)
	uomUseCase := usecase.NewUomUseCase(uomRepo)
	productUseCase := usecase.NewProductUseCase(productRepo, uomUseCase)
	employeeUseCase := usecase.NewEmployeeUseCase(employeeRepo)
	customerUseCase := usecase.NewCustomerUseCase(customerRepo)

	engine := gin.Default()
	host := fmt.Sprintf("%s:%s", cfg.ApiHost, cfg.ApiPort)
	return &Server{
		uomUC:      uomUseCase,
		productUC:  productUseCase,
		employeeUC: employeeUseCase,
		customerUC: customerUseCase,
		engine:     engine,
		host:       host}
}
