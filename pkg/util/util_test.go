// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package util

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestManageFieldByIndex(t *testing.T) {

	type TEST struct {
		INDEX1 string `index:"INDEX1"`
		INDEX2 int    `index:"INDEX2"`
	}

	t.Run("Setting using non pointed object", func(t *testing.T) {
		sample := TEST{}

		err := SetFieldByIndex(sample, "INDEX1", "INDEX1")
		require.Error(t, err)
		require.Equal(t, "could not set data", err.Error())

	})

	t.Run("Setting using pointed object", func(t *testing.T) {
		sample := TEST{}

		err := SetFieldByIndex(&sample, "INDEX1", "INDEX1")
		require.NoError(t, err)

	})

	t.Run("Setting using unknown index", func(t *testing.T) {
		sample := TEST{}

		err := SetFieldByIndex(&sample, "UNKNOWN", "INDEX1")
		require.Error(t, err)
		require.Equal(t, "unable to find matched index", err.Error())

	})

	t.Run("Setting using unmatched data type", func(t *testing.T) {
		sample := TEST{}

		err := SetFieldByIndex(&sample, "INDEX1", 111)
		require.Error(t, err)
		require.Equal(t, "doesn't match setting type", err.Error())

	})

	t.Run("Getting using non pointed object", func(t *testing.T) {
		sample := TEST{}

		SetFieldByIndex(&sample, "INDEX1", "INDEX1")
		SetFieldByIndex(&sample, "INDEX2", 111)

		got := GetFieldByIndex(sample, "INDEX1")
		require.Equal(t, "INDEX1", got)

		got = GetFieldByIndex(sample, "INDEX2")
		require.Equal(t, 111, got)
	})

	t.Run("Getting using pointed object", func(t *testing.T) {
		sample := TEST{}

		SetFieldByIndex(&sample, "INDEX1", "INDEX1")
		SetFieldByIndex(&sample, "INDEX2", 111)

		got := GetFieldByIndex(&sample, "INDEX1")
		require.Equal(t, "INDEX1", got)

		got = GetFieldByIndex(&sample, "INDEX2")
		require.Equal(t, 111, got)
	})

	t.Run("Checking duplicated control number", func(t *testing.T) {

		numbers := []string{
			"number1",
			"number2",
			"number3",
		}

		exist, got := GetDuplicateControlNumber(numbers)
		require.Equal(t, false, exist)
		require.Equal(t, "", got)

		numbers = append(numbers, "number1")
		exist, got = GetDuplicateControlNumber(numbers)
		require.Equal(t, true, exist)
		require.Equal(t, "number1", got)
	})
}
