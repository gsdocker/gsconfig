package gsconfig

import (
	"errors"
	"strconv"
	"time"

	"github.com/gsdocker/gserrors"
)

//gsconfig package's public errors
var (
	ErrConvert = errors.New("config object convert error")
)

var config map[string]string

func init() {
	config = make(map[string]string)
}

//Save save as global config
func Save(kvs map[string]string) {
	for k, v := range kvs {
		config[k] = v
	}
}

// Set set new config item
func Set(k, v string) {
	config[k] = v
}

//Get .
func Get(key string) (val string, ok bool) {

	val, ok = config[key]

	return
}

//String .
func String(key string, defaultval string) string {
	val, ok := Get(key)

	if ok {

		return val
	}

	return defaultval
}

//Int64 .
func Int64(key string, defaultval int64) int64 {
	val, ok := Get(key)

	if ok {
		i, err := strconv.ParseInt(val, 0, 64)

		if err != nil {
			gserrors.Panic(ErrConvert)
		}

		return i
	}

	return defaultval
}

//Int32 .
func Int32(key string, defaultval int32) int32 {
	return int32(Int64(key, int64(defaultval)))
}

//Int16 .
func Int16(key string, defaultval int16) int16 {
	return int16(Int64(key, int64(defaultval)))
}

//Int .
func Int(key string, defaultval int) int {
	return int(Int64(key, int64(defaultval)))
}

//Uint64 .
func Uint64(key string, defaultval uint64) uint64 {
	val, ok := Get(key)

	if ok {
		i, err := strconv.ParseUint(val, 0, 64)

		if err != nil {
			gserrors.Panic(ErrConvert)
		}

		return i
	}

	return defaultval
}

//Uint32 .
func Uint32(key string, defaultval uint32) uint32 {
	return uint32(Uint64(key, uint64(defaultval)))
}

//Uint16 .
func Uint16(key string, defaultval uint16) uint16 {
	return uint16(Uint64(key, uint64(defaultval)))
}

//Uint .
func Uint(key string, defaultval uint) uint {
	return uint(Uint64(key, uint64(defaultval)))
}

//Bool .
func Bool(key string, defaultval bool) bool {
	val, ok := Get(key)

	if ok {
		i, err := strconv.ParseBool(val)

		if err != nil {
			gserrors.Panic(ErrConvert)
		}

		return i
	}

	return defaultval
}

//Float64 .
func Float64(key string, defaultval float64) float64 {
	val, ok := Get(key)

	if ok {
		i, err := strconv.ParseFloat(val, 64)

		if err != nil {
			gserrors.Panic(ErrConvert)
		}

		return i
	}

	return defaultval
}

//Float32 .
func Float32(key string, defaultval float32) float32 {
	return float32(Float64(key, float64(defaultval)))
}

//Seconds .
func Seconds(key string, defaultval int64) time.Duration {
	return time.Duration(Int64(key, defaultval)) * time.Second
}

//Milliseconds .
func Milliseconds(key string, defaultval int64) time.Duration {
	return time.Duration(Int64(key, defaultval)) * time.Millisecond
}
