package main

import (
	"fmt"

	"github.com/albar2305/enigma-laundry-apps/config"
	"github.com/albar2305/enigma-laundry-apps/model"
	"github.com/albar2305/enigma-laundry-apps/repository"
	"github.com/albar2305/enigma-laundry-apps/usecase"
	_ "github.com/lib/pq"
)

// func createUom(db *sql.DB, uom model.Uom) error {
// 	_, err := db.Exec("INSERT INTO uom (id,name) VALUES ($1,$2)", uom.Id, uom.Name)
// 	if err != nil {
// 		return err
// 	}
// 	fmt.Println("UOM created successfully")
// 	return nil
// }
//
// func mainMenuForm() {
// 	fmt.Println(`
// |++++ ENIGMA LAUNDRY MENU ++++|
// | 1. Master UOM               |
// | 2. Master Product           |
// | 3. Master Customer          |
// | 4. Master Employee          |
// | 5. Transaksi                |
// | 6. Keluar                   |
// 			`)
// 	fmt.Print("Pilih Menu (1-6):")
// }

// func uomMenuForm() {
// 	fmt.Println(`
// |==== Master UOM ====|
// | 1. Tambah Data     |
// | 2. Lihat Data      |
// | 3. Update Data     |
// | 4. Hapus Data      |
// | 5. Kembali ke Menu |
// 			`)
// 	fmt.Print("Pilih Menu (1-5):")
// 	db := connectDB()
// 	defer db.Close()
// 	for {
// 		var selectedMenu string
// 		fmt.Scanln(&selectedMenu)
// 		switch selectedMenu {
// 		case "1":
// 			uom := uomCreateForm()
// 			err := createUom(db, uom)
// 			checkErr(err)
// 			return
// 		case "2":
// 			fmt.Println("Lihat Data")
// 		case "3":
// 			fmt.Println("Update Data")
// 		case "4":
// 			fmt.Println("Hapus Data")
// 		case "5":
// 			return
// 		default:
// 			fmt.Println("Menu tidak ditemunkan")
// 		}
// 	}
// }

// func uomCreateForm() model.Uom {
// 	var (
// 		uomId, uomName, saveConfirmation string
// 	)

// 	fmt.Print("UOM ID: ")
// 	fmt.Scanln(&uomId)
// 	fmt.Print("UOM Name: ")
// 	fmt.Scanln(&uomName)
// 	fmt.Printf("UOM Id: %s, Name: %s akan disimpan (y/t)", uomId, uomName)
// 	fmt.Scanln(&saveConfirmation)
// 	if saveConfirmation == "y" {
// 		uom := model.Uom{
// 			Id:   uomId,
// 			Name: uomName,
// 		}
// 		return uom
// 	}

// 	return model.Uom{}
// }

// func runConsole() {
// 	for {
// 		mainMenuForm()
// 		var selectedMenu string
// 		fmt.Scanln(&selectedMenu)
// 		switch selectedMenu {
// 		case "1":
// 			uomMenuForm()
// 		case "6":
// 			os.Exit(0)
// 		default:
// 			fmt.Println("Menu tidak ditemunkan")
// 		}
// 	}
// }

// func checkErr(err error) {
// 	if err != nil {
// 		log.Println(err)
// 	}
// }

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		fmt.Println(err)
	}
	dbConn, _ := config.NewDbConnection(cfg)
	db := dbConn.Conn()
	uomRepo := repository.NewUomRepository(db)
	uomUseCase := usecase.NewUomUseCase(uomRepo)

	uom := model.Uom{
		Id:   "8",
		Name: "Kaca",
	}

	err = uomUseCase.RegisterNewUom(uom)
	if err != nil {
		fmt.Println(err)
	}
}
