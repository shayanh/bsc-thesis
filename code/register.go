rpcContract := &contracts.UnaryRPCContract{
    Method: PlaceOrder,
    PreConditions: ...,
    PostConditions: ...
}

var log *logrus.Logger
serverContract := contracts.NewServerContract(log)
serverContract.RegisterUnaryRPCContract(rpcContract)

// Server middleware
srv := grpc.NewServer(grpc.UnaryInterceptor(serverContract.UnaryServerInterceptor()))
srv.serve(...)

// Client middleware
conn, err := grpc.DialContext(ctx, serviceAddress,
    grpc.WithUnaryInterceptor(serverContract.UnaryClientInterceptor()),
)
