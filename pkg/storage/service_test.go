package storage_test

import (
	"errors"
	"fizz-buzz-gin/pkg/storage"
	"log"
	"os"
	"reflect"
	"testing"
)

func TestMain(m *testing.M) {
	code := m.Run()

	err := os.Remove(storage.DBPath)
	if err != nil {
		log.Println(err)
	}

	os.Exit(code)
}

func TestInsert(t *testing.T) {
	for i, tc := range getTestData() {
		// insert data
		storage.SaveLastCall(tc.str1, tc.str2, tc.int1, tc.int2, tc.limit, tc.result)

		// get data
		result, errG := storage.GetLastCalls()

		// check the error wether it is expected or not
		if !errors.Is(errG, tc.expectErr) {
			t.Errorf("test %d: expected error %v, but got: %v", i, tc.expectErr, errG)
		}

		// check the result wether it is expected or not
		if !reflect.DeepEqual(result.Stats[i].Result, tc.expectResult) {
			t.Errorf("test %d: expected result %v, but got: %v", i, tc.expectResult, result.Stats[i].Result)
		}
	}
}

type fizzBuzzTestData struct {
	int1, int2, limit int
	str1, str2        string
	result            []string
	expectResult      string
	expectErr         error
}

func getTestData() []fizzBuzzTestData {
	return []fizzBuzzTestData{
		{
			int1:  2,
			int2:  3,
			limit: 10,
			str1:  "foo",
			str2:  "bar",
			result: []string{
				"1", "foo", "bar", "foo", "5", "foobar", "7", "foo", "bar", "foo",
			},
			expectResult: "[1,foo,bar,foo,5,foobar,7,foo,bar,foo]",
		},
	}
}
