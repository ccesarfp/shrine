package application

import (
	"context"
	grpcrecovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"runtime/debug"
	"time"
)

type circuitBreaker struct {
	quantity int32
	lastTime time.Time
}

func newCircuitBreaker() *circuitBreaker {
	return &circuitBreaker{
		quantity: 0,
		lastTime: time.Time{},
	}
}

// errorHandler - error handler to be used in grpc middleware
func (cb *circuitBreaker) errorHandler() grpcrecovery.Option {
	return grpcrecovery.WithRecoveryHandlerContext(
		func(ctx context.Context, p interface{}) error {
			errCh <- ctx.Err()
			return status.Errorf(codes.Internal, "%s", p)
		},
	)
}

// verifyError - verify if error persist, case not, reset errors variables
func (cb *circuitBreaker) verifyError() {
	if !i.cb.lastTime.IsZero() {
		timePassed := time.Now().Sub(i.cb.lastTime)
		if timePassed.Minutes() > 5 {
			i.cb.resetError()
		}
	}
}

// incrementError - increment error quantity and set error time
func (cb *circuitBreaker) incrementError() {
	log.Println("Critical Error:", string(debug.Stack()))
	i.cb.quantity++
	i.cb.lastTime = time.Now()
}

// resetError - reset errors variables
func (cb *circuitBreaker) resetError() {
	i.cb.quantity = 0
	i.cb.lastTime = time.Time{}
}
