package cache

import (
	"encoding/json"
	"fmt"

	"github.com/fyerfyer/trade-dependency/dto/customer"
	"github.com/fyerfyer/trade-dependency/dto/order"
)

func LookUpOrderInCache(cache Cache, id uint64, status string) (bool, *order.OrderDTO) {
	var order *order.OrderDTO
	if cache.Exists(GetOrderKey(id, status)) {
		data, err := cache.Get(GetOrderKey(id, status))
		if err != nil {
			return false, nil
		}
		json.Unmarshal(data, &order)
	} else {
		return false, nil
	}

	return true, order
}

func LookUpCustomerInCache(cache Cache, id uint64) (bool, *customer.CustomerDTO) {
	var customer *customer.CustomerDTO
	if cache.Exists(GetCustomerKey(id)) {
		data, err := cache.Get(GetCustomerKey(id))
		if err != nil {
			return false, nil
		}
		json.Unmarshal(data, &customer)
	} else {
		return false, nil
	}

	return true, customer
}

func GetCustomerKey(id uint64) string {
	return fmt.Sprintf("Customer_%v", id)
}

func GetOrderKey(id uint64, status string) string {
	return fmt.Sprintf("Order_%v_%v", id, status)
}
