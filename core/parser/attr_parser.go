package parser

import (
	"strings"

	"github.com/ray1422/SML/utils"
)

func parseAttr(_s string) (*utils.Dict, int) {
	s := []rune(_s)
	if len(s) < 2 || s[0] != '{' {
		return nil, 0
	}
	strOrPtr := func(u interface{}, strip bool) interface{} {
		if v, ok := u.(*string); ok {
			if strip {
				return strings.TrimSpace(*v) // dereference *string
			}
			return *v
		} else {
			return u // keep void*
		}
	}
	stk := utils.NewStack()
	var cur *utils.Dict = nil
	parsingKey := false
	key := ""

	var val interface{}

	setParsingKey := func(b bool) {
		if b {
			key = ""
			parsingKey = true
		} else {
			val = nil
			parsingKey = false
		}
	}
	n := 0
mainFor:
	for n = 0; n < len(s); n++ {
		// _, _r := range s
		_r := s[n]
		r := rune(_r) // for debug LOL
		str_r := string(r)
		_ = str_r

		switch _r {
		case '{':
			if parsingKey { // fail // can't using dict as key for simplification
				return nil, 0
			}
			setParsingKey(true)
			stk.Push(&utils.Dict{})
		case '}':
			if stk.Empty() {
				// err
				return nil, 0
			}
			if !parsingKey { // 結算上一次的 val 但是 用戶可能已經輸入逗號，那就已經除存了
				(*stk.Top().(*utils.Dict))[key] = strOrPtr(val, true)
			}
			val = stk.Pop()

			if stk.Empty() { // 全部解析完了，剩下的還給後面的 parser
				cur = val.(*utils.Dict)
				break mainFor
			}
			(*stk.Top().(*utils.Dict))[key] = val
			val = nil
			setParsingKey(true)
		case '"', '\'':
			if parsingKey {
				if len(strings.TrimSpace(key)) > 0 {
					key += string(_r)
					break
				}
			} else {
				if v, ok := val.(*string); ok && len(strings.TrimSpace(*v)) > 0 {
					*v += string(_r)
					break
				}
			}
			tok := _r
			buf := ""
			for n++; n < len(s) && s[n] != tok; n++ {
				if s[n] == '\\' && n+1 < len(s) {

					buf += utils.DeEsc(string(s[n+1]))
					n++
					continue
				}
				buf += string(s[n])
			}

			if n >= len(s) || s[n] != '"' {
				return nil, 0
			}
			if parsingKey {
				key = buf
			} else {
				(*stk.Top().(*utils.Dict))[key] = buf
				setParsingKey(true)
			}
		case ':':
			setParsingKey(false)
		case '\\':
			// TODO escape
		case '\n':
			// TODO
		case ',':
			if parsingKey {
				continue
			}
			(*stk.Top().(*utils.Dict))[key] = strOrPtr(val, true)

			val = nil
			setParsingKey(true)
		case ' ', '\t':
			if !parsingKey {
				if v, ok := val.(*string); ok {
					*v += string(r)
				}
			}
		default:
			if parsingKey {
				key += string(r)
			} else {
				if val == nil {
					tmp := ""
					val = &tmp
				}
				if v, ok := val.(*string); ok {
					*v += string(rune(r))
				}
			}
		}
	}
	return cur, n + 1
}
