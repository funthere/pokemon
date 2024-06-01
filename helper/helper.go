package helper

import (
	"fmt"
	"strconv"
)

func StringToUint32(str string) (uint32Num uint32) {
	// Convert string to uint32
	num, err := strconv.ParseUint(str, 10, 32)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Type assert to uint32
	uint32Num = uint32(num)
	return
}
