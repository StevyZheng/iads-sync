package base

import (
	"fmt"
	"regexp"
	"sort"
	"strings"
)

func Trim(srcStr string, trimStr string) string {
	regStrTmp := fmt.Sprintf("^%[1]s*|%[1]s*$", trimStr)
	re := regexp.MustCompile(regStrTmp)
	ret := re.ReplaceAllString(srcStr, "")
	return ret
}

func SearchString(srcStr string, regStr string) []string {
	regStr1 := fmt.Sprintf("(?m:%s)", regStr)
	re := regexp.MustCompile(regStr1)
	return re.FindAllString(srcStr, -1)
}

func SplitString(srcStr string, splitStr string) []string {
	ret := strings.Split(srcStr, splitStr)
	return ret
}

func SearchSplitString(srcStr string, regStr string, splitStr string) [][]string {
	re := SearchString(srcStr, regStr)
	var ret [][]string
	for _, v := range re {
		v_re := SplitString(v, splitStr)
		ret = append(ret, v_re)
	}
	return ret
}

func SearchSplitStringColumn(srcStr string, regStr string, splitStr string, col int) []string {
	tmp := SearchSplitString(srcStr, regStr, splitStr)
	var ret []string
	for _, v := range tmp {
		ret = append(ret, Trim(v[col-1], " "))
	}
	return ret
}

func SearchStringFirst(srcStr string, regStr string) string {
	regStr1 := fmt.Sprintf("(?m:%s)", regStr)
	re := regexp.MustCompile(regStr1)
	findStr := re.FindAllString(srcStr, -1)
	if findStr != nil {
		return findStr[0]
	} else {
		return "nil"
	}
}

func SearchSplitStringFirst(srcStr string, regStr string, splitStr string) []string {
	re := SearchStringFirst(srcStr, regStr)
	if re == "nil" {
		return nil
	}
	var ret []string
	ret = SplitString(re, splitStr)
	return ret
}

func SearchSplitStringColumnFirst(srcStr string, regStr string, splitStr string, col int) string {
	tmp := SearchSplitStringFirst(srcStr, regStr, splitStr)
	if tmp == nil {
		return "nil"
	}
	return Trim(tmp[col-1], " ")
}

func UniqStringList(strList []string) []string {
	newArr := make([]string, 0)
	sort.Strings(strList)
	for i := 0; i < len(strList); i++ {
		repeat := false
		for j := i + 1; j < len(strList); j++ {
			if strList[i] == strList[j] {
				repeat = true
				break
			}
		}
		if !repeat {
			newArr = append(newArr, strList[i])
		}
	}
	return newArr
}
