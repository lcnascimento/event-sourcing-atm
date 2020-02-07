package server

import (
	"context"
	"net"

	"github.com/golang/protobuf/proto"
	"github.com/google/uuid"
	"google.golang.org/grpc"

	pb "github.com/lcnascimento/event-sourcing-atm/proto/impl"

	"github.com/lcnascimento/event-sourcing-atm/infra"
	"github.com/lcnascimento/event-sourcing-atm/infra/errors"

	"github.com/lcnascimento/event-sourcing-atm/domain"
	command "github.com/lcnascimento/event-sourcing-atm/domain/command"
)

// ServiceInput ...
type ServiceInput struct {
	Stream infra.EventStreamPublisher

	Accounts     command.AccountsProvider
	Transactions command.TransactionsProvider
}

// Service ...
type Service struct {
	in ServiceInput
}

// NewService ...
func NewService(in ServiceInput) (*Service, *infra.Error) {
	return &Service{in: in}, nil
}

// CreateAccountCommand ...
func (s Service) CreateAccountCommand(ctx context.Context, req *pb.CreateAccountRequest) (*pb.CreateAccountResponse, error) {
	const opName infra.OpName = "command.server.Create"

	if err := s.in.Accounts.Insert(ctx, domain.Account{}); err != nil {
		return &pb.CreateAccountResponse{Error: errors.ToProtoError(err)}, errors.New(opName, err)
	}

	dataEncoded, err := proto.Marshal(&pb.AccountCreated{
		Account: &pb.Account{
			Id:       "123",
			Password: "321",
			Owner: &pb.User{
				Id:    "789",
				Name:  "Luís Nascimento",
				Cpf:   "123.456.789-10",
				Email: "contact@lcnascimento.me",
			},
		},
	})
	if err != nil {
		return &pb.CreateAccountResponse{
			Error: errors.ToProtoError(errors.New(ctx, opName, err)),
		}, errors.New(opName, err)
	}

	id, _ := uuid.NewRandom()

	event, err := proto.Marshal(&pb.Event{
		Id:   id.String(),
		Type: pb.EventType_ACCOUNT_CREATED,
		Data: dataEncoded,
	})
	if err != nil {
		return &pb.CreateAccountResponse{
			Error: errors.ToProtoError(errors.New(ctx, opName, err)),
		}, errors.New(opName, err)
	}

	if err := s.in.Stream.Publish(ctx, "accounts", event); err != nil {
		return &pb.CreateAccountResponse{Error: errors.ToProtoError(err)}, errors.New(opName, err)
	}

	return &pb.CreateAccountResponse{}, nil
}

// DeleteAccountCommand ...
func (s Service) DeleteAccountCommand(ctx context.Context, req *pb.DeleteAccountRequest) (*pb.DeleteAccountResponse, error) {
	const opName infra.OpName = "command.server.Delete"

	if err := s.in.Accounts.Remove(ctx, infra.ObjectID(req.AccountId)); err != nil {
		return &pb.DeleteAccountResponse{}, errors.New(opName, err)
	}

	return &pb.DeleteAccountResponse{}, nil
}

// DepositCommand ...
func (s Service) DepositCommand(ctx context.Context, req *pb.DepositRequest) (*pb.DepositResponse, error) {
	const opName infra.OpName = "command.server.Deposit"

	if err := s.in.Transactions.Deposit(ctx, infra.ObjectID(req.AccountId), req.Amount); err != nil {
		return &pb.DepositResponse{}, errors.New(opName, err)
	}

	return &pb.DepositResponse{}, nil
}

// WithdrawCommand ...
func (s Service) WithdrawCommand(ctx context.Context, req *pb.WithdrawRequest) (*pb.WithdrawResponse, error) {
	const opName infra.OpName = "command.server.Withdraw"

	if err := s.in.Transactions.Withdraw(ctx, infra.ObjectID(req.AccountId), req.Amount); err != nil {
		return &pb.WithdrawResponse{}, errors.New(opName, err)
	}

	return &pb.WithdrawResponse{}, nil
}

// TransferCommand ...
func (s Service) TransferCommand(ctx context.Context, req *pb.TransferRequest) (*pb.TranseResponse, error) {
	const opName infra.OpName = "command.server.Transfer"

	err := s.in.Transactions.Transfer(
		ctx,
		infra.ObjectID(req.SourceAccountId),
		infra.ObjectID(req.DeliveryAccountId),
		req.Amount,
	)
	if err != nil {
		return &pb.TranseResponse{}, errors.New(opName, err)
	}

	return &pb.TranseResponse{}, nil
}

// LoanCommand ...
func (s Service) LoanCommand(ctx context.Context, req *pb.LoanRequest) (*pb.LoanResponse, error) {
	const opName infra.OpName = "command.server.Loan"

	if err := s.in.Transactions.Withdraw(ctx, infra.ObjectID(req.AccountId), req.Amount); err != nil {
		return &pb.LoanResponse{}, errors.New(opName, err)
	}

	return &pb.LoanResponse{}, nil
}

// Run ...
func (s Service) Run(ctx context.Context) <-chan *infra.Error {
	const opName infra.OpName = "command.server.Run"

	ch := make(chan *infra.Error)

	lis, err := net.Listen("tcp", "localhost:8081")
	if err != nil {
		ch <- errors.New(ctx, opName, err)
		close(ch)
	}

	go func() {
		grpcServer := grpc.NewServer()
		pb.RegisterAccountsCommandServiceServer(grpcServer, s)
		if err := grpcServer.Serve(lis); err != nil {
			ch <- errors.New(ctx, opName, err)
		}
	}()

	return ch
}
