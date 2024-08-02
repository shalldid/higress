package proper_noun

import "encoding/json"

type Stack struct {
	Row    uint8
	Groups []Group
}

type Group struct {
	Ele []string
}

func (s *Stack) Print() string {
	Str, _ := json.Marshal(s)
	return string(Str)
}
