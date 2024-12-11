package e

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	DUPLICATE_CUSTOMER_ERROR     = status.New(codes.AlreadyExists, "duplicate customer").Err()
	CUSTOMER_NOT_FOUND           = status.New(codes.NotFound, "customer not found").Err()
	PAYMENT_NOT_FOUND            = status.New(codes.NotFound, "payment not found").Err()
	ORDER_NOT_FOUND              = status.New(codes.NotFound, "order not found").Err()
	ORDER_INVALID                = status.New(codes.FailedPrecondition, "order status invalid").Err()
	BALANCE_INSUFFICIENT         = status.New(codes.FailedPrecondition, "insufficient balance").Err()
	CUSTOMER_INACTIVE            = status.New(codes.FailedPrecondition, "customer is inactive").Err()
	FAILED_TO_STORE_DB_ERROR     = status.New(codes.Internal, "failed to store into database").Err()
	FAILED_TO_STORE_CACHE_ERROR  = status.New(codes.Internal, "failed to store into cache").Err()
	FAILED_TO_UPDATE_DB_ERROR    = status.New(codes.Internal, "failed to update database").Err()
	FAILED_TO_UPDATE_CACHE_ERROR = status.New(codes.Internal, "failed to update cache").Err()
	FAILED_TLS_CERT              = status.New(codes.Internal, "failed to append CA certs").Err()
)
