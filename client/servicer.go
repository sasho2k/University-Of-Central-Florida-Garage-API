package client

import (
	"fmt"
	"strings"
)

// ReturnParam will check the garage param to see if it is valid. -1 if not.
func ReturnParam(param string) (garageIdentifier int, err error) {
	param = strings.ToLower(param)
	switch param {
	case "a":
		return 0, nil
	case "b":
		return 1, nil
	case "c":
		return 2, nil
	case "d":
		return 3, nil
	case "h":
		return 4, nil
	case "i":
		return 5, nil
	case "libra":
		return 6, nil
	}
	return -1, fmt.Errorf("ERROR -> INVALID GARAGE. PLEASE INPUT A, B, C, D, H, I, LIBRA, OR ALL")
}
