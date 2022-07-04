package utils

import (
	"fmt"
	"math"
	"strconv"
)

//格式化
func RoundFloat(f float64, n int) float64 {
	n10 := math.Pow10(n)
	return math.Trunc((f+0.5/n10)*n10) / n10
}

//float64*100
func Float64C100(totalAmount string) float64 {
	totalAmount2 := StringToFloat64(totalAmount) * 100
	return RoundFloat(totalAmount2, 0)
}

//float64*100
func Float64C100ToString(totalAmount string) string {
	totalAmount2 := StringToFloat64(totalAmount) * 100
	return Float64ToString(RoundFloat(totalAmount2, 0))
}

//float64*100
func Float64C100ToStr(totalAmount float64) string {
	totalAmount2 := totalAmount * 100
	return Float64ToString(RoundFloat(totalAmount2, 0))
}

//float64
func Float64ToStr(totalAmount float64) string {
	return Float64ToString(RoundFloat(totalAmount, 0))
}

//float64/100
func Float64D100(f float64) float64 {
	f2 := f / 100
	return RoundFloat(f2, 2)
}

//64浮点数四舍五入
func Float64Decimal(value float64, bit int) float64 {
	target := "%." + IntToString(bit) + "f"
	value, _ = strconv.ParseFloat(fmt.Sprintf(target, value), 64)
	return value
}

//32浮点数四舍五入
func Float32Decimal(value float32, bit int) float32 {
	target := "%." + IntToString(bit) + "f"
	result, _ := strconv.ParseFloat(fmt.Sprintf(target, value), 32)
	return float32(result)
}
