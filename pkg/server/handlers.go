package server

import (
	"context"
	"errors"
	"fizz-buzz-gin/pkg/business"
	"fizz-buzz-gin/pkg/storage"
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

// fizzBuzzHandler launches the game with the required parameters
// @Summary Launches the game with the required parameters
// @Description The original fizz-buzz consists in writing all numbers from `1` to `100`, and replacing all multiples of `3` by `fizz`, all multiples of `5` by `buzz`, and all multiples of `15` by `fizzbuzz`. The output would look like this: 1,2,fizz,4,buzz,fizz,7,8,fizz,buzz,11,fizz,13,14,fizzbuzz,fizz,...
// @ID launch-fizz-buzz
// @Tags fizz-buzz
// @Produce  json
// @Param string1 query string false "First string related to multiples of int1"
// @Param string2 query string false "Second string related to multiples of int2"
// @Param int1 query int true "multiple related to string1"
// @Param int2 query int true "multiple related to string2"
// @Param limit query int true "number of elements to return"
// @Success 200 {object} []string
// @Failure 400 {object} server.HTTPError
// @Failure 503 {object} server.HTTPError
// @Router /fizz-buzz [get]
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

		// Save the parameters and result in storage
		err = storage.SaveLastCall(str1, str2, int1, int2, limit, result)
		if err != nil {
			c.JSON(http.StatusInternalServerError, "Error saving the last call to the fizz-buzz endpoint")
			return
		}

		// Return the result
		c.JSON(200, result)
	}
}

// statsHandler returns the parameters and result of the last calls to the fizz-buzz endpoint
// @Summary Returns the parameters and result of the last calls to the fizz-buzz endpoint
// @Description Everytime the fizz-buzz endpoint is called, the parameters related to the call are saved in a database, along with the time of the call, and the result given to the user. Calling this endpoint will return all statistics about the calls made so far.
// @ID get-fizz-buzz-stats
// @Tags fizz-buzz
// @Produce  json
// @Success 200 {object} storage.StatsCounterDto
// @Failure 503 {object} server.HTTPError
// @Router /stats [get]
func statisticsHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Retrieve the parameters and stats of the last call to the fizz-buzz endpoint
		stats, err := storage.GetLastCalls()
		if err != nil {
			c.JSON(http.StatusServiceUnavailable, NewHTTPError(http.StatusServiceUnavailable, err.Error()))
			return
		}

		// Return the result
		c.JSON(200, stats)
	}
}
