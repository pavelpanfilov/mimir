// SPDX-License-Identifier: AGPL-3.0-only

package util

import (
	"errors"

	"github.com/prometheus/prometheus/model/histogram"

	"github.com/grafana/mimir/pkg/util/globalerror"
)

func ConvertHistogramErrorToGlobalError(err error) globalerror.ID {
	switch {
	case errors.Is(err, histogram.ErrHistogramCountMismatch):
		return globalerror.NativeHistogramCountMismatch
	case errors.Is(err, histogram.ErrHistogramCountNotBigEnough):
		return globalerror.NativeHistogramCountNotBigEnough
	case errors.Is(err, histogram.ErrHistogramNegativeBucketCount):
		return globalerror.NativeHistogramNegativeBucketCount
	case errors.Is(err, histogram.ErrHistogramSpanNegativeOffset):
		return globalerror.NativeHistogramSpanNegativeOffset
	case errors.Is(err, histogram.ErrHistogramSpansBucketsMismatch):
		return globalerror.NativeHistogramSpansBucketsMismatch
	case errors.Is(err, histogram.ErrHistogramCustomBucketsMismatch):
		return globalerror.NativeHistogramCustomBucketsMismatch
	case errors.Is(err, histogram.ErrHistogramCustomBucketsInvalid):
		return globalerror.NativeHistogramCustomBucketsInvalid
	case errors.Is(err, histogram.ErrHistogramCustomBucketsInfinite):
		return globalerror.NativeHistogramCustomBucketsInfinite
	default:
		return ""
	}
}
