package stringsx

// import (
// 	"strings"
// 	"unicode"
// )

// type Stringx struct {
// 	str string
// }

// //初始化
// func New(s string) *Stringx {
// 	return &Stringx{s}
// }

// //转化为string
// func (s *Stringx) ToString() string {
// 	return s.str
// }

// //判断两个utf-8编码字符串（将unicode大写、小写、标题三种格式字符视为相同）是否相同
// func (s *Stringx) EqualFold(t string) bool {
// 	return strings.EqualFold(s.str, t)
// }

// //判断s是否有前缀字符串prefix
// func (s *Stringx) HasPrefix(prefix string) bool {
// 	return strings.HasPrefix(s.str, prefix)
// }

// //判断s是否有后缀字符串suffix
// func (s *Stringx) HasSuffix(suffix string) bool {
// 	return strings.HasSuffix(s.str, suffix)
// }

// //判断字符串s是否包含子串substr
// func (s *Stringx) Contains(substr string) bool {
// 	return strings.Contains(s.str, substr)
// }

// //判断字符串s是否包含utf-8码值r
// func (s *Stringx) ContainsRune(r rune) bool {
// 	return strings.ContainsRune(s.str, r)
// }

// //判断字符串s是否包含字符串chars中的任一字符
// func (s *Stringx) ContainsAny(chars string) bool {
// 	return strings.ContainsAny(s.str, chars)
// }

// //返回字符串s中有几个不重复的sep子串
// func (s *Stringx) Count(sep string) int {
// 	return strings.Count(s.str, sep)
// }

// //子串sep在字符串s中第一次出现的位置，不存在则返回-1
// func (s *Stringx) Index(substr string) int {
// 	return strings.Index(s.str, substr)
// }

// //字符c在s中第一次出现的位置，不存在则返回-1
// func (s *Stringx) IndexByte(c byte) int {
// 	return strings.IndexByte(s.str, c)
// }

// //unicode码值r在s中第一次出现的位置，不存在则返回-1
// func (s *Stringx) IndexRune(r rune) int {
// 	return strings.IndexRune(s.str, r)
// }

// //字符串chars中的任一utf-8码值在s中第一次出现的位置，如果不存在或者chars为空字符串则返回-1
// func (s *Stringx) IndexAny(chars string) int {
// 	return strings.IndexAny(s.str, chars)
// }

// //s中第一个满足函数f的位置i（该处的utf-8码值r满足f(r)==true），不存在则返回-1
// func (s *Stringx) IndexFunc(f func(rune) bool) int {
// 	return strings.IndexFunc(s.str, f)
// }

// //子串sep在字符串s中最后一次出现的位置，不存在则返回-1
// func (s *Stringx) LastIndex(substr string) int {
// 	return strings.LastIndex(s.str, substr)
// }

// //字符c在s中最后一次出现的位置，不存在则返回-1
// func (s *Stringx) LastIndexByte(c byte) int {
// 	return strings.LastIndexByte(s.str, c)
// }

// //字符串chars中的任一utf-8码值在s中最后一次出现的位置，如不存在或者chars为空字符串则返回-1
// func (s *Stringx) LastIndexAny(chars string) int {
// 	return strings.LastIndexAny(s.str, chars)
// }

// //s中最后一个满足函数f的unicode码值的位置i，不存在则返回-1
// func (s *Stringx) LastIndexFunc(f func(rune) bool) int {
// 	return strings.LastIndexFunc(s.str, f)
// }

// //返回s中每个单词的首字母都改为标题格式的字符串拷贝。
// //BUG: Title用于划分单词的规则不能很好的处理Unicode标点符号。
// func (s *Stringx) Title() string {
// 	s.str = strings.Title(s.str)
// 	return s.str
// }

// //返回将所有字母都转为对应的小写版本的拷贝
// func (s *Stringx) ToLower() string {
// 	s.str = strings.ToLower(s.str)
// 	return s.str
// }

// //使用_case规定的字符映射，返回将所有字母都转为对应的小写版本的拷贝
// func (s *Stringx) ToLowerSpecial(c unicode.SpecialCase) string {
// 	s.str = strings.ToLowerSpecial(c, s.str)
// 	return s.str
// }

// //返回将所有字母都转为对应的大写版本的拷贝
// func (s *Stringx) ToUpper() string {
// 	s.str = strings.ToUpper(s.str)
// 	return s.str
// }

// //使用_case规定的字符映射，返回将所有字母都转为对应的大写版本的拷贝
// func (s *Stringx) ToUpperSpecial(c unicode.SpecialCase) string {
// 	s.str = strings.ToUpperSpecial(c, s.str)
// 	return s.str
// }

// //返回将所有字母都转为对应的标题版本的拷贝
// func (s *Stringx) ToTitle() string {
// 	s.str = strings.ToTitle(s.str)
// 	return s.str
// }

// //使用_case规定的字符映射，返回将所有字母都转为对应的标题版本的拷贝。
// func (s *Stringx) ToTitleSpecial(c unicode.SpecialCase) string {
// 	s.str = strings.ToTitleSpecial(c, s.str)
// 	return s.str
// }

// //返回count个s串联的字符串
// func (s *Stringx) Repeat(count int) string {
// 	s.str = strings.Repeat(s.str, count)
// 	return s.str
// }

// //返回将s中前n个不重叠old子串都替换为new的新字符串，如果n<0会替换所有old子串
// func (s *Stringx) Replace(old, new string, n int) string {
// 	s.str = strings.Replace(s.str, old, new, n)
// 	return s.str
// }

// //将s的每一个unicode码值r都替换为mapping(r)，返回这些新码值组成的字符串拷贝。
// //如果mapping返回一个负值，将会丢弃该码值而不会被替换。（返回值中对应位置将没有码值）
// func (s *Stringx) Map(mapping func(rune) rune) string {
// 	s.str = strings.Map(mapping, s.str)
// 	return s.str
// }

// //返回将s前后端所有cutset包含的utf-8码值都去掉的字符串
// func (s *Stringx) Trim(cutset string) string {
// 	s.str = strings.Trim(s.str, cutset)
// 	return s.str
// }

// //返回将s前后端所有空白（unicode.IsSpace指定）都去掉的字符串
// func (s *Stringx) TrimSpace() string {
// 	s.str = strings.TrimSpace(s.str)
// 	return s.str
// }

// //返回将s前后端所有满足f的unicode码值都去掉的字符串
// func (s *Stringx) TrimFunc(f func(rune) bool) string {
// 	s.str = strings.TrimFunc(s.str, f)
// 	return s.str
// }

// //返回将s前端所有cutset包含的utf-8码值都去掉的字符串
// func (s *Stringx) TrimLeft(cutset string) string {
// 	s.str = strings.TrimLeft(s.str, cutset)
// 	return s.str
// }

// //返回将s前端所有满足f的unicode码值都去掉的字符串
// func (s *Stringx) TrimLeftFunc(f func(rune) bool) string {
// 	s.str = strings.TrimLeftFunc(s.str, f)
// 	return s.str
// }

// //返回将s后端所有cutset包含的utf-8码值都去掉的字符串
// func (s *Stringx) TrimRight(cutset string) string {
// 	s.str = strings.TrimRight(s.str, cutset)
// 	return s.str
// }

// //返回将s后端所有满足f的unicode码值都去掉的字符串
// func (s *Stringx) TrimRightFunc(f func(rune) bool) string {
// 	s.str = strings.TrimRightFunc(s.str, f)
// 	return s.str
// }

// //返回去除s可能的前缀prefix的字符串
// func (s *Stringx) TrimPrefix(prefix string) string {
// 	s.str = strings.TrimPrefix(s.str, prefix)
// 	return s.str
// }

// //返回去除s可能的后缀suffix的字符串
// func (s *Stringx) TrimSuffix(suffix string) string {
// 	s.str = strings.TrimSuffix(s.str, suffix)
// 	return s.str
// }

// //返回将字符串按照空白（unicode.IsSpace确定，可以是一到多个连续的空白字符）分割的多个字符串。
// //如果字符串全部是空白或者是空字符串的话，会返回空切片
// func (s *Stringx) Fields() []string {
// 	return strings.Fields(s.str)
// }

// //类似Fields，但使用函数f来确定分割符（满足f的unicode码值）。
// //如果字符串全部是分隔符或者是空字符串的话，会返回空切片。
// func (s *Stringx) FieldsFunc(f func(rune) bool) []string {
// 	return strings.FieldsFunc(s.str, f)
// }

// //用去掉s中出现的sep的方式进行分割，会分割到结尾，并返回生成的所有片段组成的切片
// //（每一个sep都会进行一次切割，即使两个sep相邻，也会进行两次切割）。
// //如果sep为空字符，Split会将s切分成每一个unicode码值一个字符串。
// func (s *Stringx) Split(sep string) []string {
// 	return strings.Split(s.str, sep)
// }

// //用去掉s中出现的sep的方式进行分割，会分割到结尾，并返回生成的所有片段组成的切片
// //（每一个sep都会进行一次切割，即使两个sep相邻，也会进行两次切割）。如果sep为空字符，
// //Split会将s切分成每一个unicode码值一个字符串。参数n决定返回的切片的数目：
// // n > 0 : 返回的切片最多n个子字符串；最后一个子字符串包含未进行切割的部分。
// // n == 0: 返回nil
// // n < 0 : 返回所有的子字符串组成的切片
// func (s *Stringx) SplitX(sep string, n int) []string {
// 	return strings.SplitN(s.str, sep, n)
// }

// //用从s中出现的sep后面切断的方式进行分割，会分割到结尾，并返回生成的所有片段组成的切片
// //（每一个sep都会进行一次切割，即使两个sep相邻，也会进行两次切割）。如果sep为空字符，
// //Split会将s切分成每一个unicode码值一个字符串
// func (s *Stringx) SplitAfter(sep string) []string {
// 	return strings.SplitAfter(s.str, sep)
// }

// //用从s中出现的sep后面切断的方式进行分割，会分割到结尾，并返回生成的所有片段组成的切片
// // （每一个sep都会进行一次切割，即使两个sep相邻，也会进行两次切割）。如果sep为空字符，
// // Split会将s切分成每一个unicode码值一个字符串。参数n决定返回的切片的数目：
// // n > 0 : 返回的切片最多n个子字符串；最后一个子字符串包含未进行切割的部分。
// // n == 0: 返回nil
// // n < 0 : 返回所有的子字符串组成的切片
// func (s *Stringx) SplitAfterN(sep string, n int) []string {
// 	return strings.SplitAfterN(s.str, sep, n)
// }
