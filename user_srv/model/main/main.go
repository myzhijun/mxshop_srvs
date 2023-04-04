package main

import (
	"crypto/sha512"
	"fmt"
	"github.com/anaskhan96/go-password-encoder"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"mxshop_srvs/user_srv/model"
	"os"
	"strings"
	"time"
)

func main() {
	dsn := "root:root@tcp(127.0.0.1:3306)/mxshop_user_srv?charset=utf8mb4&parseTime=True"
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Info,
			Colorful:      true,
		})
	//全局模式
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic(err)
	}
	_ = db.AutoMigrate(&model.User{})
}

//加密算法
func genMd5(code string) string {
	//Md5 := md5.New()
	//_, _ = io.WriteString(Md5, code)
	//return hex.EncodeToString(Md5.Sum(nil))
	options := &password.Options{16, 100, 32, sha512.New}
	salt, encodedPwd := password.Encode("generic password", options)
	newPassword := fmt.Sprintf("$pdkdf2-sha512$%s$%s", salt, encodedPwd)
	return newPassword
}

//解密
func CheckPwd(encodepwd string) bool {
	options := &password.Options{16, 100, 32, sha512.New}
	pwd := strings.Split(encodepwd, "$")
	salt := strings.Split(pwd[0], "-")[1]
	check := password.Verify("generic password", salt, pwd[1], options)
	return check
}
