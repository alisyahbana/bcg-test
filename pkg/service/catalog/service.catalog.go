package catalog

import (
	"fmt"
	"github.com/alisyahbana/bcg-test/pkg/service/catalog/data"
	"math"
	"strings"
)

type CatalogService struct {
	data data.MysqlCatalogData
}

func New() CatalogService {
	return CatalogService{
		data: data.MysqlCatalogData{},
	}
}

type PurchasePayload struct {
	Items []string `json:"items"`
}

type PurchaseResponse struct {
	TotalPrice float64 `json:"total_price"`
}

func (s CatalogService) Purchase(payload PurchasePayload) (*PurchaseResponse, error) {
	var items []data.Item
	var totalPrice float64
	var countGoogleHome int
	var purchasedMacbook bool
	var countAlexa int
	for _, keyword := range payload.Items {
		// if purchased macbook
		if strings.ToLower(keyword) == "macbook pro" {
			purchasedMacbook = true
		}
		item, err := s.data.GetItem(keyword)
		if err != nil {
			return nil, err
		}
		if item == nil {
			return nil, fmt.Errorf("item '%s' not found", keyword)
		}
		if item.SKU == "120P90" {
			countGoogleHome += 1
		}
		if item.SKU == "A304SD" {
			countAlexa += 1
		}
		price := item.Price
		// if purchased macbook then free raspberry
		if purchasedMacbook && item.SKU == "234234" {
			price = 0
		}
		totalPrice += price
		items = append(items, *item)
	}

	if countGoogleHome%3 == 0 {
		discount := 49.99
		numberDiscount := countGoogleHome / 3
		totalDiscount := (discount * float64(numberDiscount))
		totalPrice = totalPrice - totalDiscount
	}

	if countAlexa >= 3 {
		totalPrice = totalPrice - float64(totalPrice*10/100)
	}

	return &PurchaseResponse{
		TotalPrice: math.Round(totalPrice*100) / 100,
	}, nil
}
