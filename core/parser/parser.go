package parser

import (
	"math"
	"regexp"
	"sync"

	"github.com/ray1422/SML/core/container"
)

var (
	specialParser = regexp.MustCompile(`^(([^#\x60*_~\\$\-.0-9\n{}!\[\]()])+)`)
	Parse         = ParseNonAsync
)

type Parser struct {
	doc string
}

// SetDoc set doc
func (p *Parser) SetDoc(doc string) error {
	p.doc = doc
	return nil
}
func ParseAsync(s string) (b container.VarBlock) {
	root := &container.BaseBlock{}
	str := ""
	for s != "" {
		flag := true
		type ret struct {
			order int
			val   container.Block
			idx   int
		}
		var wg sync.WaitGroup
		ch := make(chan ret, len(BlockParsers)+5)
		for i, bp := range BlockParsers {
			wg.Add(1)
			go func(i int, bp BlockParser) {
				defer wg.Done()
				if val, idx := bp.Parse(s); val != nil {
					if str != "" {
						// clean up last str
						root.Append(&container.TextBlock{Content: str})
						str = ""
					}
					ch <- ret{order: i, val: val, idx: idx}
					// root.Append(val)
					// s = s[idx:]
					flag = false
					// break
				}
			}(i, bp)
		}
		wg.Wait()
		l := 0
		r := ret{order: int(math.Inf(1)), val: nil, idx: -1}
		for len(ch) > 0 {
			item := <-ch
			if item.order > r.order {
				r = item
			}
			l++
		}
		flag = l == 0
		if flag {
			strs := specialParser.FindStringSubmatch(s)
			if len(strs) > 2 && len(strs[1]) > 0 {
				str += strs[1]
				s = s[len(strs[1]):]
			} else {
				str += string(s[0])
				s = s[1:]
			}
		} else {
			root.Append(r.val)
			s = s[r.idx:]
		}
	}
	if str != "" {
		// clean up last str
		root.Append(&container.TextBlock{Content: str})
		str = ""
	}
	return root
}

func ParseNonAsync(s string) (b container.VarBlock) {
	root := &container.BaseBlock{}
	str := ""
	for s != "" {
		flag := true
		for _, bp := range BlockParsers {
			if val, idx := bp.Parse(s); val != nil {
				if str != "" {
					// clean up last str
					root.Append(&container.TextBlock{Content: str})
					str = ""
				}
				root.Append(val)
				s = s[idx:]
				flag = false
				break
			}
		}
		if flag {
			strs := specialParser.FindStringSubmatch(s)
			if len(strs) > 2 && len(strs[1]) > 0 {
				str += strs[1]
				s = s[len(strs[1]):]
			} else {
				str += string(s[0])
				s = s[1:]
			}
		}
	}
	if str != "" {
		// clean up last str
		root.Append(&container.TextBlock{Content: str})
		str = ""
	}
	return root
}
