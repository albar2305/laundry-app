package cli

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/albar2305/enigma-laundry-apps/model"
	"github.com/albar2305/enigma-laundry-apps/model/dto"
	"github.com/albar2305/enigma-laundry-apps/usecase"
	"github.com/albar2305/enigma-laundry-apps/utils/exceptions"
	"github.com/google/uuid"
)

type ProductController struct {
	productUC usecase.ProductUseCase
}

func (p *ProductController) HandlerMainForm() {
	fmt.Println(`
|==== Master Product ====|
| 1. Tambah Data         |
| 2. Lihat Data          |
| 3. Detail Data	     |
| 4. Update Data         |
| 5. Hapus Data          |
| 6. Kembali ke Menu     |
			     `)
	fmt.Print("Pilih Menu (1-6): ")
	for {
		var selectedMenu string
		fmt.Scanln(&selectedMenu)
		switch selectedMenu {
		case "1":
			product := p.createHandlerForm()
			err := p.productUC.RegisterNewProduct(product)
			exceptions.CheckErr(err)
			return
		case "2":
			requestPaging := dto.PaginationParam{
				Page: 1,
			}
			products, paging, err := p.productUC.FindAllProduct(requestPaging)
			exceptions.CheckErr(err)
			p.findAllHandlerForm(products, paging)
			return

		case "3":
			p.getHandleForm()
			return
		case "4":
			product := p.updateHandlerForm()
			err := p.productUC.UpdateProduct(product)
			exceptions.CheckErr(err)
			return
		case "5":
			id := p.deleteHandleForm()
			err := p.productUC.DeleteProduct(id)
			exceptions.CheckErr(err)
			return
		case "6":
			return
		default:
			fmt.Println("Menu tidak ditemukan!")
		}
	}
}

func (p *ProductController) createHandlerForm() model.Product {
	var (
		id, name, price, uomId, saveConfirmation string
	)
	fmt.Print("Product Name: ")
	fmt.Scanln(&name)

	fmt.Print("Product Price: ")
	fmt.Scanln(&price)

	fmt.Print("Product UOM ID: ")
	fmt.Scanln(&uomId)
	id = uuid.New().String()
	priceConv, _ := strconv.Atoi(price)
	fmt.Printf("Product Id: %s, Name: %s, Price: %d, UOM ID: %s akan disimpan (y/t)", id, name, priceConv, uomId)
	fmt.Scanln(&saveConfirmation)
	if saveConfirmation == "y" {
		product := model.Product{
			Id:    id,
			Name:  name,
			Price: priceConv,
			Uom:   model.Uom{Id: uomId},
		}
		return product
	}
	return model.Product{}
}

func (p *ProductController) findAllHandlerForm(products []model.Product, paging dto.Paging) {
	fmt.Println("Product List")
	for _, product := range products {
		fmt.Printf("ID: %s \n", product.Id)
		fmt.Printf("Name: %s \n", product.Name)
		fmt.Printf("Price: %d \n", product.Price)
		fmt.Printf("UOM ID: %s \n", product.Uom.Id)
		fmt.Printf("UOM Name: %s \n", product.Uom.Name)
		fmt.Println()
	}
	fmt.Println("Paging: ")
	fmt.Printf("Page: %d \n", paging.Page)
	fmt.Printf("RowsPerPage: %d \n", paging.RowsPerPage)
	fmt.Printf("TotalPages: %d \n", paging.TotalPages)
	fmt.Printf("TotalRows: %d \n", paging.TotalRows)
}

func (p *ProductController) updateHandlerForm() model.Product {
	var (
		id, name, price, uomId, saveConfirmation string
	)
	fmt.Print("Product ID: ")
	fmt.Scanln(&id)

	fmt.Print("Product Name: ")
	fmt.Scanln(&name)

	fmt.Print("Product Price: ")
	fmt.Scanln(&price)

	fmt.Print("Product UOM ID: ")
	fmt.Scanln(&uomId)
	priceConv, _ := strconv.Atoi(price)
	fmt.Printf("Product Id: %s, Name: %s, Price: %d, UOM ID: %s akan disimpan (y/t)", id, name, priceConv, uomId)
	fmt.Scanln(&saveConfirmation)
	if saveConfirmation == "y" {
		product := model.Product{
			Id:    id,
			Name:  name,
			Price: priceConv,
			Uom:   model.Uom{Id: uomId},
		}
		return product
	}
	return model.Product{}
}

func (p *ProductController) deleteHandleForm() string {
	var id string
	fmt.Print("UOM ID: ")
	fmt.Scanln(&id)
	return id
}

func (p *ProductController) getHandleForm() {
	var id string
	fmt.Print("UOM ID: ")
	fmt.Scanln(&id)
	product, err := p.productUC.FindByIdProduct(id)
	exceptions.CheckErr(err)
	fmt.Printf("UOM ID %s \n", id)
	fmt.Println(strings.Repeat("=", 15))
	fmt.Printf("UOM ID: %s \n", product.Id)
	fmt.Printf("UOM Name: %s \n", product.Name)
	fmt.Printf("Price: %d \n", product.Price)
	fmt.Printf("UOM ID: %s \n", product.Uom.Id)
	fmt.Printf("UOM Name: %s \n", product.Uom.Name)
	fmt.Println()
}

func NewProductController(usecase usecase.ProductUseCase) *ProductController {
	return &ProductController{productUC: usecase}
}
