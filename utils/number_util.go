package utils

import (
	"crypto/rand"
	"errors"
	"fmt"
	"io"
)

var table = [...]byte{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0'}

func RandInt(max int) string {
	b := make([]byte, max)
	n, err := io.ReadAtLeast(rand.Reader, b, max)
	if n != max {
		panic(err)
	}
	for i := 0; i < len(b); i++ {
		b[i] = table[int(b[i])%len(table)]
	}
	return string(b)
}

func GetIndexSliceInt(value int, slice [6]int) (int, error) {
	for i, v := range slice {
		if value == v {
			return i, nil
		}
	}
	return -1, errors.New("value not exist")
}

func FloatToString64(value float64) string {
	return fmt.Sprintf("%f", value)
}
