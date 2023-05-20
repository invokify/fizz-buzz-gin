package server_test

import (
	"context"
	"fizz-buzz-gin/pkg/business"
	"fizz-buzz-gin/pkg/server"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"strings"
	"testing"
	"time"
)

func TestPingHandler(t *testing.T) {
	e := server.NewServer(time.Duration(1) * time.Second)
	req := httptest.NewRequest(http.MethodGet, "/ping", nil)
	rec := httptest.NewRecorder()

	e.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("expected status code %d but got %d", http.StatusOK, rec.Code)
	}
}

func TestFizzBuzzHandler(t *testing.T) {
	for i, tc := range getTestData() {
		q := make(url.Values)
		q.Add("int1", tc.int1)
		q.Add("int2", tc.int2)
		q.Add("limit", tc.limit)
		q.Add("string1", tc.string1)
		q.Add("string2", tc.string2)

		e := server.NewServer(tc.timeout)
		req := httptest.NewRequest(http.MethodGet, "/api/v1/fizz-buzz?"+q.Encode(), nil)
		rec := httptest.NewRecorder()

		e.ServeHTTP(rec, req)

		if tc.expectStatusCode != rec.Code {
			t.Errorf("test %d: expected status code %d but got %d", i, tc.expectStatusCode, rec.Code)
		}

		actualJSON := strings.TrimSpace(rec.Body.String())
		if !reflect.DeepEqual(tc.expectJSON, actualJSON) {
			t.Errorf("test %d: expected response '%s', but got: '%s'", i, tc.expectJSON, actualJSON)
		}
	}
}

type fizzBuzzHandlerTestData struct {
	int1, int2, limit, string1, string2 string
	timeout                             time.Duration
	expectStatusCode                    int
	expectJSON                          string
}

func getTestData() []fizzBuzzHandlerTestData {
	return []fizzBuzzHandlerTestData{
		{
			int1:             "",
			expectStatusCode: http.StatusBadRequest,
			expectJSON:       `{"message":"int1 required field value is empty"}`,
		},
		{
			int1:             "foo",
			expectStatusCode: http.StatusBadRequest,
			expectJSON:       `{"message":"int1 failed to bind field value to int"}`,
		},
		{
			int1:             "1000000000000000000000000000000000000000000000000",
			expectStatusCode: http.StatusBadRequest,
			expectJSON:       `{"message":"int1 failed to bind field value to int"}`,
		},
		{
			int1:             "1",
			int2:             "",
			expectStatusCode: http.StatusBadRequest,
			expectJSON:       `{"message":"int2 required field value is empty"}`,
		},
		{
			int1:             "1",
			int2:             "foo",
			expectStatusCode: http.StatusBadRequest,
			expectJSON:       `{"message":"int2 failed to bind field value to int"}`,
		},
		{
			int1:             "1",
			int2:             "1000000000000000000000000000000000000000000000000",
			expectStatusCode: http.StatusBadRequest,
			expectJSON:       `{"message":"int2 failed to bind field value to int"}`,
		},
		{
			int1:             "1",
			int2:             "1",
			limit:            "",
			expectStatusCode: http.StatusBadRequest,
			expectJSON:       `{"message":"limit required field value is empty"}`,
		},
		{
			int1:             "1",
			int2:             "1",
			limit:            "foo",
			expectStatusCode: http.StatusBadRequest,
			expectJSON:       `{"message":"limit failed to bind field value to int"}`,
		},
		{
			int1:             "1",
			int2:             "1",
			limit:            "1000000000000000000000000000000000000000000000000",
			expectStatusCode: http.StatusBadRequest,
			expectJSON:       `{"message":"limit failed to bind field value to int"}`,
		},
		{
			int1:             "0",
			int2:             "1",
			limit:            "1",
			expectStatusCode: http.StatusBadRequest,
			expectJSON: func() string {
				return fmt.Sprintf(`{"message":"%s"}`, business.ErrIntIsZero)
			}(),
		},
		{
			int1:             "1",
			int2:             "0",
			limit:            "1",
			expectStatusCode: http.StatusBadRequest,
			expectJSON: func() string {
				return fmt.Sprintf(`{"message":"%s"}`, business.ErrIntIsZero)
			}(),
		},
		{
			int1:             "1",
			int2:             "1",
			limit:            "-1",
			expectStatusCode: http.StatusBadRequest,
			expectJSON: func() string {
				return fmt.Sprintf(`{"message":"%s"}`, business.ErrLimitIsNegativeOrZero)
			}(),
		},
		{
			int1:             "1",
			int2:             "1",
			limit:            "0",
			expectStatusCode: http.StatusBadRequest,
			expectJSON: func() string {
				return fmt.Sprintf(`{"message":"%s"}`, business.ErrLimitIsNegativeOrZero)
			}(),
		},
		{
			int1:             "2",
			int2:             "3",
			limit:            "100000",
			expectStatusCode: http.StatusServiceUnavailable,
			timeout:          time.Duration(1) * time.Nanosecond,
			expectJSON: func() string {
				return fmt.Sprintf(`{"message":"%s"}`, context.DeadlineExceeded)
			}(),
		},
		{
			int1:             "2",
			int2:             "3",
			limit:            "10",
			string1:          "foo",
			string2:          "bar",
			timeout:          time.Duration(10) * time.Second,
			expectStatusCode: http.StatusOK,
			expectJSON:       `["1","foo","bar","foo","5","foobar","7","foo","bar","foo"]`,
		},
	}
}
