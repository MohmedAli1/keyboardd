// this is a description for my package
package keyboardd

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

// this is a description for my function "getfloat"
func GetFloat() (float64, error) {

	m := bufio.NewReader(os.Stdin)
	y, err := m.ReadString('\n')
	if err != nil {
		return 0, err
	}
	number, err := strconv.ParseFloat(strings.TrimSpace(y), 64)
	if err != nil {
		return 0, err
	}
	return number, nil

}
