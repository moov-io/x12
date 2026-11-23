// Copyright 2026 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package util

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUpdateErrorReason(t *testing.T) {
	t.Run("segment line and stack", func(t *testing.T) {
		err := &ErrorObject{
			Reason:     "bpr segment can't parse all input data",
			Segment:    "BPR*I*1*C*ACH~",
			Originated: []string{"BPR", "Transaction Loop", "TransactionSet", "FunctionalGroup"},
		}
		got := UpdateErrorReason(err)
		require.EqualError(t, got, "bpr segment can't parse all input data, error segment line: 'BPR*I*1*C*ACH~', error stack: 'functionalgroup->transactionset->transaction loop->bpr'")
	})

	t.Run("stack only", func(t *testing.T) {
		err := &ErrorObject{
			Reason:     "segment(amt) don't accept specified rule(tds)",
			Originated: []string{"AMT", "Summary"},
		}
		got := UpdateErrorReason(err)
		require.EqualError(t, got, "segment(amt) don't accept specified rule(tds), error stack: 'summary->amt'")
	})

	t.Run("segment line only", func(t *testing.T) {
		err := &ErrorObject{
			Reason:     "ge segment contains invalid code",
			Segment:    "ST*810*0002~",
			Originated: []string{"GE"},
		}
		got := UpdateErrorReason(err)
		require.EqualError(t, got, "ge segment contains invalid code, error segment line: 'ST*810*0002~'")
	})

	t.Run("other errors", func(t *testing.T) {
		require.NoError(t, UpdateErrorReason(nil))

		other := errors.New("other")
		require.Equal(t, other, UpdateErrorReason(other))
	})
}
