// Copyright 2020 PingCAP, Inc. Licensed under Apache-2.0.

package utils

import (
	"context"

	"github.com/pingcap/br/pkg/retry"
)

// RetryableFunc presents a retryable operation.
type RetryableFunc = retry.RetryableFunc

// Backoffer implements a backoff policy for retrying operations.
type Backoffer = retry.Backoffer

// WithRetry retries a given operation with a backoff policy.
//
// Returns nil if `retryableFunc` succeeded at least once. Otherwise, returns a
// multierr containing all errors encountered.
func WithRetry(
	ctx context.Context,
	retryableFunc RetryableFunc,
	backoffer Backoffer,
) error {
	return retry.WithRetry(ctx, retryableFunc, backoffer)
}
