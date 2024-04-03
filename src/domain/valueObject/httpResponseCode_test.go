package valueObject

import (
	"encoding/json"
	"strconv"
	"strings"
	"testing"
)

func TestNewHttpResponseCode(t *testing.T) {
	t.Run("ValidResponseCode", func(t *testing.T) {
		validResponseCodes := []interface{}{
			"100",
			"200",
			"300",
			"400",
			"500",
			100,
			200,
			300,
			400,
			500,
		}

		for _, responseCode := range validResponseCodes {
			_, err := NewHttpResponseCode(responseCode)
			if err != nil {
				t.Errorf(
					"Expected no error for %v, got %s",
					responseCode,
					err.Error(),
				)
			}
		}
	})

	t.Run("InvalidResponseCode", func(t *testing.T) {
		invalidResponseCodes := []interface{}{
			"@blabla",
			"<script>alert('xss')</script>",
			"1000",
			"0",
			"-1",
			"UNION SELECT * FROM USERS",
		}

		for _, responseCode := range invalidResponseCodes {
			_, err := NewHttpResponseCode(responseCode)
			if err == nil {
				t.Errorf("Expected error for %s, got nil", responseCode)
			}
		}
	})

	t.Run("ValidUnmarshalJSON", func(t *testing.T) {
		var testStruct struct {
			DataToTest HttpResponseCode
		}

		dataToTest := 200
		mapToTest := map[string]int{
			"dataToTest": dataToTest,
		}
		mapBytesToTest, _ := json.Marshal(mapToTest)

		reader := strings.NewReader(string(mapBytesToTest))
		jsonDecoder := json.NewDecoder(reader)
		err := jsonDecoder.Decode(&testStruct)
		if err != nil {
			t.Fatalf("Expected no error on UnmarshalJSON valid test, got %s", err.Error())
		}

		dataToTestFromStructStr := testStruct.DataToTest.String()
		dataToTestStr := strconv.Itoa(dataToTest)
		if dataToTestFromStructStr != dataToTestStr {
			t.Errorf(
				"VO data '%s' after UnmarshalJSON is not the same as the original data '%s'",
				dataToTestFromStructStr,
				dataToTestStr,
			)
		}
	})

	t.Run("InvalidUnmarshalJSON", func(t *testing.T) {
		var testStruct struct {
			DataToTest HttpResponseCode
		}

		dataToTest := ""
		mapToTest := map[string]string{
			"dataToTest": dataToTest,
		}
		mapBytesToTest, _ := json.Marshal(mapToTest)

		reader := strings.NewReader(string(mapBytesToTest))
		jsonDecoder := json.NewDecoder(reader)
		err := jsonDecoder.Decode(&testStruct)
		if err == nil {
			t.Fatal("Expected error on UnmarshalJSON invalid test, got nil")
		}
	})
}
