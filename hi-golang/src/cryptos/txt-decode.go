package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
)

// @Date：   2019/1/27 0027 16:23
// @Author:  Joshua Conero
// @Name:    名称描述

// md5
func MdsEnApp(s string) string {
	dd := md5.Sum([]byte(s))
	return fmt.Sprintf("%x", dd)
}

// sha1
func Sha1EnApp(s string) string {
	dd := sha1.Sum([]byte(s))
	return fmt.Sprintf("%x", dd)
}

// sha256
func sha256EnApp(s string) string {
	dd := sha256.Sum256([]byte(s))
	return fmt.Sprintf("%x", dd)
}


// sha512
func sha512EnApp(s string) string {
	dd := sha512.Sum512([]byte(s))
	return fmt.Sprintf("%x", dd)
}
