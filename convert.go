package pot

import (
	"encoding/json"
	"strconv"
)

/*
valConvertString
@Desc: 数据转换
@param: value
@return: string
*/
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
	var (
		result uint64
	)
	if value == nil {
		return 0
	}

	switch value.(type) {
	case float64:
		result = uint64(value.(float64))
	case float32:
		result = uint64(value.(float32))
	case int:
		result = uint64(value.(int))
	case uint:
		result = uint64(value.(uint))
	case int8:
		result = uint64(value.(int8))
	case uint8:
		result = uint64(value.(uint8))
	case int16:
		result = uint64(value.(int16))
	case uint16:
		result = uint64(value.(int16))
	case int32:
		result = uint64(value.(int32))
	case uint32:
		result = uint64(value.(uint32))
	case int64:
		result = uint64(value.(int64))
	case uint64:
		result = value.(uint64)
	}
	return result
}

/*
valConvertInt64
@Desc: parse to int64
@param: value
@return: uint64
*/
func valConvertInt64(value interface{}) int64 {
	var (
		result int64
	)
	if value == nil {
		return 0
	}

	switch value.(type) {
	case float64:
		result = int64(value.(float64))
	case float32:
		result = int64(value.(float32))
	case int:
		result = int64(value.(int))
	case uint:
		result = int64(value.(uint))
	case int8:
		result = int64(value.(int8))
	case uint8:
		result = int64(value.(uint8))
	case int16:
		result = int64(value.(int16))
	case uint16:
		result = int64(value.(int16))
	case int32:
		result = int64(value.(int32))
	case uint32:
		result = int64(value.(uint32))
	case int64:
		result = value.(int64)
	case uint64:
		result = int64(value.(uint64))
	}
	return result
}

/*
valConvertFloat64
@Desc: parse to float64
@param: value
@return: uint64
*/
func valConvertFloat64(value interface{}) float64 {
	var (
		result float64
	)
	if value == nil {
		return 0
	}

	switch value.(type) {
	case float64:
		result = value.(float64)
	case float32:
		result = float64(value.(float32))
	case int:
		result = float64(value.(int))
	case uint:
		result = float64(value.(uint))
	case int8:
		result = float64(value.(int8))
	case uint8:
		result = float64(value.(uint8))
	case int16:
		result = float64(value.(int16))
	case uint16:
		result = float64(value.(int16))
	case int32:
		result = float64(value.(int32))
	case uint32:
		result = float64(value.(uint32))
	case int64:
		result = float64(value.(int64))
	case uint64:
		result = float64(value.(uint64))
	}
	return result
}
