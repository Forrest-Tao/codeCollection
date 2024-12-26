package main

import (
	"context"
	"fmt"
	order "forrest/codeCollection/grpc/grpc-on-k8s/client/pb"
	payment "forrest/codeCollection/grpc/grpc-on-k8s/client/pb"
	user "forrest/codeCollection/grpc/grpc-on-k8s/client/pb"
	"google.golang.org/grpc"
	"log"
	"sync"
	"time"
)

type Client struct {
	userClient    user.UserServiceClient
	orderClient   order.OrderServiceClient
	paymentClient payment.PaymentServiceClient
	userCon       *grpc.ClientConn
	OrderCon      *grpc.ClientConn
	paymentCon    *grpc.ClientConn
}

func NewClient() *Client {
	// 连接到各个服务，使用 Kubernetes 服务名称
	userConn, err := grpc.Dial("user-service.default.svc.cluster.local:50001", grpc.WithInsecure()) // 使用服务的 DNS 名称
	if err != nil {
		log.Fatalf("failed to connect to UserService: %v", err)
	}
	orderConn, err := grpc.Dial("order-service.default.svc.cluster.local:50002", grpc.WithInsecure()) // 使用服务的 DNS 名称
	if err != nil {
		log.Fatalf("failed to connect to OrderService: %v", err)
	}
	paymentConn, err := grpc.Dial("payment-service.default.svc.cluster.local:50003", grpc.WithInsecure()) // 使用服务的 DNS 名称
	if err != nil {
		log.Fatalf("failed to connect to PaymentService: %v", err)
	}

	// 创建客户端
	return &Client{
		userClient:    user.NewUserServiceClient(userConn),
		orderClient:   order.NewOrderServiceClient(orderConn),
		paymentClient: payment.NewPaymentServiceClient(paymentConn),
		userCon:       userConn,
		OrderCon:      orderConn,
		paymentCon:    paymentConn,
	}
}

func (c *Client) GetAggregatedData(userID int32) (string, error) {
	var wg sync.WaitGroup
	wg.Add(3)

	var userInfo *user.UserInfo
	var orders *order.Orders
	var payments *payment.Payments
	var errUser, errOrder, errPayment error

	// 并发调用多个服务
	go func() {
		defer wg.Done()
		userInfo, errUser = c.userClient.GetUserInfo(context.Background(), &user.UserRequest{UserId: userID})
	}()

	go func() {
		defer wg.Done()
		orders, errOrder = c.orderClient.GetOrders(context.Background(), &order.OrderRequest{UserId: userID})
	}()

	go func() {
		defer wg.Done()
		payments, errPayment = c.paymentClient.GetPayments(context.Background(), &payment.PaymentRequest{UserId: userID})
	}()

	// 等待所有 goroutine 完成
	wg.Wait()

	// 如果有任何错误，返回错误
	if errUser != nil {
		return "", fmt.Errorf("failed to get user info: %v", errUser)
	}
	if errOrder != nil {
		return "", fmt.Errorf("failed to get orders: %v", errOrder)
	}
	if errPayment != nil {
		return "", fmt.Errorf("failed to get payments: %v", errPayment)
	}

	// 聚合数据
	result := fmt.Sprintf("User: %s\nOrders: %v\nPayments: %v", userInfo.UserName, orders.OrderList, payments.PaymentList)
	return result, nil
}

func main() {
	client := NewClient()
	defer client.userCon.Close()
	defer client.OrderCon.Close()
	defer client.paymentCon.Close()

	// 获取聚合数据
	userID := int32(1)
	result, err := client.GetAggregatedData(userID)
	if err != nil {
		log.Fatalf("failed to get aggregated data: %v", err)
	}

	fmt.Println(result)
	time.Sleep(3 * time.Second)
}
