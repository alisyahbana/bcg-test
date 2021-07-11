package catalog

import (
	"fmt"
	"github.com/alisyahbana/bcg-test/pkg/service/catalog/data"
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

	var purchasedMacbook bool
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

		price := item.Price

		// if purchased macbook then free raspberry
		if purchasedMacbook && item.SKU == "234234" {
			price = 0
		}

		totalPrice += price

		items = append(items, *item)
	}

	return &PurchaseResponse{
		TotalPrice: totalPrice,
	}, nil
}
