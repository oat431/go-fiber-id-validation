package service

import (
	"strconv"

	"github.com/oat431/go-fiber-id-validation/payload"
)

const (
	IDLengthWithoutChecksum = 12
	ChecksumBaseMultiplier  = 13
	ChecksumModulus         = 11
)

func lastDigit(generated string) int {
	sum := 0
	for i := range IDLengthWithoutChecksum {
		digit, err := strconv.Atoi(string(generated[i]))
		if err != nil {
			return -1
		}

		sum += digit * (ChecksumBaseMultiplier - i)
	}

	return (ChecksumModulus - (sum % ChecksumModulus)) % 10
}

func isNIDValid(id string) bool {
	if len(id) != IDLengthWithoutChecksum+1 {
		return false
	}

	lastChar := string(id[len(id)-1])
	lastDigitInt, err := strconv.Atoi(lastChar)

	if err != nil {
		return false
	}

	calculatedChecksum := lastDigit(id[:IDLengthWithoutChecksum])

	return calculatedChecksum == lastDigitInt
}

func ValidateNID(id string) payload.ValidateDto {
	return payload.ValidateDto{
		NID:     id,
		IsValid: isNIDValid(id),
	}
}
