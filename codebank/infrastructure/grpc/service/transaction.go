package service

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/hrsniper/imersao-fullstack-fullcycle-3/dto"
	"github.com/hrsniper/imersao-fullstack-fullcycle-3/infrastructure/grpc/pb"
	"github.com/hrsniper/imersao-fullstack-fullcycle-3/usecase"
)

type TransactionService struct {
	ProcessTransactionUseCase usecase.UseCaseTransaction
	pb.UnimplementedPaymentServiceServer
}

func NewTransactionService() *TransactionService {
	return &TransactionService{}
}

// method of TransactionService struct
func (transactionService *TransactionService) Payment(ctx context.Context, paymentRequest *pb.PaymentRequest) (*empty.Empty, error) {
	transactionDto := dto.Transaction{
		Name:            paymentRequest.GetCreditCard().GetName(),
		Number:          paymentRequest.CreditCard.GetNumber(),
		ExpirationMonth: paymentRequest.GetCreditCard().GetExpirationMonth(),
		ExpirationYear:  paymentRequest.GetCreditCard().GetExpirationYear(),
		CVV:             paymentRequest.GetCreditCard().GetCvv(),
		Amount:          paymentRequest.GetAmount(),
		Store:           paymentRequest.GetStore(),
		Description:     paymentRequest.GetDescription(),
	}

	transaction, err := transactionService.ProcessTransactionUseCase.ProcessTransaction(transactionDto)

	if err != nil {
		return &empty.Empty{}, status.Error(codes.FailedPrecondition, err.Error())
	}

	if transaction.Status != "approved" {
		return &empty.Empty{}, status.Error(codes.FailedPrecondition, "transaction rejected by the bank")
	}

	return &empty.Empty{}, nil
}
