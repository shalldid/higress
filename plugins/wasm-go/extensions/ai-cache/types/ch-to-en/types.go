package ch_to_en

import (
	"encoding/json"
	"strings"
)

type Stack struct {
	Row    uint8
	Groups []Group
}

type Group struct {
	Ch string
	En string
}

func (s *Stack) Translate(str string) string {
	for _, g := range s.Groups {
		if strings.Contains(str, g.Ch) {
			return strings.ReplaceAll(str, g.Ch, g.En)
		}
	}
	return str
}

func (s *Stack) Print() string {
	Str, _ := json.Marshal(s)
	return string(Str)
}
