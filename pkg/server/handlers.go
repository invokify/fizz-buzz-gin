package server

import (
	"context"
	"errors"
	"fizz-buzz-gin/pkg/business"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// constants
const (
	ErrMsgNotProvided = "%s required field value is empty"
	ErrMsgNotInteger  = "%s failed to bind field value to int"
)

func fizzBuzzHandler(timeout time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Retrieve the query parameters if exist otherwise use default values
		str1 := c.DefaultQuery("string1", "Fizz")
		str2 := c.DefaultQuery("string2", "Buzz")

		// checks wether int1 exists or is empty
		int1str, ok := c.GetQuery("int1")
		if !ok || int1str == "" {
			c.JSON(http.StatusBadRequest, NewHTTPError(http.StatusBadRequest, fmt.Sprintf(ErrMsgNotProvided, "int1")))
			return
		}

		// convert int1 to integer
		int1, err := strconv.Atoi(int1str)
		if err != nil {
			c.JSON(http.StatusBadRequest, NewHTTPError(http.StatusBadRequest, fmt.Sprintf(ErrMsgNotInteger, "int1")))
			return
		}

		// checks wether int2 exists or is empty
		int2str, ok := c.GetQuery("int2")
		if !ok || int2str == "" {
			c.JSON(http.StatusBadRequest, NewHTTPError(http.StatusBadRequest, fmt.Sprintf(ErrMsgNotProvided, "int2")))
			return
		}

		// convert int2 to integer
		int2, err := strconv.Atoi(int2str)
		if err != nil {
			c.JSON(http.StatusBadRequest, NewHTTPError(http.StatusBadRequest, fmt.Sprintf(ErrMsgNotInteger, "int2")))
			return
		}

		// checks wether limit exists or is empty
		limitstr, ok := c.GetQuery("limit")
		if !ok || limitstr == "" {
			c.JSON(http.StatusBadRequest, NewHTTPError(http.StatusBadRequest, fmt.Sprintf(ErrMsgNotProvided, "limit")))
			return
		}

		// convert limit to integer
		limit, err := strconv.Atoi(limitstr)
		if err != nil {
			c.JSON(http.StatusBadRequest, NewHTTPError(http.StatusBadRequest, fmt.Sprintf(ErrMsgNotInteger, "limit")))
			return
		}

		// Create a context with a timeout
		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()

		// Call the fizzbuzz function
		result, err := business.FizzBuzz(ctx, int1, int2, limit, str1, str2)
		if err != nil {
			if errors.Is(err, business.ErrIntIsZero) || errors.Is(err, business.ErrLimitIsNegativeOrZero) {
				c.JSON(http.StatusBadRequest, NewHTTPError(http.StatusBadRequest, err.Error()))
			}

			if errors.Is(err, ctx.Err()) {
				c.JSON(http.StatusServiceUnavailable, NewHTTPError(http.StatusServiceUnavailable, err.Error()))
			}

			return
		}

		// Return the result
		c.JSON(200, result)
	}
}
