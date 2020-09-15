rpcContract.PreConditions: []contracts.Condition{
    // Requires validity of user's email
    func(req PlaceOrderRequest) error {
        if !isValidEmail(req.Email) {
            return errors.New("email address is not valid")
        }
        return nil
    },
    // Requires user authentication
    func(req *pb.PlaceOrderRequest) error {
        if req.UserId == "" {
            return errors.New("user must be authenticated")
        }
        return nil
    },
}
