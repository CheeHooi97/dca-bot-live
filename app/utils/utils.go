package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	config "dca-bot-live/app/config"
	"dca-bot-live/app/errcode"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"log"
	r "math/rand"
	"regexp"
	"strconv"
	"strings"
	"time"
	"unicode"

	"github.com/labstack/echo/v4"
	"github.com/sony/sonyflake"
)

var sf = sonyflake.NewSonyflake(sonyflake.Settings{})

func ValidateRequest(c echo.Context, i any) (string, error) {
	if err := c.Bind(i); err != nil {
		return errcode.InvalidRequest.Message, err
	}

	if err := c.Validate(i); err != nil {
		return errcode.ValidationError.Message, err
	}

	return "", nil
}

func EncryptAES(plaintext string) (string, error) {
	keyBytes, err := base64.StdEncoding.DecodeString(config.SystemAesKey)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(keyBytes)
	if err != nil {
		return "", err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonce := make([]byte, aesGCM.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	ciphertext := aesGCM.Seal(nil, nonce, []byte(plaintext), nil)

	encrypted := append(nonce, ciphertext...)

	return base64.StdEncoding.EncodeToString(encrypted), nil
}

func DecryptAES(encryptedText string) (string, error) {
	keyBytes, err := base64.StdEncoding.DecodeString(config.SystemAesKey)
	if err != nil {
		return "", err
	}

	data, err := base64.StdEncoding.DecodeString(encryptedText)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(keyBytes)
	if err != nil {
		return "", err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonceSize := aesGCM.NonceSize()
	if len(data) < nonceSize {
		return "", errors.New(errcode.InvalidEncryptedText.Message)
	}

	nonce, ciphertext := data[:nonceSize], data[nonceSize:]

	plaintext, err := aesGCM.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return "", err
	}

	return string(plaintext), nil
}

// random
var randomSeed = r.New(r.NewSource(time.Now().UnixNano()))

func Alphanumeric(n int) string {
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890-")
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[randomSeed.Intn(len(letterRunes))]
	}
	return string(b)
}

func Numeric(n int) string {
	var letterRunes = []rune("1234567890")
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[randomSeed.Intn(len(letterRunes))]
	}

	s := string(b)
	return s
}

func DecodeKey(encoded string) (string, error) {
	// Re-add padding.
	if m := len(encoded) % 4; m != 0 {
		encoded += strings.Repeat("=", 4-m)
	}

	decoded, err := base64.URLEncoding.DecodeString(encoded)
	if err != nil {
		return "", err
	}

	parts := strings.SplitN(string(decoded), ":", 2)
	if len(parts) != 2 {
		return "", fmt.Errorf("invalid key format")
	}

	return parts[1], nil
}

func Contains(slice []string, target string) bool {
	for _, item := range slice {
		if item == target {
			return true
		}
	}
	return false
}

func ConvertLocalTimeToUTC(dt string) (*time.Time, error) {
	t, err := time.Parse(time.DateOnly, dt)
	if err != nil {
		return nil, err
	}

	localTime := time.Date(
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second(), t.Nanosecond(),
		time.Local,
	)

	utcTime := localTime.UTC()

	return &utcTime, nil
}

func IsRentalFeeInRange(rentalFee int64, priceRange string) bool {
	priceRange = strings.TrimPrefix(priceRange, "RM ")
	priceRange = strings.TrimSpace(priceRange)

	parts := strings.Split(priceRange, "-")
	if len(parts) != 2 {
		return false
	}

	min, err1 := strconv.ParseInt(strings.TrimSpace(parts[0]), 10, 64)
	max, err2 := strconv.ParseInt(strings.TrimSpace(parts[1]), 10, 64)

	if err1 != nil || err2 != nil {
		return false
	}

	return rentalFee >= min && rentalFee <= max
}

func FormatPhoneNumber(phone string) string {
	if strings.HasPrefix(phone, "+60") && len(phone) > 3 {
		return "0" + phone[3:]
	}
	return phone
}

func CheckAllValuesInArray(source, target []string) bool {
	set := make(map[string]struct{}, len(source))
	for _, s := range source {
		set[strings.ToLower(s)] = struct{}{}
	}

	for _, t := range target {
		if _, found := set[strings.ToLower(t)]; !found {
			return false
		}
	}
	return true
}

func ValidatePin(pin string) bool {
	// Define a regular expression pattern for a six-digit PIN
	pattern := "^[0-9]{6}$"

	// Compile the regular expression
	regex, err := regexp.Compile(pattern)
	if err != nil {
		return false
	}

	// Use the regular expression to match the PIN
	return regex.MatchString(pin)
}

func ToCamelCase(input string) string {
	words := strings.Fields(input)
	for i, w := range words {
		runes := []rune(w)
		if len(runes) > 0 {
			runes[0] = unicode.ToUpper(runes[0])
			words[i] = string(runes)
		}
	}
	return strings.Join(words, "")
}

func ExtractCity(address string) string {
	parts := strings.Split(address, ",")
	for i := range parts {
		parts[i] = strings.TrimSpace(parts[i])
	}

	re := regexp.MustCompile(`\d+`)

	for i := len(parts) - 1; i >= 0; i-- {
		part := parts[i]

		if strings.EqualFold(part, "Malaysia") {
			continue
		}

		if strings.EqualFold(part, "Kuala Lumpur") {
			return "Kuala Lumpur"
		}

		return strings.TrimSpace(re.ReplaceAllString(part, ""))
	}
	return ""
}

func TenancyEndDate(start *time.Time, months int) *time.Time {
	end := start.AddDate(0, months, 0)

	endDate := end.AddDate(0, 0, -1)

	return &endDate
}

func FormatCamelCase(name string) string {
	var result []rune
	for i, r := range name {
		if unicode.IsUpper(r) && i > 0 {
			result = append(result, ' ')
		}
		result = append(result, r)
	}

	words := strings.Fields(string(result))
	for i, w := range words {
		runes := []rune(w)
		if len(runes) > 0 {
			runes[0] = unicode.ToUpper(runes[0])
			words[i] = string(runes)
		}
	}
	return strings.Join(words, " ")
}

func LowerFirst(s string) string {
	if s == "" {
		return s
	}
	runes := []rune(s)
	runes[0] = unicode.ToLower(runes[0])
	return string(runes)
}

func FormatDate1(date *time.Time) string {
	return date.Format("2006-01-02")
}

func FormatDate2(date *time.Time) string {
	return date.Format("02 Jan 2006")
}

func FormatTenurePeriod(period string) string {
	var p string
	switch period {
	case "3_MONTH", "3 Months", "3 months":
		p = "3 months"
	case "6_MONTH", "6 Months", "6 months":
		p = "6 months"
	case "12_MONTH", "12 Months", "12 months":
		p = "12 months"
	}
	return p
}

func FormatTenureEndDate(period string, date *time.Time) time.Time {
	var endDate time.Time
	switch period {
	case "3_MONTH", "3 Months", "3 months":
		endDate = date.AddDate(0, 3, -1)
	case "6_MONTH", "6 Months", "6 months":
		endDate = date.AddDate(0, 6, -1)
	case "12_MONTH", "12 Months", "12 months":
		endDate = date.AddDate(1, 0, -1)
	}
	return endDate
}

// GenerateUNIQ :
func UniqueID() string {
	id, err := sf.NextID()
	if err != nil {
		log.Fatal(err)
	}
	return strconv.FormatUint(id, 10)
}
