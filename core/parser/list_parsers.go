package parser

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/ray1422/SML/core/container"
	"github.com/ray1422/SML/utils"
)

var ListLineRe = regexp.MustCompile(`^((?P<indent>[ \t]*)[ \t]*(?P<listType>-|[0-9]+\.) ?(?P<content>[^\n]+))`)

func readLine(s *string) (string, int) {
	s_idx := -1
	for i, r := range *s {
		if r != '\n' {
			s_idx = i
			break
		}
	}
	if s_idx == -1 {
		// all are \n
		l := len(*s)
		*s = ""
		return "", l
	}
	*s = (*s)[s_idx:]
	idx := strings.IndexRune(*s, '\n')
	if idx == -1 {
		// no \n found
		tmp := *s
		*s = ""
		return tmp, len(tmp) + s_idx
	} else {
		tmp := (*s)[:idx]
		*s = (*s)[idx:]
		return tmp, idx + s_idx
	}

}
func parseList(s string) (int, *container.ListBlock) {
	subs := utils.RegexNamedGroupMap(ListLineRe.FindStringSubmatch(s), ListLineRe.SubexpNames())
	if len(subs) == 0 {
		return 0, nil
	}
	ordered := subs["listType"] != "-"
	fmt.Println("parseList", subs["content"], len(subs["indent"]))
	blk, _ := childUtil([]string{subs["content"]}, 0, &container.ListBlock{Ordered: ordered})
	return len(subs["indent"]), blk.(*container.ListBlock)

}
func parseListGroup(ss []string) (*container.ListGroupBlock, int) {
	n_line_processed := 0
	if len(ss) == 0 {
		return nil, n_line_processed
	}
	if len(ss[len(ss)-1]) != 0 {
		ss = append(ss, "")
	}
	blk := &container.ListGroupBlock{}
	stk := utils.NewStack()
	stk.Push(blk)
	rootIndent := -1
	for i := 0; i < len(ss); i++ {
		line := ss[i]

		indent, val := parseList(line)
		var _cur interface{} = nil
		_ = _cur
		var child interface{} = nil
		if rootIndent == -1 {
			rootIndent = indent
		}
		if indent > rootIndent {
			stk.Push(&container.ListGroupBlock{})
		} else if indent < rootIndent {
			child = stk.Pop()
			if stk.Top() == nil {
				if v, ok := child.(*container.ListGroupBlock); ok && v != nil {
					return v, n_line_processed
				} else {
					return nil, n_line_processed
				}
			}

		}
		_cur = stk.Top()

		// if _cur == nil { // pop fail, could caused by bad indent
		// 	stk.Push(&container.ListGroupBlock{})
		// }
		cur := stk.Top().(*container.ListGroupBlock)
		if v, ok := child.(*container.ListGroupBlock); ok && v != nil {
			cur.Append(v)
			child = nil
		}
		if val != nil {
			cur.Append(val)
		}
		n_line_processed++
		rootIndent = indent
	}
	for stk.Top() != nil {
		c := stk.Pop().(*container.ListGroupBlock)
		if v, ok := stk.Top().(*container.ListGroupBlock); ok {
			v.Append(c)
		}
	}
	return blk, n_line_processed
}
func parseListGroups(s *string) *container.BaseBlock {
	*s = strings.ReplaceAll(*s, "\t", "    ")
	lines := strings.Split(*s, "\n")
	for _, line := range lines {
		fmt.Printf("'%s'\n", line)
	}

	wrapper2 := &container.BaseBlock{}
	for len(lines) > 0 {
		blks, n := parseListGroup(lines)
		wrapper2.Append(blks)
		lines = lines[n:]
	}

	return wrapper2
}

// just a helper
func registerLists() {
	// UL Group
	RegisterBlockParser(&RegexParser{
		/*
			match the whole list content
		*/
		re: regexp.MustCompile(`^\n?(?P<listContent>(?P<firstIndent>[ \t]*)(([ \t]*)(-|[0-9]+\.) ?([^\n]+)(\n))+)`),
		parse: func(re *regexp.Regexp, s string) (container.Block, int) {
			matches := re.FindStringSubmatch(s)
			names := re.SubexpNames()
			if matches == nil {
				return nil, 0
			}
			subs := utils.RegexNamedGroupMap(matches, names)
			content := subs["listContent"]
			return parseListGroups(&content), len(subs[0])
		},
	})
}
