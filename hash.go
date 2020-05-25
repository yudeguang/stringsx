// Copyright 2020 stringsx Author(https://github.com/yudeguang/stringsx). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/yudeguang/stringsx.
//字符串处理包，对标准库字符串的补充
package stringsx

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"hash/crc32"
)

// 生成md5
func MD5(str string) string {
	c := md5.New()
	c.Write([]byte(str))
	return hex.EncodeToString(c.Sum(nil))
}

//生成sha1
func SHA1(str string) string {
	c := sha1.New()
	c.Write([]byte(str))
	return hex.EncodeToString(c.Sum(nil))
}

//生成crc32
func CRC32(str string) uint32 {
	return crc32.ChecksumIEEE([]byte(str))
}
