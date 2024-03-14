package application

import (
	"context"
	statusEnum "github.com/ccesarfp/shrine/internal/enum/status"
	"github.com/ccesarfp/shrine/internal/errors/circuit_open"
	grpcrecovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"runtime/debug"
	"time"
)

type circuitBreaker struct {
	verifyingError  bool
	maxErrors       int32     // maximum amount of error until the circuit breaker Opens
	errorsQuantity  int32     // current error quantity
	lastTime        time.Time // time of last error occurred
	timeoutDuration float64   // time to check if the error persists in minutes
	status          uint8
}

func newCircuitBreaker() *circuitBreaker {
	return &circuitBreaker{
		verifyingError:  false,
		errorsQuantity:  0,
		lastTime:        time.Time{},
		maxErrors:       3,
		timeoutDuration: 5,
		status:          statusEnum.Close,
	}
}

// errorHandler - error handler to be used in grpc middleware
func (cb *circuitBreaker) errorHandler() grpcrecovery.Option {
	return grpcrecovery.WithRecoveryHandlerContext(
		func(ctx context.Context, p interface{}) error {
			log.Println("[Circuit Breaker] Critical Error:", string(debug.Stack()))
			errCh <- ctx.Err()
			return status.Errorf(codes.Internal, "%s", p)
		},
	)
}

// verifyError - verify if error persist, case not, reset errors variables
func (cb *circuitBreaker) verifyError() {
	if !i.cb.lastTime.IsZero() {
		timePassed := time.Now().Sub(i.cb.lastTime)
		if timePassed.Minutes() >= cb.timeoutDuration {
			i.cb.changeStatus(statusEnum.HalfOpen)
		}
	}
}

// incrementError - increment error quantity and set error time
func (cb *circuitBreaker) incrementError() {
	i.cb.errorsQuantity++
	i.cb.lastTime = time.Now()
}

// resetError - reset errors variables
func (cb *circuitBreaker) closeCircuit() {
	i.cb.errorsQuantity = 0
	i.cb.lastTime = time.Time{}
	i.cb.verifyingError = false
	i.cb.changeStatus(statusEnum.Close)
}

// changeStatus - change circuit breaker status
func (cb *circuitBreaker) changeStatus(newStatus uint8) {
	if i.cb.status != newStatus {
		oldStatus := i.cb.status

		switch newStatus {
		case statusEnum.Close:
			i.cb.status = statusEnum.Close
		case statusEnum.HalfOpen:
			i.cb.status = statusEnum.HalfOpen
		case statusEnum.Open:
			i.cb.status = statusEnum.Open
		default:
			i.cb.status = 3
		}

		log.Println("[Circuit Breaker] Changing Status from:", statusEnum.String(oldStatus), "- to:", statusEnum.String(i.cb.status))
	}
}

// circuitBreakerInterceptor - verify circuit breaker status and handled the request based on this status
func (cb *circuitBreaker) circuitBreakerInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {

	c := circuit_open.Error{}
	if i.cb.status == statusEnum.Open {
		return nil, status.Error(codes.Unavailable, c.Error())
	}

	if i.cb.status == statusEnum.HalfOpen && i.cb.verifyingError == true {
		return nil, status.Error(codes.Unavailable, c.Error())
	}

	if i.cb.status == statusEnum.HalfOpen && i.cb.verifyingError == false {
		i.cb.verifyingError = true
	}

	resp, err := handler(ctx, req)

	if i.cb.status == statusEnum.HalfOpen && i.cb.verifyingError == true && err != nil {
		i.cb.verifyingError = false
		i.cb.changeStatus(statusEnum.Open)
	}

	if i.cb.status == statusEnum.HalfOpen && i.cb.verifyingError == true && err == nil {
		i.cb.closeCircuit()
	}

	return resp, err
}
