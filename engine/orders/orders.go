package orders

import (
	"github.com/Techbite-sudo/MediTrack-Backend/database"
	"github.com/Techbite-sudo/MediTrack-Backend/graph/model"
)

func CreateOrder(userID string, items []*model.OrderItemInput) (*model.Order, error) {
	var orderItems []*model.OrderItem
	for _, item := range items {
		orderItem := &model.OrderItem{
			MedicationID: item.MedicationID,
			Quantity:     item.Quantity,
			Price:        0, // Set the price accordingly
		}
		orderItems = append(orderItems, orderItem)
	}

	order := model.Order{
		UserID: userID,
		Items:  orderItems,
		// Set other fields as needed
	}
	err := database.DB.Create(&order).Error
	if err != nil {
		return nil, err
	}
	return &order, nil
}


func FetchOrders() ([]*model.Order, error) {
	var orders []*model.Order
	err := database.DB.Find(&orders).Error
	if err != nil {
		return nil, err
	}
	return orders, nil
}

func FetchOrder(id string) (*model.Order, error) {
	var order model.Order
	err := database.DB.Where("id = ?", id).First(&order).Error
	if err != nil {
		return nil, err
	}
	return &order, nil
}
