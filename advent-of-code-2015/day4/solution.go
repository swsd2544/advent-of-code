package day4

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
)

func SolveLowestPositiveNumber(puzzle, prefix string) uint64 {
	var number uint64
	for {
		hash := md5.Sum([]byte(puzzle + strconv.FormatUint(number, 10)))
		if strings.HasPrefix(hex.EncodeToString(hash[:]), prefix) {
			return number
		}
		if number == 609043 {
			fmt.Println(hex.EncodeToString(hash[:]))
		}
		number += 1
	}
}
