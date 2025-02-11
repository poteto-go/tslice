package tslice

import (
	"fmt"
	"strconv"
)

func toString[V any](data V) string {
	switch asserted := any(data).(type) {
	case string:
		return asserted
	case []byte:
		return string(asserted)
	case []rune:
		return string(asserted)
	case int:
		return strconv.Itoa(asserted)
	case int8:
		return strconv.Itoa(int(asserted))
	case int16:
		return strconv.Itoa(int(asserted))
	case int32:
		return strconv.Itoa(int(asserted))
	case int64:
		return strconv.Itoa(int(asserted))
	case uint:
		return strconv.FormatUint(uint64(asserted), 10)
	case uint8:
		return strconv.FormatUint(uint64(asserted), 10)
	case uint16:
		return strconv.FormatUint(uint64(asserted), 10)
	case uint32:
		return strconv.FormatUint(uint64(asserted), 10)
	case uint64:
		return strconv.FormatUint(uint64(asserted), 10)
	case float32:
		return strconv.FormatFloat(float64(asserted), 'f', -1, 32)
	case float64:
		return strconv.FormatFloat(asserted, 'f', -1, 64)
	case bool:
		return strconv.FormatBool(asserted)
	default:
		return fmt.Sprintf("%v", data)
	}
}
