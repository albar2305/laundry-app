package delivery

import (
	"fmt"
	"os"

	"github.com/albar2305/enigma-laundry-apps/config"
	"github.com/albar2305/enigma-laundry-apps/delivery/controller/cli"
	"github.com/albar2305/enigma-laundry-apps/repository"
	"github.com/albar2305/enigma-laundry-apps/usecase"
	"github.com/albar2305/enigma-laundry-apps/utils/exceptions"
)

type Console struct {
	uomUC     usecase.UomUseCase
	prodcutUC usecase.ProductUseCase
}

func (c *Console) mainMenuForm() {
	fmt.Println(`
|++++ ENIGMA LAUNDRY MENU ++++|
| 1. Master UOM               |
| 2. Master Product           |
| 3. Master Customer          |
| 4. Master Employee          |
| 5. Transaksi                |
| 6. Keluar                   |
			`)
	fmt.Print("Pilih Menu (1-6):")
}

func (c *Console) Run() {
	for {
		c.mainMenuForm()

		var selectedMenu string
		fmt.Scanln(&selectedMenu)
		switch selectedMenu {
		case "1":
			controller := cli.NewUomController(c.uomUC)
			controller.UomMenuForm()
		case "2":
			productController := cli.NewProductController(c.prodcutUC)
			productController.HandlerMainForm()
		case "3":
			fmt.Println("Master Customer")
		case "4":
			fmt.Println("Master Employee")
		case "5":
			fmt.Println("Transaksi")
		case "6":
			os.Exit(0)
		default:
			fmt.Println("Menu tidak ditemunkan")
		}
	}
}

func NewConsole() *Console {
	cfg, err := config.NewConfig()
	exceptions.CheckErr(err)
	dbConn, _ := config.NewDbConnection(cfg)
	db := dbConn.Conn()
	uomRepo := repository.NewUomRepository(db)
	productRepo := repository.NewProductRepository(db)
	uomUseCase := usecase.NewUomUseCase(uomRepo)
	productUseCase := usecase.NewProductUseCase(productRepo, uomUseCase)
	return &Console{
		uomUC:     uomUseCase,
		prodcutUC: productUseCase,
	}
}
