/*
 * Name : Okto Prima Jaya
 * Github : https://github.com/oktopriima
 * Email : octoprima93@gmail.com
 * Created At : 18/11/2019, 20:52
 * Copyright (c) 2019
 */

package helper

import (
	"crypto/md5"
	"encoding/hex"
)

// GetMD5Hash accepts string input to hash with MD5.
// Return hash string of input.
func GetMD5Hash(input string) string {
	hash := md5.Sum([]byte(input))
	return hex.EncodeToString(hash[:])
}
