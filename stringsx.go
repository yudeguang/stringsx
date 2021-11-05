// Copyright 2020 stringsx Author(https://github.com/yudeguang/stringsx). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/yudeguang/stringsx.
//字符串处理包，对标准库字符串的补充
package stringsx

import (
	iox "github.com/yudeguang/iox"
	"math/rand"
	"strconv"
	"strings"
	"time"
	"unicode"
	"fmt"
)

//返回第一次出现sep之后的字串符
func After(s, sep string) string {
	if s == "" || sep == "" {
		return s
	}
	pos := strings.Index(s, sep)
	if pos == -1 {
		return ""
	}
	return s[pos+len(sep):]
}

//返回最后一次出现sep之后的字符串
func AfterLast(s, sep string) string {
	if sep == "" || s == "" {
		return s
	}
	pos := strings.LastIndex(s, sep)
	if pos == -1 {
		return ""
	}
	return s[pos+len(sep):]
}

//返回第N次出现sep之后的字符串
func AfterNSep(s, sep string, nTimes int) string {
	if sep == "" || s == "" || nTimes <= 0 {
		return s
	}
	f := iox.New(strings.NewReader(s))
	pos := int(f.IndexN(0, []byte(sep), nTimes))
	if pos == -1 {
		return ""
	}
	return s[pos+len(sep):]
}

//返回第一次出现sep之前的字符串
func Before(s, sep string) string {
	if s == "" || sep == "" {
		return s
	}
	pos := strings.Index(s, sep)
	if pos == -1 {
		return ""
	}
	return s[:pos]
}

//返回最后一次出现sep之前的字符串
func BeforeLast(s, sep string) string {
	if sep == "" || s == "" {
		return s
	}
	pos := strings.LastIndex(s, sep)
	if pos == -1 {
		return ""
	}
	return s[:pos]
}

//返回第N次出现sep之前的字符串
func BeforeNSep(s, sep string, nTimes int) string {
	if sep == "" || s == "" || nTimes <= 0 {
		return s
	}
	f := iox.New(strings.NewReader(s))
	pos := int(f.IndexN(0, []byte(sep), nTimes))
	if pos == -1 {
		return ""
	}
	return s[:pos]
}

//返回第一次出现在两个字符串接之间的字符串
func Between(s, begin, end string) string {
	if s == "" || begin == "" || end == "" {
		return ""
	}
	beginPos := strings.Index(s, begin)
	if beginPos != -1 {
		f := iox.New(strings.NewReader(s))
		endPos := int(f.IndexGen(int64(beginPos+len(begin)), int64(len(s)-1), []byte(end)))
		if endPos != -1 {
			return s[beginPos+len(begin) : endPos]
		}
	}
	return ""
}

//返回左侧N个字符
func Left(s string, n int) string {
	if n <= 0 || s == "" {
		return ""
	}
	runes := []rune(s)
	if len(runes) <= n {
		return s
	}
	return string(runes[0:n])
}

//返回右侧N个字符
func Right(s string, n int) string {
	if n <= 0 || s == "" {
		return ""
	}
	runes := []rune(s)
	if len(runes) <= n {
		return s
	}
	return string(runes[len(runes)-n:])
}

//用分隔符sep把若干个字符拼接在一起,实际为strings.Join的变体形式
func JoinStrings(sep string, args ...string) string {
	return strings.Join(args, sep)
}

//用分隔符号把若干个数字排接在一起
/*
func JoinInts(sep string, args ...int) string {
	l := len(args)
	switch l {
	case 0:
		return ""
	case 1:
		return strconv.Itoa(args[0])
	case 2:
		return strconv.Itoa(args[0]) + sep + strconv.Itoa(args[1])
	case 3:
		return strconv.Itoa(args[0]) + sep + strconv.Itoa(args[1]) + sep + strconv.Itoa(args[2])
	}
	var buffer bytes.Buffer
	//前面若干条中间要加sep
	for i := 0; i < l-1; i++ {
		buffer.WriteString(strconv.Itoa(args[i]) + sep)
	}
	//最后次不加sep
	buffer.WriteString(strconv.Itoa(args[l-1]))
	return buffer.String()
}
 */

func JoinInts( sep string,elems ...int) string {
	switch len(elems) {
	case 0:
		return ""
	case 1:
		return strconv.Itoa(elems[0])
	}
	n := len(sep) * (len(elems) - 1)
	for i := 0; i < len(elems); i++ {
		n += len( strconv.Itoa(elems[i]))
	}
	var b strings.Builder
	b.Grow(n)
	b.WriteString(strconv.Itoa(elems[0]) )
	for _, s := range elems[1:] {
		b.WriteString(sep)
		b.WriteString(strconv.Itoa(s))
	}
	return b.String()
}
//用分隔符sep把若干个字符或int,double等类型数据拼接在一起,实际为strings.Join的变体形式
/*
func JoinInterface(sep string, args ...interface{}) string {
	l := len(args)
	switch l {
	case 0:
		return ""
	case 1:
		return fmt.Sprint(args[0])
	case 2:
		return fmt.Sprint(args[0]) + sep + fmt.Sprint(args[1])
	case 3:
		return fmt.Sprint(args[0]) + sep + fmt.Sprint(args[1]) + sep + fmt.Sprint(args[2])
	}
	var buffer bytes.Buffer
	//前面若干条中间要加sep
	for i := 0; i < l-1; i++ {
		buffer.WriteString(fmt.Sprint(args[i]) + sep)
	}
	//最后次不加sep
	buffer.WriteString(fmt.Sprint(args[l-1]))
	return buffer.String()
}
*/
func JoinInterface( sep string,elems  ...interface{}) string {
	switch len(elems) {
	case 0:
		return ""
	case 1:
		return  fmt.Sprint(elems[0])
	}
	n := len(sep) * (len(elems) - 1)
	for i := 0; i < len(elems); i++ {
		n += len( fmt.Sprint(elems[i]))
	}
	var b strings.Builder
	b.Grow(n)
	b.WriteString(fmt.Sprint(elems[0]) )
	for _, s := range elems[1:] {
		b.WriteString(sep)
		b.WriteString(fmt.Sprint(s))
	}
	return b.String()
}
//返回倒序字符串
func Reverse(s string) string {
	runes := []rune(s)
	for from, to := 0, len(runes)-1; from < to; from, to = from+1, to-1 {
		runes[from], runes[to] = runes[to], runes[from]
	}
	return string(runes)
}

//返回把中间字符按一定规则替换后的字符串接
//CenterPad("hello", 10, "*") => "he*****llo"
func CenterPad(s string, length int, pad string) string {
	l := lenRune(s)
	if length <= l {
		return Left(s, length)
	}
	//对应mysql中是返回NULL
	if pad == "" && length >= l {
		return ""
	}
	//取得重复若干次pad之后剩余的文本"12312", Rightpad("hello", 10, "123") => "hello12312"
	pads := Right(Rightpad(s, length, pad), length-l)
	//判断s长度是单数还是双数
	remainder := l % 2
	quotient := l / 2
	if remainder == 0 {
		return Left(s, quotient) + pads + Right(s, quotient)
	} else {
		return Left(s, quotient) + pads + Right(s, quotient+1)
	}
}

//返回两侧字符按一定规则替换后的字符串接
// LeftRightPad("hello", 4, " ")    => "hell"
// LeftRightPad("hello", 10, " ")   => "  hello   "
// LeftRightPad("hello", 10, "123") => "12hello123"
func LeftRightPad(s string, length int, pad string) string {
	l := lenRune(s)
	if length <= l {
		return Left(s, length)
	}
	//对应mysql中是返回NULL
	if pad == "" && length >= l {
		return ""
	}
	//取得重复若干次pad之后剩余的文本"12312", Rightpad("hello", 10, "123") => "hello12312"
	pads := Right(Rightpad(s, length, pad), length-l)
	//判断pads长度是单数还是双数
	remainder := (length - l) % 2
	quotient := (length - l) / 2
	if remainder == 0 {
		return Left(pads, quotient) + s + Right(pads, quotient)
	} else {
		return Left(pads, quotient) + s + Right(pads, quotient+1)
	}
}

//返回字符串str，右面用字符串padstr填补直到str是len个字符长,此函数与mysql中RPAD()行为保持一致
// Rightpad("hello", 4, " ")    => "hello"
// Rightpad("hello", 10, " ")   => "hello     "
// Rightpad("hello", 10, "123") => "hello12312"
func Rightpad(s string, length int, pad string) string {
	l := lenRune(s)
	if length <= l {
		return Left(s, length)
	}
	//对应mysql中是返回NULL
	if pad == "" && length >= l {
		return ""
	}
	for {
		if lenRune(s) >= length {
			return Left(s, length)
		}
		s = s + pad
	}
}

//返回字符串str，左面用字符串padstr填补直到str是len个字符长,,此函数与mysql中LPAD()行为保持一致
// LeftPad("hello", 4, " ")    => "hello"
// LeftPad("hello", 10, " ")   => "     hello"
// LeftPad("hello", 10, "123") => "12312hello"
func LeftPad(s string, length int, pad string) string {
	l := lenRune(s)
	if length <= l {
		return Right(s, length)
	}
	//对应mysql中是返回NULL
	if pad == "" && length >= l {
		return ""
	}
	for {
		if lenRune(s) >= length {
			return Right(s, length)
		}
		s = pad + s
	}
}

//返回s中的字符数
func lenRune(s string) int {
	return len([]rune(s))
}

//返回随机打乱后的字符串
func Rand(s string) string {
	if s == "" {
		return s
	}
	runes := []rune(s)
	rand.Seed(time.Now().UnixNano())
	randedSlice := rand.Perm(len(runes))
	newRunes := make([]rune, 0, len(runes))
	for _, randedIndex := range randedSlice {
		newRunes = append(newRunes, runes[randedIndex])
	}
	return string(newRunes)
}

const numbersAndLetters = `abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789`
const commaAndNumbersAndLetters = `,abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789`
const numbers = `0123456789`
const letters = `abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ`

//只保留数字和英文字母,删除其它类型字母及标点符号
func NumbersLettersLeft(s string) string {
	runes := []rune(s)
	newRunes := make([]rune, 0, len(runes))
	for _, r := range runes {
		if strings.Contains(numbersAndLetters, string(r)) {
			newRunes = append(newRunes, r)
		}
	}
	return string(newRunes)
}

//只保留阿拉伯数字
func NumbersLeft(s string) string {
	runes := []rune(s)
	newRunes := make([]rune, 0, len(runes))
	for _, r := range runes {
		if strings.Contains(numbers, string(r)) {
			newRunes = append(newRunes, r)
		}
	}
	return string(newRunes)
}

//只保留英文字母
func LettersLeft(s string) string {
	runes := []rune(s)
	newRunes := make([]rune, 0, len(runes))
	for _, r := range runes {
		if strings.Contains(letters, string(r)) {
			newRunes = append(newRunes, r)
		}
	}
	return string(newRunes)
}

//保留输入的相关字符
func RelevantCharactersLeft(s, RelevantCharacter string) string {
	runes := []rune(s)
	newRunes := make([]rune, 0, len(runes))
	for _, r := range runes {
		if strings.Contains(RelevantCharacter, string(r)) {
			newRunes = append(newRunes, r)
		}
	}
	return string(newRunes)
}

//只保留逗号以及数字和英文字母，因为逗号一般用于分隔文本
func CommaNumbersLettersLeft(s string) string {
	runes := []rune(s)
	newRunes := make([]rune, 0, len(runes))
	for _, r := range runes {
		if strings.Contains(commaAndNumbersAndLetters, string(r)) {
			newRunes = append(newRunes, r)
		}
	}
	return string(newRunes)
}

//移除字符串中的汉字
func RemoveHan(s string) string {
	runes := []rune(s)
	var newRunes []rune
	for _, r := range runes {
		if !RuneIsHan(r) {
			newRunes = append(newRunes, r)
		}
	}
	return string(newRunes)
}

//只保留字符串中的汉字
func HanLeft(s string) string {
	runes := []rune(s)
	var newRunes []rune
	for _, r := range runes {
		if RuneIsHan(r) {
			newRunes = append(newRunes, r)
		}
	}
	return string(newRunes)
}

//按固定的长度拆分字符串
func SplitByLen(s string, sepLen int) []string {
	x := len(s) % sepLen //余数
	y := len(s) / sepLen
	var ret_len int
	if x > 0 {
		ret_len = y + 1
	} else {
		ret_len = y
	}
	var ret = make([]string, 0, ret_len)
	cur := 0
	for {
		if len(s[cur:]) >= sepLen {
			ret = append(ret, s[cur:cur+sepLen])
			cur = cur + sepLen
		} else {
			if len(s[cur:]) > 0 {
				ret = append(ret, s[cur:])
			}
			break
		}
	}
	return ret
}

//判断单个rune是否是汉字
func RuneIsHan(r rune) bool {
	return unicode.Is(unicode.Han, r)
}

//判断rune是否包含汉字
func ContainsHan(s string) bool {
	runes := []rune(s)
	for _, r := range runes {
		if RuneIsHan(r) {
			return true
		}
	}
	return false
}

//判断字符串是否是由纯数字组成
func IsNumber(s string) bool {
	runes := []rune(s)
	for _, r := range runes {
		if !unicode.IsDigit(r) {
			return false
		}
	}
	return true
}

//判断是否是大写字母
func IsLetterUpper(b byte) bool {
	if b >= byte('A') && b <= byte('Z') {
		return true
	}
	return false
}

//判断是否是小写字母
func IsLetterLower(b byte) bool {
	if b >= byte('a') && b <= byte('z') {
		return true
	}
	return false
}
