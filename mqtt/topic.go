package mqtt

import (
	"bytes"
	"errors"
	"fmt"
)

func ValidTopic(topic []byte) error {
	//no + #
	if bytes.ContainsAny(topic, "+#") {
		return fmt.Errorf("+ # is not valid (%s)", string(topic))
	}
	//no //
	if bytes.Contains(topic, []byte("//")) {
		return fmt.Errorf("// is not valid (%s)", string(topic))
	}
	return nil
}

func ValidSubscribe(topic []byte) error {
	if len(topic) == 0 {
		return errors.New("blank topic")
	}
	topics := bytes.Split(topic, []byte("/"))
	if len(topics[0]) == 0 {
		topics[0] = []byte("/")
	}
	for i, tt := range topics {
		l := len(tt)
		if l == 0 {
			return errors.New("inner blank")
		}
		if l == 1 && tt[0] == '#' && i < len(topics)-1 {
			return errors.New("# must be the last one")
		}
		if l > 1 && bytes.ContainsAny(tt, "+#") {
			return errors.New("+ # is alone")
		}
	}
	return nil
}

func MatchSubscribe(topic []byte, sub []byte) bool {
	i, j := 0, 0
	for i < len(topic) && j < len(sub) {
		t, s := topic[i], sub[j]
		if s == '#' {
			return true
		} else if s == '+' {
			for t != '/' {
				i++
				t = topic[i]
			}
			j++ // skip /
		} else if s != t {
			break
		}
		// else s==t
		i++
		j++
	}

	//Just match
	if i == len(topic) && j == len(sub) {
		return true
	}

	return false
}
