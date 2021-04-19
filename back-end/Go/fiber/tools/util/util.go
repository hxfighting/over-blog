package util

import (
	"crypto/rand"
	"time"
	"unsafe"

	"github.com/cenkalti/backoff/v4"
	"github.com/nyaruka/phonenumbers"
	"golang.org/x/crypto/bcrypt"
)

const (
	YMDHIS = "2006-01-02 15:04:05"
	YMD    = "2006-01-02"
)

func BytesToStr(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

func StrToBytes(s string) []byte {
	x := (*[2]uintptr)(unsafe.Pointer(&s))
	l := [3]uintptr{x[0], x[1], x[1]}
	return *(*[]byte)(unsafe.Pointer(&l))
}

func Retry(f func() error, maxRetry uint8, retryInterval time.Duration) error {
	if maxRetry == 0 {
		maxRetry = 2
	} else {
		maxRetry--
	}
	if retryInterval.Seconds() == 0 {
		retryInterval = time.Second * 10
	}
	b := backoff.NewExponentialBackOff()
	b.InitialInterval = retryInterval
	bkf := backoff.WithMaxRetries(b, uint64(maxRetry))
	err := backoff.Retry(f, bkf)
	if err != nil {
		return err
	}
	return nil
}

func RandString(n int) (string, error) {
	randStr := []byte("0123456789ABCDEFGHJKLMNPQRSTUVWXYZ")
	maxrb := 255 - (256 % len(randStr))
	b := make([]byte, n)
	r := make([]byte, n+(n/4))
	i := 0
	for {
		if _, err := rand.Read(r); err != nil {
			return "", err
		}
		for _, rb := range r {
			c := int(rb)
			if c > maxrb {
				continue
			}
			b[i] = randStr[c%len(randStr)]
			i++
			if i == n {
				return *(*string)(unsafe.Pointer(&b)), nil
			}
		}
	}
}

// 验证电话号码
func ValidatePhone(phone string) bool {
	p, err := phonenumbers.Parse(phone, "CN")
	if err != nil {
		return false
	}
	return phonenumbers.IsValidNumberForRegion(p, "CN")
}

//密码加密
func PasswordHash(password string) (string, error) {
	passBytes, err := bcrypt.GenerateFromPassword(StrToBytes(password), bcrypt.DefaultCost)
	return BytesToStr(passBytes), err
}

//密码验证
func PasswordVerify(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword(StrToBytes(hash), StrToBytes(password))
	return err == nil
}

func GetDateTime(unix int64, format string) string {
	return time.Unix(unix, 0).Format(format)
}
