package paymentRepository

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/phetployst/sekai-shop-microservices/modules/item/itemPb"
	"github.com/phetployst/sekai-shop-microservices/pkg/grpccon"
	"github.com/phetployst/sekai-shop-microservices/pkg/jwtauth"
	"go.mongodb.org/mongo-driver/mongo"
)

type (
	PaymentRepositoryService interface {
		FindItemsInIds(pctx context.Context, grpcUrl string, req *itemPb.FindItemsInIdsReq) (*itemPb.FindItemsInIdsRes, error)
	}

	paymentRepository struct {
		db *mongo.Client
	}
)

func NewPaymentRepository(db *mongo.Client) PaymentRepositoryService {
	return &paymentRepository{db}
}

func (r *paymentRepository) paymentDbConn(pctx context.Context) *mongo.Database {
	return r.db.Database("payment_db")
}

func (r *paymentRepository) FindItemsInIds(pctx context.Context, grpcUrl string, req *itemPb.FindItemsInIdsReq) (*itemPb.FindItemsInIdsRes, error) {
	ctx, cancel := context.WithTimeout(pctx, 30*time.Second)
	defer cancel()

	jwtauth.SetApiKeyInContext(&ctx)
	conn, err := grpccon.NewGrpcClient(grpcUrl)
	if err != nil {
		log.Printf("Error: gRPC connection failed: %s", err.Error())
		return nil, errors.New("error: gRPC connection failed")
	}

	result, err := conn.Item().FindItemsInIds(ctx, req)
	if err != nil {
		log.Printf("Error: FindItemsInIds failed: %s", err.Error())
		return nil, errors.New("error: items not found")
	}

	if result == nil {
		log.Printf("Error: FindItemsInIds failed: %s", err.Error())
		return nil, errors.New("error: items not found")
	}

	if len(result.Items) == 0 {
		log.Printf("Error: FindItemsInIds failed: %s", err.Error())
		return nil, errors.New("error: items not found")
	}

	return result, nil
}
