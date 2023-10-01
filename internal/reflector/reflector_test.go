package reflector_test

import (
	"errors"
	"reflect"
	"testing"

	"github.com/livebud/bud/internal/reflector"
	"github.com/matryer/is"
)

type customStruct struct {
	Field1 int
	Field2 string
}

type customInterface interface {
	Method1() int
}

type customFunc func() int

type customAlias = int

func TestName(t *testing.T) {
	is := is.New(t)
	name, err := reflector.TypeOf(customStruct{})
	is.NoErr(err)
	is.Equal(name, "github.com/livebud/bud/internal/reflector_test.customStruct")
	name, err = reflector.TypeOf(&customStruct{})
	is.NoErr(err)
	is.Equal(name, "github.com/livebud/bud/internal/reflector_test.*customStruct")
	name, err = reflector.TypeOf(reflect.ValueOf((*customInterface)(nil)).Elem())
	is.NoErr(err)
	is.Equal(name, "reflect.Value")
	name, err = reflector.TypeOf(customFunc(func() int { return 0 }))
	is.NoErr(err)
	is.Equal(name, "github.com/livebud/bud/internal/reflector_test.customFunc")
}

func TestNoName(t *testing.T) {
	is := is.New(t)
	name, err := reflector.TypeOf(bool(false))
	is.True(err != nil)
	is.True(errors.Is(err, reflector.ErrNoName))
	is.Equal(name, "")
	name, err = reflector.TypeOf(int(0))
	is.True(err != nil)
	is.True(errors.Is(err, reflector.ErrNoName))
	is.Equal(name, "")
	name, err = reflector.TypeOf(int8(0))
	is.True(err != nil)
	is.True(errors.Is(err, reflector.ErrNoName))
	is.Equal(name, "")
	name, err = reflector.TypeOf(int16(0))
	is.True(err != nil)
	is.True(errors.Is(err, reflector.ErrNoName))
	is.Equal(name, "")
	name, err = reflector.TypeOf(int32(0))
	is.True(err != nil)
	is.True(errors.Is(err, reflector.ErrNoName))
	is.Equal(name, "")
	name, err = reflector.TypeOf(int64(0))
	is.True(err != nil)
	is.True(errors.Is(err, reflector.ErrNoName))
	is.Equal(name, "")
	name, err = reflector.TypeOf(uint(0))
	is.True(err != nil)
	is.True(errors.Is(err, reflector.ErrNoName))
	is.Equal(name, "")
	name, err = reflector.TypeOf(uint8(0))
	is.True(err != nil)
	is.True(errors.Is(err, reflector.ErrNoName))
	is.Equal(name, "")
	name, err = reflector.TypeOf(uint16(0))
	is.True(err != nil)
	is.True(errors.Is(err, reflector.ErrNoName))
	is.Equal(name, "")
	name, err = reflector.TypeOf(uint32(0))
	is.True(err != nil)
	is.True(errors.Is(err, reflector.ErrNoName))
	is.Equal(name, "")
	name, err = reflector.TypeOf(uint64(0))
	is.True(err != nil)
	is.True(errors.Is(err, reflector.ErrNoName))
	is.Equal(name, "")
	name, err = reflector.TypeOf(uintptr(0))
	is.True(err != nil)
	is.True(errors.Is(err, reflector.ErrNoName))
	is.Equal(name, "")
	name, err = reflector.TypeOf(float32(0))
	is.True(err != nil)
	is.True(errors.Is(err, reflector.ErrNoName))
	is.Equal(name, "")
	name, err = reflector.TypeOf(float64(0))
	is.True(err != nil)
	is.True(errors.Is(err, reflector.ErrNoName))
	is.Equal(name, "")
	name, err = reflector.TypeOf(complex64(0))
	is.True(err != nil)
	is.True(errors.Is(err, reflector.ErrNoName))
	is.Equal(name, "")
	name, err = reflector.TypeOf(complex128(0))
	is.True(err != nil)
	is.True(errors.Is(err, reflector.ErrNoName))
	is.Equal(name, "")
	name, err = reflector.TypeOf(string(""))
	is.True(err != nil)
	is.True(errors.Is(err, reflector.ErrNoName))
	is.Equal(name, "")
	name, err = reflector.TypeOf(byte(0))
	is.True(err != nil)
	is.True(errors.Is(err, reflector.ErrNoName))
	is.Equal(name, "")
	name, err = reflector.TypeOf(rune(0))
	is.True(err != nil)
	is.True(errors.Is(err, reflector.ErrNoName))
	is.Equal(name, "")
	name, err = reflector.TypeOf(errors.New(""))
	is.True(err == nil)
	is.Equal(name, "errors.*errorString")
	name, err = reflector.TypeOf(new(int))
	is.True(err != nil)
	is.True(errors.Is(err, reflector.ErrNoName))
	is.Equal(name, "")
	name, err = reflector.TypeOf(new(int8))
	is.True(err != nil)
	is.True(errors.Is(err, reflector.ErrNoName))
	is.Equal(name, "")
	name, err = reflector.TypeOf(new(int16))
	is.True(err != nil)
	is.True(errors.Is(err, reflector.ErrNoName))
	is.Equal(name, "")
	name, err = reflector.TypeOf(new(int32))
	is.True(err != nil)
	is.True(errors.Is(err, reflector.ErrNoName))
	is.Equal(name, "")
	name, err = reflector.TypeOf(new(int64))
	is.True(err != nil)
	is.True(errors.Is(err, reflector.ErrNoName))
	is.Equal(name, "")
	name, err = reflector.TypeOf(new(uint))
	is.True(err != nil)
	is.True(errors.Is(err, reflector.ErrNoName))
	is.Equal(name, "")
	name, err = reflector.TypeOf(new(uint8))
	is.True(err != nil)
	is.True(errors.Is(err, reflector.ErrNoName))
	is.Equal(name, "")
	name, err = reflector.TypeOf(new(uint16))
	is.True(err != nil)
	is.True(errors.Is(err, reflector.ErrNoName))
	is.Equal(name, "")
	name, err = reflector.TypeOf(new(uint32))
	is.True(err != nil)
	is.True(errors.Is(err, reflector.ErrNoName))
	is.Equal(name, "")
	name, err = reflector.TypeOf(new(uint64))
	is.True(err != nil)
	is.True(errors.Is(err, reflector.ErrNoName))
	is.Equal(name, "")
	name, err = reflector.TypeOf(new(uintptr))
	is.True(err != nil)
	is.True(errors.Is(err, reflector.ErrNoName))
	is.Equal(name, "")
	name, err = reflector.TypeOf(new(float32))
	is.True(err != nil)
	is.True(errors.Is(err, reflector.ErrNoName))
	is.Equal(name, "")
	name, err = reflector.TypeOf(new(float64))
	is.True(err != nil)
	is.True(errors.Is(err, reflector.ErrNoName))
	is.Equal(name, "")
	name, err = reflector.TypeOf(new(complex64))
	is.True(err != nil)
	is.True(errors.Is(err, reflector.ErrNoName))
	is.Equal(name, "")
	name, err = reflector.TypeOf(new(complex128))
	is.True(err != nil)
	is.True(errors.Is(err, reflector.ErrNoName))
	is.Equal(name, "")
	name, err = reflector.TypeOf(new(bool))
	is.True(err != nil)
	is.True(errors.Is(err, reflector.ErrNoName))
	is.Equal(name, "")
	name, err = reflector.TypeOf(new(string))
	is.True(err != nil)
	is.True(errors.Is(err, reflector.ErrNoName))
	is.Equal(name, "")
	name, err = reflector.TypeOf(new(byte))
	is.True(err != nil)
	is.True(errors.Is(err, reflector.ErrNoName))
	is.Equal(name, "")
	name, err = reflector.TypeOf(new(rune))
	is.True(err != nil)
	is.True(errors.Is(err, reflector.ErrNoName))
	is.Equal(name, "")
	name, err = reflector.TypeOf(new(error))
	is.True(err != nil)
	is.True(errors.Is(err, reflector.ErrNoName))
	is.Equal(name, "")
	name, err = reflector.TypeOf(new(interface{}))
	is.True(err != nil)
	is.True(errors.Is(err, reflector.ErrNoName))
	is.Equal(name, "")
	name, err = reflector.TypeOf([]int{})
	is.True(err != nil)
	is.True(errors.Is(err, reflector.ErrNoName))
	is.Equal(name, "")
	name, err = reflector.TypeOf([]int8{})
	is.True(err != nil)
	is.True(errors.Is(err, reflector.ErrNoName))
	is.Equal(name, "")
	name, err = reflector.TypeOf([]int16{})
	is.True(err != nil)
	is.True(errors.Is(err, reflector.ErrNoName))
	is.Equal(name, "")
	name, err = reflector.TypeOf([]int32{})
	is.True(err != nil)
	is.True(errors.Is(err, reflector.ErrNoName))
	is.Equal(name, "")
	name, err = reflector.TypeOf([]int64{})
	is.True(err != nil)
	is.True(errors.Is(err, reflector.ErrNoName))
	is.Equal(name, "")
	name, err = reflector.TypeOf([]uint{})
	is.True(err != nil)
	is.True(errors.Is(err, reflector.ErrNoName))
	is.Equal(name, "")
	name, err = reflector.TypeOf([]uint8{})
	is.True(err != nil)
	is.True(errors.Is(err, reflector.ErrNoName))
	is.Equal(name, "")
	name, err = reflector.TypeOf([]uint16{})
	is.True(err != nil)
	is.True(errors.Is(err, reflector.ErrNoName))
	is.Equal(name, "")
	name, err = reflector.TypeOf([]uint32{})
	is.True(err != nil)
	is.True(errors.Is(err, reflector.ErrNoName))
	is.Equal(name, "")
	name, err = reflector.TypeOf([]uint64{})
	is.True(err != nil)
	is.True(errors.Is(err, reflector.ErrNoName))
	is.Equal(name, "")
	name, err = reflector.TypeOf([]uintptr{})
	is.True(err != nil)
	is.True(errors.Is(err, reflector.ErrNoName))
	is.Equal(name, "")
	name, err = reflector.TypeOf([]float32{})
	is.True(err != nil)
	is.True(errors.Is(err, reflector.ErrNoName))
	is.Equal(name, "")
	name, err = reflector.TypeOf([]float64{})
	is.True(err != nil)
	is.True(errors.Is(err, reflector.ErrNoName))
	is.Equal(name, "")
	name, err = reflector.TypeOf([]complex64{})
	is.True(err != nil)
	is.True(errors.Is(err, reflector.ErrNoName))
	is.Equal(name, "")
	name, err = reflector.TypeOf([]bool{})
	is.True(err != nil)
	is.True(errors.Is(err, reflector.ErrNoName))
	is.Equal(name, "")
	name, err = reflector.TypeOf([]string{})
	is.True(err != nil)
	is.True(errors.Is(err, reflector.ErrNoName))
	is.Equal(name, "")
	name, err = reflector.TypeOf([]byte{})
	is.True(err != nil)
	is.True(errors.Is(err, reflector.ErrNoName))
	is.Equal(name, "")
	name, err = reflector.TypeOf([]rune{})
	is.True(err != nil)
	is.True(errors.Is(err, reflector.ErrNoName))
	is.Equal(name, "")
	name, err = reflector.TypeOf([]error{})
	is.True(err != nil)
	is.True(errors.Is(err, reflector.ErrNoName))
	is.Equal(name, "")
	name, err = reflector.TypeOf([]interface{}{})
	is.True(err != nil)
	is.True(errors.Is(err, reflector.ErrNoName))
	is.Equal(name, "")
	name, err = reflector.TypeOf([1]int{})
	is.True(err != nil)
	is.True(errors.Is(err, reflector.ErrNoName))
	is.Equal(name, "")
	name, err = reflector.TypeOf([1]int8{})
	is.True(err != nil)
	is.True(errors.Is(err, reflector.ErrNoName))
	is.Equal(name, "")
	name, err = reflector.TypeOf([1]int16{})
	is.True(err != nil)
	is.True(errors.Is(err, reflector.ErrNoName))
	is.Equal(name, "")
	name, err = reflector.TypeOf([1]int32{})
	is.True(err != nil)
	is.True(errors.Is(err, reflector.ErrNoName))
	is.Equal(name, "")
	name, err = reflector.TypeOf([1]int64{})
	is.True(err != nil)
	is.True(errors.Is(err, reflector.ErrNoName))
	is.Equal(name, "")
	name, err = reflector.TypeOf([1]uint{})
	is.True(err != nil)
	is.True(errors.Is(err, reflector.ErrNoName))
	is.Equal(name, "")
	name, err = reflector.TypeOf([1]uint8{})
	is.True(err != nil)
	is.True(errors.Is(err, reflector.ErrNoName))
	is.Equal(name, "")
	name, err = reflector.TypeOf([1]uint16{})
	is.True(err != nil)
	is.True(errors.Is(err, reflector.ErrNoName))
	is.Equal(name, "")
	name, err = reflector.TypeOf([1]uint32{})
	is.True(err != nil)
	is.True(errors.Is(err, reflector.ErrNoName))
	is.Equal(name, "")
	name, err = reflector.TypeOf([1]uint64{})
	is.True(err != nil)
	is.True(errors.Is(err, reflector.ErrNoName))
	is.Equal(name, "")
	name, err = reflector.TypeOf([1]uintptr{})
	is.True(err != nil)
	is.True(errors.Is(err, reflector.ErrNoName))
	is.Equal(name, "")
	name, err = reflector.TypeOf([1]float32{})
	is.True(err != nil)
	is.True(errors.Is(err, reflector.ErrNoName))
	is.Equal(name, "")
	name, err = reflector.TypeOf([1]float64{})
	is.True(err != nil)
	is.True(errors.Is(err, reflector.ErrNoName))
	is.Equal(name, "")
	name, err = reflector.TypeOf([1]complex64{})
	is.True(err != nil)
	is.True(errors.Is(err, reflector.ErrNoName))
	is.Equal(name, "")
	name, err = reflector.TypeOf([1]complex128{})
	is.True(err != nil)
	is.True(errors.Is(err, reflector.ErrNoName))
	is.Equal(name, "")
	name, err = reflector.TypeOf([1]bool{})
	is.True(err != nil)
	is.True(errors.Is(err, reflector.ErrNoName))
	is.Equal(name, "")
	name, err = reflector.TypeOf([1]string{})
	is.True(err != nil)
	is.True(errors.Is(err, reflector.ErrNoName))
	is.Equal(name, "")
	name, err = reflector.TypeOf([1]byte{})
	is.True(err != nil)
	is.True(errors.Is(err, reflector.ErrNoName))
	is.Equal(name, "")
	name, err = reflector.TypeOf([1]rune{})
	is.True(err != nil)
	is.True(errors.Is(err, reflector.ErrNoName))
	is.Equal(name, "")
	name, err = reflector.TypeOf([1]error{})
	is.True(err != nil)
	is.True(errors.Is(err, reflector.ErrNoName))
	is.Equal(name, "")
	name, err = reflector.TypeOf([1]interface{}{})
	is.True(err != nil)
	is.True(errors.Is(err, reflector.ErrNoName))
	is.Equal(name, "")
	name, err = reflector.TypeOf(struct{}{})
	is.True(err != nil)
	is.True(errors.Is(err, reflector.ErrNoName))
	is.Equal(name, "")
	name, err = reflector.TypeOf(struct {
		field1 int
		field2 string
	}{})
	is.True(err != nil)
	is.True(errors.Is(err, reflector.ErrNoName))
	is.Equal(name, "")
	name, err = reflector.TypeOf(map[string]int{})
	is.True(err != nil)
	is.True(errors.Is(err, reflector.ErrNoName))
	is.Equal(name, "")
	name, err = reflector.TypeOf(map[string]string{})
	is.True(err != nil)
	is.True(errors.Is(err, reflector.ErrNoName))
	is.Equal(name, "")
	name, err = reflector.TypeOf(map[string]bool{})
	is.True(err != nil)
	is.True(errors.Is(err, reflector.ErrNoName))
	is.Equal(name, "")
	name, err = reflector.TypeOf(map[string]float32{})
	is.True(err != nil)
	is.True(errors.Is(err, reflector.ErrNoName))
	is.Equal(name, "")
	name, err = reflector.TypeOf(map[string]float64{})
	is.True(err != nil)
	is.True(errors.Is(err, reflector.ErrNoName))
	is.Equal(name, "")
	name, err = reflector.TypeOf(map[string]complex64{})
	is.True(err != nil)
	is.True(errors.Is(err, reflector.ErrNoName))
	is.Equal(name, "")
	name, err = reflector.TypeOf(map[string]complex128{})
	is.True(err != nil)
	is.True(errors.Is(err, reflector.ErrNoName))
	is.Equal(name, "")
	name, err = reflector.TypeOf(map[string]interface{}{})
	is.True(err != nil)
	is.True(errors.Is(err, reflector.ErrNoName))
	is.Equal(name, "")
	name, err = reflector.TypeOf(int(42))
	is.True(err != nil)
	is.True(errors.Is(err, reflector.ErrNoName))
	is.Equal(name, "")
	name, err = reflector.TypeOf(float64(42.0))
	is.True(err != nil)
	is.True(errors.Is(err, reflector.ErrNoName))
	is.Equal(name, "")
	name, err = reflector.TypeOf(true)
	is.True(err != nil)
	is.True(errors.Is(err, reflector.ErrNoName))
	is.Equal(name, "")
	name, err = reflector.TypeOf("hello")
	is.True(err != nil)
	is.True(errors.Is(err, reflector.ErrNoName))
	is.Equal(name, "")
	name, err = reflector.TypeOf([]int{1, 2, 3})
	is.True(err != nil)
	is.True(errors.Is(err, reflector.ErrNoName))
	is.Equal(name, "")
	name, err = reflector.TypeOf(customAlias(10))
	is.True(err != nil)
	is.True(errors.Is(err, reflector.ErrNoName))
	is.Equal(name, "")
}
