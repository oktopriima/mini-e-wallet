/*
 * Name : Okto Prima Jaya
 * Github : https://github.com/oktopriima
 * Email : octoprima93@gmail.com
 * Created At : 18/11/2019, 20:52
 * Copyright (c) 2019
 */

package helper

import (
	"math/rand"
	"net"
	"net/http"

	"github.com/oktopriima/mini-e-wallet/domain/middleware"
)

func RandString(n int) string {
	const letterBytes = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}

	return string(b)
}

func GetAuthenticatedUser(r *http.Request) (int64, error) {
	userID, err := middleware.ExtractToken(r, "user_id")
	if err != nil {
		return 0, err
	}
	return int64(userID.(float64)), nil
}

func GetLimitOffset(page, size int) (limit int, offset int) {
	if page == 0 || size == 0 {
		// using -1 to disable gorm size and offset in case page and size not set
		size = -1
		offset = -1
		return size, offset
	}
	offset = (page - 1) * size
	return size, offset
}

func GetIP(r *http.Request) string {
	var ip net.IP

	nets, err := net.Interfaces()
	if err != nil {
		return ""
	}
	// handle err
	for _, i := range nets {
		addr, err := i.Addrs()
		if err != nil {
			return ""
		}
		// handle err
		for _, addr := range addr {
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}
		}
	}

	if ip == nil {
		return ""
	}

	return ip.String()
}
