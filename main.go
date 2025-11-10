package main

import (
	"errors"
	"fmt"
)

func main() {
	// bools := []bool{true, false, true, false, false, false, false, false}
	// packedBools, err := packBoolArray(bools)

	// if err != nil {
	// 	fmt.Println(err.Error())
	// } else {
	// 	fmt.Printf("Packed bools: %v \n\n", packedBools)
	// }
	byteTest(1)
	byteTest(7)
	byteTest(10)
}

func packBoolArray(unpackedBools []bool) (packedBools byte, err error) {
	if len(unpackedBools) > 8 {
		return 0, errors.New("There is a max size of 8 booleans")
	}
	fmt.Printf("Unpacked: %v \n\n", unpackedBools)

	return packedBools, nil
}

func byteTest(shiftLeft int) {
	fmt.Printf("Shift left %d times", shiftLeft)
	myWeeByte := byte(0)
	mask := byte(1 << shiftLeft)
	myWeeByte = myWeeByte | mask

	fmt.Printf("Decimal: %d\n", myWeeByte)
	fmt.Printf("Binary: %08b\n\n", myWeeByte)
}
