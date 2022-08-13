// A utility package for getting various types of data from the keyboard (int, float, bool, string)
package keyboard

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader *bufio.Reader
var input string
var err error

func getReader() {
	reader = bufio.NewReader(os.Stdin)
}
func GetFloat() (float64, error) {
	if reader == nil {
		getReader()
	}

	fmt.Print("Enter a floating point number: ")
	input, err = reader.ReadString('\n')
	input = strings.TrimSpace(input)
	if err != nil {
		return 0, fmt.Errorf("Error getting input from user.")
	}

	float, err := strconv.ParseFloat(input, 64)
	if err != nil {
		return 0, fmt.Errorf("Input must be a number.")
	}

	return float, nil
}
