package data

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var ErrInvalidSportsEquipmentNumberFormat = errors.New("invalid sports equipment format")

type SportsEquipmentNumber int32

func (s *SportsEquipmentNumber) MarshalJSON() ([]byte, error) {
	jsonValue := fmt.Sprintf("%d mins", s)

	quotedJSONValue := strconv.Quote(jsonValue)

	return []byte(quotedJSONValue), nil
}

func (s *SportsEquipmentNumber) UnmarshalJSON(jsonValue []byte) error {
	unquotedJSONValue, err := strconv.Unquote(string(jsonValue))
	if err != nil {
		return ErrInvalidSportsEquipmentNumberFormat
	}

	parts := strings.Split(unquotedJSONValue, " ")
	if len(parts) != 2 || parts[1] != "mins" {
		return ErrInvalidSportsEquipmentNumberFormat
	}

	i, err := strconv.ParseInt(parts[0], 10, 32)
	if err != nil {
		return ErrInvalidSportsEquipmentNumberFormat
	}

	*s = SportsEquipmentNumber(i)
	return nil
}
