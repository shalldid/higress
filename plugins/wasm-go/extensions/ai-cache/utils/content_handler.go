package utils

import (
	Antonym "ai-cache/types/antonym"
	ProperNoun "ai-cache/types/proper-noun"
	"strconv"
	"strings"
)

func TrimQuote(source string) string {
	TempKey := strings.Trim(source, `"`)
	Key, _ := zhToUnicode([]byte(TempKey))
	return string(Key)
}

func zhToUnicode(raw []byte) ([]byte, error) {
	str, err := strconv.Unquote(strings.Replace(strconv.Quote(string(raw)), `\\u`, `\u`, -1))
	if err != nil {
		return nil, err
	}
	return []byte(str), nil
}

func Generate(k1 string, k2 string, s *ProperNoun.Stack) string {
	for _, ProperNounGroup := range s.Groups {
		HitCount := 0
		HitFirstNoun := ""
		HitSecondNoun := ""
		for _, ProperNounGroupEle := range ProperNounGroup.Ele {
			if strings.Contains(k1, ProperNounGroupEle) {
				HitCount++
				HitFirstNoun = ProperNounGroupEle
				continue
			}
			if strings.Contains(k2, ProperNounGroupEle) {
				HitCount++
				HitSecondNoun = ProperNounGroupEle
				continue
			}
		}
		if HitCount >= 2 {
			return strings.ReplaceAll(k1, HitFirstNoun, HitSecondNoun)
		}
	}

	return k1 + k2

}

func IsDiffWithAntonymSet(a Antonym.Stack, s1, s2 string) bool {
	diff1, diff2 := getDiff(s1, s2)
	for _, AntonymGroup := range a.Groups {
		HitCount := 0
		for _, AntonymGroupEle := range AntonymGroup.Ele {
			if AntonymGroupEle == diff1 {
				HitCount++
			}
			if AntonymGroupEle == diff2 {
				HitCount++
			}
		}
		if HitCount >= 2 {
			return true
		}
	}
	return false
}

func getDiff(s1, s2 string) (string, string) {
	rs1 := []rune(s1)
	rs2 := []rune(s2)
	var diffS1 string
	var diffS2 string
	for i := 0; i < len(rs1)-1; i++ {
		if i >= len(rs2) || rs2[i] != rs1[i] {
			diffS1 = string(rs1[i:])
			diffS2 = string(rs2[i:])
			break
		}
	}
	rs3 := []rune(diffS1)
	rs4 := []rune(diffS2)
	var diffS3 string
	var diffS4 string
	for i := 1; i <= len(rs3); i++ {
		if len(rs4) < i || rs3[len(rs3)-i] != rs4[len(rs4)-i] {
			diffS3 = string(rs3[:len(rs3)-i+1])
			diffS4 = string(rs4[:len(rs4)-i+1])
			break
		}
	}
	return strings.Trim(diffS3, " "), strings.Trim(diffS4, " ")
}
