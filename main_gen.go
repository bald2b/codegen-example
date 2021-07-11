package main

var statusesGenerated = map[string]OrderStatus{}

func getStatusesV2() map[string]OrderStatus {
	return statusesGenerated
}
func init() {
	statusesGenerated["OrderStatusNew"] = OrderStatus(1)
	statusesGenerated["OrderStatusAccepted"] = OrderStatus(2)
	statusesGenerated["OrderStatusCancelled"] = OrderStatus(3)
}
