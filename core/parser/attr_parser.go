package parser

import (
	"io"
	"strings"
	"unicode"

	"github.com/ray1422/SML/utils"
)

func parseAttr(s string) (utils.Dict, int) {

	stk := utils.NewStack()
	keys := utils.NewStack()
	keys.Push(new(string))
	stk.Push(&utils.Dict{})
	parsingKey := true
	key := func() *string { return keys.Top().(*string) }
	cur := func() *utils.Dict { return stk.Top().(*utils.Dict) }
	var val interface{} = nil
	saveVal := func(trim bool) {
		if v, ok := val.(*string); ok {
			if trim {
				*v = strings.TrimSpace(*v)
			}
			(*cur())[*key()] = *v
			val = nil
		}
	}
	reader := strings.NewReader(s)
	r, i, err := reader.ReadRune()
	if err != nil || r != rune('{') {
		return nil, 0
	}
	for {
		r, size, err := reader.ReadRune()
		if err == io.EOF {
			break
		}
		i += size
		switch r {
		case rune('{'):
			if parsingKey { // must parsing v// al
				return nil, 0
			}
			el := &utils.Dict{}
			stk.Push(el)
			keys.Push(new(string))
			*key() = ""
			parsingKey = true
		case rune('}'):
			// save last element if not ended
			if !parsingKey || *key() != "" {
				// last val not saved
				saveVal(true)
				parsingKey = true
			}
			// pop keys
			keys.Pop()
			// save current element to parent
			el := stk.Pop().(*utils.Dict)
			if stk.Empty() {
				return *el, i
			}

			(*cur())[*key()] = *el
		case rune(':'):
			if !parsingKey {
				return nil, 0
			}
			// set parsineKey to false
			parsingKey = false
		case rune(','):
			// save last element
			saveVal(true)
			parsingKey = true
			*key() = ""
		case rune('"'), rune('\''):
			r0 := r
			if v, ok := val.(*string); !parsingKey && ok {
				if strings.TrimSpace(*v) != "" {
					*v += string(r)
					break
				}
			} else if parsingKey {
				if strings.TrimSpace(*key()) != "" {
					*key() += string(r)
					break
				}
			}
			buf := ""
			j := 0
			for {
				r2, size, err := reader.ReadRune()
				if err != nil {
					return nil, 0
				}
				j += size
				if r2 == r {
					buf = s[i : i+j-size]
					i += j
					break
				}
			}
			if parsingKey {
				*key() = buf
				buf2 := ""
			keyFor:
				for {
					r, size, err := reader.ReadRune()
					if err != nil {
						return nil, 0
					}
					switch {
					case r == rune(':'): // ?????????????????????
						reader.UnreadRune()
						break keyFor
					case unicode.IsSpace(r):
						buf2 += string(r)
						i += size
					default:
						return nil, 0
					}
				}

			} else {
				val = &buf
				saveVal(false)
				buf2 := ""
			ValFor:
				for {
					r, size, err := reader.ReadRune()
					if err != nil {
						return nil, 0
					}
					switch {
					case r == rune('}'), r == rune(','): // ?????????????????????
						reader.UnreadRune()
						break ValFor
					case unicode.IsSpace(r):
						buf2 += string(r)
						i += size
					default:
						i += size
						val = new(string)
						*val.(*string) += string(r0) + buf + string(r0) + buf2 + string(r)
					}
				}
			}
			// fmt.Printf("buf: '%s'\n", buf)
		case ('\r'):
			// TODO
			fallthrough
		case ('\n'):
			// TODO
		case (' '), ('\t'):
			if parsingKey {
				break
			}
			fallthrough
		default:
			if parsingKey {
				*key() += string(r)
			} else {
				if val == nil {
					val = new(string)
					*val.(*string) = ""
				}
				if v, ok := val.(*string); ok {
					*v += string(r)
				}
			}
		}
	}
	return nil, 0
}
