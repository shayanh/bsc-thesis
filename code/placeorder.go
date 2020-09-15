func PlaceOrder(req *PlaceOrderRequest) (*PlaceOrderResponse, error) {
    cost := 0
    for _, item := range req.Items {
        cost += item.Cost
    }
    _, err := PaymentService.Charge(&ChargeRequest{
        Amount: cost, 
        CreditCard: req.CreditCard})
    if err != nil {
        return err
    }
    _, err := ShippingService.ShipOrder(&ShipOrderRequest{
        Address: req.Address,
        Items: req.Items})
    if err != nil {
        return err
    }
    orderID := generateNewID()
    return &PlaceOrderResponse{OrderID: orderID}, nil
}
