package service

import (
	"context"

	"github.com/filipe-freitas/codecoin/dto"
	"github.com/filipe-freitas/codecoin/infrastructure/grpc/pb"
	"github.com/filipe-freitas/codecoin/usecase"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type TransactionService struct {
	ProcessTransactionUseCase usecase.UseCaseTransaction
	pb.UnimplementedPaymentServiceServer
}

func NewTransactionService() *TransactionService {
	return &TransactionService{}
}

func (t *TransactionService) Payment(ctx context.Context, in *pb.PaymentRequest) (*empty.Empty, error) {
	transactionDTO := dto.Transaction{
		Name:            in.GetCreditCard().GetName(),
		Number:          in.GetCreditCard().GetNumber(),
		ExpirationMonth: uint16(in.GetCreditCard().GetExpirationMonth()),
		ExpirationYear:  uint16(in.GetCreditCard().GetExpirationYear()),
		CVV:             uint16(in.GetCreditCard().GetCvv()),
		Amount:          in.GetAmount(),
		Store:           in.GetStore(),
		Description:     in.GetDescription(),
	}

	transaction, err := t.ProcessTransactionUseCase.ProcessTransaction(transactionDTO)
	if err != nil {
		return &empty.Empty{}, status.Error(codes.FailedPrecondition, err.Error())
	}
	if transaction.Status != "approved" {
		return &empty.Empty{}, status.Error(codes.FailedPrecondition, "Transaction rejected by the bank")
	}
	return &empty.Empty{}, nil
}
