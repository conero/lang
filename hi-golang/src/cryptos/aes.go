package main

import (
	"crypto/aes"
	"log"
)

// @Date：   2019/1/27 0027 17:34
// @Author:  Joshua Conero
// @Name:    名称描述

func aesEnTxt(d string, key string) string{
	c, e := aes.NewCipher([]byte(key))
	if e != nil{
		log.Fatal(e.Error())
	}
	var dst []byte
	c.Encrypt(dst, []byte(d))
	return string(dst)
}