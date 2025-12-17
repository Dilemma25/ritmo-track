package parser

import (
	"fmt"
	"strconv"
)

func ParseInt(str string) (int64, error) {
	integer, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("invalid string: %s", str)
	}
	return integer, nil
}
