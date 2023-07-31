package cli

import (
	"fmt"
	"strings"

	"github.com/albar2305/enigma-laundry-apps/model"
	"github.com/albar2305/enigma-laundry-apps/usecase"
	"github.com/albar2305/enigma-laundry-apps/utils/exceptions"
	"github.com/google/uuid"
)

type UomController struct {
	uomUC usecase.UomUseCase
}

func (u *UomController) UomMenuForm() {
	fmt.Println(`
|======== Master  UOM ========|
| 1. Tambah Data              |
| 2. Lihat Data               |
| 3. Detail Data               |
| 4. Update Data              |
| 5. Hapus Data               |
| 6. Kembali ke Menu          |
			`)
	fmt.Print("Pilih Menu (1-5):")
	for {
		var selectedMenu string
		fmt.Scanln(&selectedMenu)
		switch selectedMenu {
		case "1":
			uom := u.uomCreateForm()
			err := u.uomUC.RegisterNewUom(uom)
			exceptions.CheckErr(err)
			return
		case "2":
			uoms, err := u.uomUC.FindAllUom()
			exceptions.CheckErr(err)
			u.uomFindAll(uoms)
			return
		case "3":
			u.uomGetForm()
			return
		case "4":
			uom := u.uomUpdateForm()
			err := u.uomUC.UpdateUom(uom)
			exceptions.CheckErr(err)
			return
		case "5":
			id := u.uomDeleteForm()
			err := u.uomUC.DeleteUom(id)
			exceptions.CheckErr(err)
			return
		case "6":
			return
		default:
			fmt.Println("Menu tidak ditemunkan")
		}
	}
}

func (u *UomController) uomCreateForm() model.Uom {
	var (
		uomName, saveConfirmation string
	)

	fmt.Print("UOM Name: ")
	fmt.Scanln(&uomName)
	fmt.Printf("Name: %s akan disimpan (y/t)", uomName)
	fmt.Scanln(&saveConfirmation)
	if saveConfirmation == "y" {
		uom := model.Uom{
			Id:   uuid.New().String(),
			Name: uomName,
		}
		return uom
	}

	return model.Uom{}
}
func (u *UomController) uomFindAll(uoms []model.Uom) {
	for _, uom := range uoms {
		fmt.Println("UOM List")
		fmt.Printf("ID: %s \n", uom.Id)
		fmt.Printf("Nme: %s \n", uom.Name)
		fmt.Println()
	}
}

func (u *UomController) uomUpdateForm() model.Uom {
	var (
		uomID, uomName, saveConfirmation string
	)

	fmt.Print("UOM ID: ")
	fmt.Scanln(&uomID)
	fmt.Print("UOM Name: ")
	fmt.Scanln(&uomName)
	fmt.Printf("ID: %s,Name: %s akan disimpan (y/t)", uomID, uomName)
	fmt.Scanln(&saveConfirmation)
	if saveConfirmation == "y" {
		uom := model.Uom{
			Id:   uomID,
			Name: uomName,
		}
		return uom
	}

	return model.Uom{}
}

func (u *UomController) uomDeleteForm() string {
	var id string
	fmt.Print("ID :")
	fmt.Scanln(&id)
	return id
}

func (u *UomController) uomGetForm() {
	var id string
	fmt.Print("UOM ID: ")
	fmt.Scanln(&id)
	uom, err := u.uomUC.FindByIdUom(id)
	exceptions.CheckErr(err)
	// fmt.Printf("UOM ID: %s \n", id)
	fmt.Println(strings.Repeat("=", 15))
	fmt.Printf("UOM ID: %s \n", uom.Id)
	fmt.Printf("UOM Name: %s \n", uom.Name)
	fmt.Println()
}

func NewUomController(usecase usecase.UomUseCase) *UomController {
	return &UomController{
		uomUC: usecase,
	}
}
