package gopot

import (
	"encoding/json"
	"strconv"
)

// valConvertString
func valConvertString(value interface{}) string {
	var (
		result string
	)
	if value == nil {
		return result
	}

	switch value.(type) {
	case float64:
		ft := value.(float64)
		result = strconv.FormatFloat(ft, 'f', -1, 64)
	case float32:
		ft := value.(float32)
		result = strconv.FormatFloat(float64(ft), 'f', -1, 64)
	case int:
		it := value.(int)
		result = strconv.Itoa(it)
	case uint:
		it := value.(uint)
		result = strconv.Itoa(int(it))
	case int8:
		it := value.(int8)
		result = strconv.Itoa(int(it))
	case uint8:
		it := value.(uint8)
		result = strconv.Itoa(int(it))
	case int16:
		it := value.(int16)
		result = strconv.Itoa(int(it))
	case uint16:
		it := value.(uint16)
		result = strconv.Itoa(int(it))
	case int32:
		it := value.(int32)
		result = strconv.Itoa(int(it))
	case uint32:
		it := value.(uint32)
		result = strconv.Itoa(int(it))
	case int64:
		it := value.(int64)
		result = strconv.FormatInt(it, 10)
	case uint64:
		it := value.(uint64)
		result = strconv.FormatUint(it, 10)
	case string:
		result = value.(string)
	case []byte:
		result = string(value.([]byte))
	default:
		newValue, _ := json.Marshal(value)
		result = string(newValue)
	}
	return result
}

/*
valConvertUint64
@Desc: parse to uint64
@param: value
@return: uint64
*/
func valConvertUint64(value interface{}) uint64 {
	if value == nil {
		return 0
	}

	switch v := value.(type) {
	case float64:
		return uint64(v)
	case float32:
		return uint64(v)
	case int:
		return uint64(v)
	case uint:
		return uint64(v)
	case int8:
		return uint64(v)
	case uint8:
		return uint64(v)
	case int16:
		return uint64(v)
	case uint16:
		return uint64(v)
	case int32:
		return uint64(v)
	case uint32:
		return uint64(v)
	case int64:
		if v < 0 {
			return 0
		}
		return uint64(v)
	case uint64:
		return v
	case string:
		if u, err := strconv.ParseUint(v, 10, 64); err == nil {
			return u
		}
	}
	return 0
}

/*
valConvertInt64
@Desc: parse to int64
@param: value
@return: uint64
*/
func valConvertInt64(value interface{}) int64 {
	if value == nil {
		return 0
	}

	switch v := value.(type) {
	case float64:
		return int64(v)
	case float32:
		return int64(v)
	case int:
		return int64(v)
	case uint:
		return int64(v)
	case int8:
		return int64(v)
	case uint8:
		return int64(v)
	case int16:
		return int64(v)
	case uint16:
		return int64(v)
	case int32:
		return int64(v)
	case uint32:
		return int64(v)
	case int64:
		return v
	case uint64:
		return int64(v)
	case string:
		val, _ := strconv.ParseInt(v, 10, 64)
		return val
	}
	return 0
}

/*
valConvertFloat64
@Desc: parse to float64
@param: value
@return: uint64
*/
func valConvertFloat64(value interface{}) float64 {
	if value == nil {
		return 0
	}

	switch v := value.(type) {
	case float64:
		return v
	case float32:
		return float64(v)
	case int:
		return float64(v)
	case uint:
		return float64(v)
	case int8:
		return float64(v)
	case uint8:
		return float64(v)
	case int16:
		return float64(v)
	case uint16:
		return float64(v)
	case int32:
		return float64(v)
	case uint32:
		return float64(v)
	case int64:
		return float64(v)
	case uint64:
		return float64(v)
	case string:
		if f, err := strconv.ParseFloat(v, 64); err == nil {
			return f
		}
	}
	return 0
}
