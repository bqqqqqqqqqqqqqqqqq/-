package helping

import (
	"Doggggg/Init"
	"crypto/md5"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/jordan-wright/email"
	uuid "github.com/satori/go.uuid"
	"math/rand"
	"net/smtp"
	"strconv"
	"time"
)

type UserClaims struct {
	Uuid  string `json:"identity"`
	Name  string `json:"name"`
	Phone string `json:"phone"`
	jwt.StandardClaims
}

// GetMd5
// 生成 md5
func GetMd5(s string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(s)))
}

var myKey = []byte("gin-gorm-oj-key")

// GenerateToken
// 生成 token
func GenerateToken(uuid string, name string, phone string) (string, error) {
	UserClaim := &UserClaims{
		Uuid:           uuid,
		Name:           name,
		Phone:          phone,
		StandardClaims: jwt.StandardClaims{},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, UserClaim)
	tokenString, err := token.SignedString(myKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// AnalyseToken
// 解析 token
func AnalyseToken(tokenString string) (*UserClaims, error) {
	userClaim := new(UserClaims)
	claims, err := jwt.ParseWithClaims(tokenString, userClaim, func(token *jwt.Token) (interface{}, error) {
		return myKey, nil
	})
	if err != nil {
		return nil, err
	}
	if !claims.Valid {
		return nil, fmt.Errorf("analyse Token Error:%v", err)
	}
	return userClaim, nil
}

// SendCodeHelp
// 发送验证码
func SendCodeHelp(c *gin.Context, toUserEmail string, code string) error {
	e := email.NewEmail()
	e.From = "get <qiaobqqq@gmail.com>"
	e.To = []string{toUserEmail}
	e.Subject = "测试"
	e.HTML = []byte("您的验证码是：<b>" + code + "</b>")
	err := e.Send("smtp.gmail.com:587", smtp.PlainAuth("", "qiaobqqq@gmail.com", "stmrqvrshcxfykbu", "smtp.gmail.com"))
	if err != nil {
		return err
	}
	Init.RDB.Set(c, toUserEmail, code, time.Second*300)
	return nil
}

// GetUUID
// 生成唯一码
func GetUUID() string {
	return uuid.NewV4().String()
}

// GetRand
// 生成验证码
func GetRand() string {
	rand.Seed(time.Now().UnixNano())
	s := ""
	for i := 0; i < 6; i++ {
		s += strconv.Itoa(rand.Intn(10))
	}
	return s
}
