package business_test

import (
	"context"
	"errors"
	business "fizz-buzz-gin/pkg/business"
	"reflect"
	"testing"
	"time"
)

func TestFizzBuzz(t *testing.T) {
	for i, tc := range getTestData() {
		// execute the fizzbuzz function
		result, err := business.FizzBuzz(tc.ctx, tc.int1, tc.int2, tc.limit, tc.str1, tc.str2)

		// check the error wether it is expected or not
		if !errors.Is(err, tc.expectErr) {
			t.Errorf("test %d: expected error %v, but got: %v", i, tc.expectErr, err)
		}

		// check the result wether it is expected or not
		if !reflect.DeepEqual(result, tc.expectResult) {
			t.Errorf("test %d: expected result %v, but got: %v", i, tc.expectResult, result)
		}
	}
}

func TestFizzBuzz2(t *testing.T) {
	for i, tc := range getTestData() {
		// execute the fizzbuzz function
		result, err := business.FizzBuzz2(tc.ctx, tc.int1, tc.int2, tc.limit, tc.str1, tc.str2)

		// check the error wether it is expected or not
		if !errors.Is(err, tc.expectErr) {
			t.Errorf("test %d: expected error %v, but got: %v", i, tc.expectErr, err)
		}

		// check the result wether it is expected or not
		if !reflect.DeepEqual(result, tc.expectResult) {
			t.Errorf("test %d: expected result %v, but got: %v", i, tc.expectResult, result)
		}
	}
}

type fizzBuzzTestData struct {
	ctx               context.Context
	int1, int2, limit int
	str1, str2        string
	expectResult      []string
	expectErr         error
}

func getTestData() []fizzBuzzTestData {
	return []fizzBuzzTestData{
		{
			ctx:       context.TODO(),
			int1:      0,
			expectErr: business.ErrIntIsZero,
		},
		{
			ctx:       context.TODO(),
			int1:      1,
			int2:      0,
			expectErr: business.ErrIntIsZero,
		},
		{
			ctx:       context.TODO(),
			int1:      1,
			int2:      2,
			limit:     -1,
			expectErr: business.ErrLimitIsNegativeOrZero,
		},
		{
			ctx:       context.TODO(),
			int1:      1,
			int2:      2,
			limit:     0,
			expectErr: business.ErrLimitIsNegativeOrZero,
		},
		{
			ctx: func() context.Context {
				ctx, cancel := context.WithCancel(context.TODO())
				defer cancel()

				return ctx
			}(),
			int1:      1,
			int2:      2,
			limit:     3,
			expectErr: context.Canceled,
		},
		{
			ctx: func() context.Context {
				ctx, cancel := context.WithTimeout(context.TODO(), time.Duration(1)*time.Microsecond)
				time.Sleep(time.Duration(1) * time.Second)
				defer cancel()

				return ctx
			}(),
			int1:      1,
			int2:      2,
			limit:     3,
			expectErr: context.DeadlineExceeded,
		},
		{
			ctx:   context.TODO(),
			int1:  2,
			int2:  3,
			limit: 10,
			str1:  "foo",
			str2:  "bar",
			expectResult: []string{
				"1", "foo", "bar", "foo", "5", "foobar", "7", "foo", "bar", "foo",
			},
		},
		{
			ctx:   context.TODO(),
			int1:  1,
			int2:  1,
			limit: 10,
			str1:  "foo",
			str2:  "bar",
			expectResult: []string{
				"foobar", "foobar", "foobar", "foobar", "foobar", "foobar", "foobar", "foobar", "foobar", "foobar",
			},
		},
		{
			ctx:   context.TODO(),
			int1:  1,
			int2:  1,
			limit: 10,
			str1:  "",
			str2:  "",
			expectResult: []string{
				"", "", "", "", "", "", "", "", "", "",
			},
		},
		{
			ctx:   context.TODO(),
			int1:  11,
			int2:  12,
			limit: 10,
			str1:  "",
			str2:  "",
			expectResult: []string{
				"1", "2", "3", "4", "5", "6", "7", "8", "9", "10",
			},
		},
		{
			ctx:   context.TODO(),
			int1:  -2,
			int2:  -3,
			limit: 10,
			str1:  "foo",
			str2:  "bar",
			expectResult: []string{
				"1", "foo", "bar", "foo", "5", "foobar", "7", "foo", "bar", "foo",
			},
		},
	}
}

// Benchmark test function for FizzBuzz
func BenchmarkFizzBuzz(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range getTestData() {
			// execute the fizzbuzz function
			business.FizzBuzz(tc.ctx, tc.int1, tc.int2, tc.limit, tc.str1, tc.str2)
		}
	}
}

// Benchmark test function for FizzBuzz2
func BenchmarkFizzBuzz2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range getTestData() {
			// execute the fizzbuzz function
			business.FizzBuzz2(tc.ctx, tc.int1, tc.int2, tc.limit, tc.str1, tc.str2)
		}
	}
}
