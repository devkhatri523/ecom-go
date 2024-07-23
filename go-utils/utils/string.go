package utils

import (
	"encoding/json"
	"github.com/google/uuid"
	"log"

	"html"
	"math"
	"net/mail"
	"os"
	"reflect"
	"strconv"
	"strings"
)

func FormatAmount(amount int64) string {
	a := GetDecimalAmount(amount)
	return strconv.FormatFloat(a, 'f', 2, 64)
}
func GetDecimalAmount(amount int64) float64 {
	if amount == 0 {
		return 0
	}
	result := float64(amount) / 100.0
	return result
}
func GetAbsAmountFromStr(amount string) int64 {
	if IsBlank(amount) {
		return 0
	}
	amount = strings.ReplaceAll(amount, ",", "")
	f, err := strconv.ParseFloat(amount, 64)
	if err != nil {
		panic(err)
	}
	return GetAbsAmount(f)
}
func GetFloatAmount(amount string) float64 {
	if IsBlank(amount) {
		return 0
	}
	amount = strings.ReplaceAll(amount, ",", "")
	f, err := strconv.ParseFloat(amount, 64)
	if err != nil {
		panic(err)
	}
	return f
}
func GetAbsAmount(amount float64) int64 {
	return int64(math.Abs(amount) * 100)
}
func ParseNumber(amount string) int64 {
	if IsBlank(amount) {
		return 0
	}
	amount = strings.ReplaceAll(amount, ",", "")
	amount = strings.ReplaceAll(amount, ".", "")
	f, err := strconv.ParseInt(amount, 10, 64)
	if err != nil {
		panic(err)
	}
	return f
}
func GetAmountFromStr(amount string) int64 {
	if IsBlank(amount) {
		return 0
	}
	amount = strings.ReplaceAll(amount, ",", "")
	f, err := strconv.ParseFloat(amount, 64)
	if err != nil {
		panic(err)
	}
	return GetAmount(f)
}
func GetAmount(amount float64) int64 {
	return int64(math.Round(amount * 100))
}
func GetStructField(s any, fieldName string) reflect.StructField {
	f, ok := reflect.TypeOf(s).Elem().FieldByName(fieldName)
	if ok {
		return f
	} else {
		panic(fieldName + " not found")
	}
}
func GetStructTag(f reflect.StructField, tagName string) string {
	return f.Tag.Get(tagName)
}

func GetHostname() string {
	h, err := os.Hostname()
	if err != nil {
		return ""
	}
	return h
}

func CreateJsonStr(t interface{}) string {
	j, err := json.Marshal(t)
	if err != nil {
		log.Fatal(err)
		return ""
	}
	return string(j)
}
func CreateJsonFromStr(jsonStr string) map[string]interface{} {
	var jsonMap map[string]interface{}
	err := json.Unmarshal([]byte(jsonStr), &jsonMap)
	if err != nil {
		log.Fatal(err)
	}
	return jsonMap
}
func FromJsonToAny(jsonStr string, v any) error {
	return json.Unmarshal([]byte(jsonStr), v)
}

func GenerateUUID() string {
	id := uuid.New()
	return id.String()
}

func IsBlank(data string) bool {
	return data == ""
}

func IsNotBlank(data string) bool {
	return data != ""
}

func EscapeHtml(data string) string {
	if IsBlank(data) {
		return data
	} else {
		return html.EscapeString(data)
	}
}

func CompareTwoStrings(first string, second string) bool {
	if len(first) != len(second) {
		return false
	}
	firstByte := []byte(first)
	secondByte := []byte(second)
	var ret byte
	for i, x := range secondByte {
		ret |= x ^ firstByte[i]
	}
	return ret == 0
}

func IsValidEmail(email string) bool {
	if IsBlank(email) {
		return false
	}
	_, err := mail.ParseAddress(email)
	return err == nil
}

func DefaultString(str string, defaultStr string) string {
	if IsBlank(str) {
		return defaultStr
	} else {
		return str
	}
}
