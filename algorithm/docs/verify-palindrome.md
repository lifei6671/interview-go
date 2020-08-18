# 验证回文串

## 01、题目示例

> 见微知著，发现一组数据很有趣，分享给大家。leetcode 第一题通过次数为 993,335，第二题通过次数为 396,160，第三题通过次数为 69,508。我想说什么，请自己悟。

### 第125题：验证回文串

给定一个字符串，验证它是否是回文串，只考虑字母和数字字符，可以忽略字母的大小写。

说明： 本题中，我们将空字符串定义为有效的回文串。

> 示例 1:
>输入: "A man, a plan, a canal: Panama"
>输出: true

>示例 2:
>输入: "race a car"
>输出: false

## 02、图解教程

经典题目，你需要像掌握反转字符串一样掌握本题。

首先，我想确保你知道什么是回文串。“回文串”是一个正读和反读都一样的字符串，比如“level”或者“noon”等等就是回文串。

对于字符串中可能存在的其他字符，可以通过正则替换，但是正则替换会增加程序运行复杂度，下面给出的是在判断过程中忽略其他字符：

```go
func isPalindrome(s string) bool {
	if s == "" {
		return false
	}
	s = strings.ToLower(s)
	if len(s) == 2  {
		return s[0] == s[1]
	}
	left := 0
	right := len(s) - 1
	for left < right {
        //忽略除字母和数字之外的字符
		if !((s[left] >= 'a' && s[left] <= 'z') || (s[left] >= '0' && s[left] <= '9')) {
			left++
			continue
		}
		if !((s[right] >= 'a' && s[right] <= 'z') || (s[right] >= '0' && s[right] <= '9')){
			right--
			continue
		}
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}
```