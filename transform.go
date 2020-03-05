// Copyright 2020 stringsx Author(https://github.com/yudeguang/stringsx). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/yudeguang/stringsx.
package stringsx

import (
	"bytes"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io/ioutil"
)

type Charset string

const (
	UTF8    = Charset("UTF-8")
	GB18030 = Charset("GB18030")
	GBK     = Charset("GBK")
	GB2312  = Charset("HZGB2312")
)

//把其它中文类型转换为UTF8
func toUTF8(bt []byte, charset Charset) string {
	switch charset {
	case GB18030:
		bt, _ = simplifiedchinese.GB18030.NewDecoder().Bytes(bt)
	case GBK:
		bt, _ = simplifiedchinese.GBK.NewDecoder().Bytes(bt)
	case GB2312:
		bt, _ = simplifiedchinese.HZGB2312.NewDecoder().Bytes(bt)
	default:
		panic("unsupport type")
	}
	return string(bt)
}

//GB18030转换为UTF8，应优先使用GB18030,GB18030为GBK的超集,GBK又为GB2312的超集
func GB18030ToUTF8(s string) string {
	return toUTF8([]byte(s), GB18030)
}

//GBK转换为utf8
func GBKToUTF8(s string) string {
	return toUTF8([]byte(s), GBK)
}

//GB2312转换为utf8
func GB2312ToUTF8(s string) string {
	return toUTF8([]byte(s), GB2312)
}

//UTF8转GB18030，注意，打印出来一般是乱码
func UTF8ToGB18030(s string) string {
	data, err := ioutil.ReadAll(transform.NewReader(bytes.NewReader([]byte(s)), simplifiedchinese.GB18030.NewEncoder()))
	if err != nil {
		panic(err)
	}
	return string(data)
}

//UTF8转GBK，注意，打印出来一般是乱码
func UTF8ToGBK(s string) string {
	data, err := ioutil.ReadAll(transform.NewReader(bytes.NewReader([]byte(s)), simplifiedchinese.GBK.NewEncoder()))
	if err != nil {
		panic(err)
	}
	return string(data)
}

//UTF8转GB2312，注意，打印出来一般是乱码
func UTF8ToGB2312(s string) string {
	data, err := ioutil.ReadAll(transform.NewReader(bytes.NewReader([]byte(s)), simplifiedchinese.HZGB2312.NewEncoder()))
	if err != nil {
		panic(err)
	}
	return string(data)
}
