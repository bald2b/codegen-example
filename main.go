package main

import "fmt"

//go:generate go run generator/gen.go --file main.go --type OrderStatus

const (
	OrderStatusNew       OrderStatus = 1
	OrderStatusAccepted  OrderStatus = 2
	OrderStatusCancelled OrderStatus = 3
)
const (
	OrderStatusFoo               = 1
	TenderStatusNew TenderStatus = 200
)

type OrderStatus int64
type TenderStatus int64

var statuses = map[string]OrderStatus{"new": OrderStatusNew, "accepted": OrderStatusAccepted, "cancelled": OrderStatusCancelled}

func main() {
	fmt.Println(getStatusesV1())
	fmt.Println("--------------")
	fmt.Println(getStatusesV2())
}

func getStatusesV1() map[string]OrderStatus {
	return statuses
}
