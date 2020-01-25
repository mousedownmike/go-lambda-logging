// Illustrates structured logging with logrus and AWS Cloudwatch Metrics.
package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/mousedownmike/go-lambda-logging/pkg/mdc"
	log "github.com/sirupsen/logrus"
	"math/rand"
	"os"
	"time"
)

var (
	r = rand.New(rand.NewSource(time.Now().UnixNano()))
)


func main() {
	log.SetFormatter(&log.JSONFormatter{})
	if os.Getenv("AWS_LAMBDA_FUNCTION_NAME") != "" {
		lambda.Start(handle)
	} else {
		handle(context.Background())
	}
}

func handle(ctx context.Context) {
	// Create a parent Context with an empty Diagnostic Map
	ctx = mdc.Context(ctx)
	var err error
	if checkDiff(ctx) {
		err = processDiffs(ctx)
	}
	if err != nil {
		log.WithFields(mdc.Diagnostics(ctx)).Error(err)
	} else {
		log.WithFields(mdc.Diagnostics(ctx)).Info("success")
	}
}

func checkDiff(ctx context.Context) bool {
	var fakeDiff bool
	if r.Intn(2) == 0 {
		fakeDiff = false
	} else {
		fakeDiff = true
	}
	mdc.WithDiagnostic(ctx, "hasDiff", fakeDiff)
	return fakeDiff
}

func processDiffs(ctx context.Context) error {
	adds := r.Intn(30)
	subtracts := r.Intn(20)
	fakeError := r.Intn(4)
	if fakeError == 0 {
		return fmt.Errorf("%d causes a fake error", fakeError)
	}
	mdc.WithDiagnostic(ctx, "additions", adds)
	mdc.WithDiagnostic(ctx, "subtractions", subtracts)
	return nil
}
