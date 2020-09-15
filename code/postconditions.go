rpcContract.PostConditions: []contracts.Condition{
    // Ensure a successful place order call, has a successful ship order call
    func(resp *PlaceOrderResponse, respErr error, req *PlaceOrderRequest, calls contracts.RPCCallHistory) error {
        if respErr != nil {
            return nil
        }
        paymentCalls := calls.Filter(PaymentServiceServer.Charge)
        if len(paymentCalls) < 1 {
            return errors.New("no call to payment service")
        }
        paymentCall := paymentCalls[0]
        if paymentCall.Error != nil || paymentCall.Response.(*pb.ChargeResponse).GetTransactionId() == "" {
            return errors.New("invalid response from payment service")
        }
        return nil
    },
    // Ensure payment call happens before shipping call
    func(resp *PlaceOrderResponse, respErr error, req *PlaceOrderRequest, calls contracts.RPCCallHistory) error {
        if respErr != nil {
            return nil
        }
        shippingCalls := calls.Filter(ShippingServiceServer.ShipOrder)
        paymentCalls := calls.Filter(PaymentServiceServer.Charge)
        if len(paymentCalls) < 1 || len(shippingCalls) < 1 {
            return errors.New("no call to payment or shipping service")
        }
        paymentCall := paymentCalls[0]
        shippingCall := shippingCalls[0]
        if paymentCall.Order > shippingCall.Order {
            return errors.New("payment must called before shipping")
        }
        return nil
    },
}
