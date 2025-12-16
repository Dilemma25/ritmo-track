package parser

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func ParseInt(str string) (int64, error) {
	integer, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("invalid string: %s", str)
	}
	return integer, nil
}

func ParseRequestBody(r *http.Request, reqType interface{}) error {
	if err := json.NewDecoder(r.Body).Decode(&reqType); err != nil {
		return fmt.Errorf("invalid json")
	}

	return nil
}
