package fizzbuzz

import (
	"context"
	"errors"
	"fmt"
	"math"
)

var (
	// ErrIntIsZero holds error for int1 and/or int2 being equals to zero.
	ErrIntIsZero = errors.New("zero int1 and/or int2")
	// ErrLimitIsNegativeOrZero holds error for limit being negative or zero.
	ErrLimitIsNegativeOrZero = errors.New("limit is negative or equals zero")
)

func FizzBuzz(ctx context.Context, int1, int2, limit int, str1, str2 string) ([]string, error) {
	// Let's avoid division by zero.
	if int1 == 0 || int2 == 0 {
		return nil, ErrIntIsZero
	}

	// Our implementation goes from 1 to limit.
	if limit < 1 {
		return nil, ErrLimitIsNegativeOrZero
	}

	// An alternative would be to reallocate on the go this slice.
	result := make([]string, limit)

	for i := 0; i < limit; i++ {
		// As the limit could be huge, we check if the provided context has
		// been canceled or if its deadline has exceeded.
		if ctx.Err() != nil {
			return nil, ctx.Err()
		}

		currentNumber := i + 1
		isInt1Multiple := math.Mod(float64(currentNumber)/float64(int1), 1.0) == 0
		isInt2Multiple := math.Mod(float64(currentNumber)/float64(int2), 1.0) == 0

		if isInt1Multiple && isInt2Multiple {
			result[i] = str1 + str2
		} else if isInt1Multiple {
			result[i] = str1
		} else if isInt2Multiple {
			result[i] = str2
		} else {
			result[i] = fmt.Sprintf("%d", currentNumber)
		}
	}

	return result, nil
}
