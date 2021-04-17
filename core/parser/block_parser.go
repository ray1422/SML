package parser

import (
	"regexp"

	"github.com/ray1422/SML/core/container"
	"github.com/ray1422/SML/utils"
)

var (
	BlockParsers []BlockParser = nil
	rem                        = regexEscapeMatchHelper
)

type BlockParser interface {
	Parse(s string) (container.Block, int)
}

// RegisterBlockParser Please make sure your identifier is recorded in specialParser (parser.go) or it might be treat as normal string.
func RegisterBlockParser(f BlockParser) {
	BlockParsers = append(BlockParsers, f)
}

type RegexParser struct {
	re    *regexp.Regexp
	parse func(re *regexp.Regexp, s string) (container.Block, int)
}

func (rp *RegexParser) Parse(s string) (container.Block, int) {
	return rp.parse(rp.re, s)
}
func RegAll() {
	if BlockParsers != nil {
		return
	}
	BlockParsers = []BlockParser{}

	// drunk coding, might be wrong
	RegisterBlockParser(&RegexParser{
		re: regexp.MustCompile(`^\x60\x60\x60(.*?)\n(([^\x60\\]|\\\x60)+)\x60\x60\x60`), // code block
		parse: func(re *regexp.Regexp, s string) (container.Block, int) {
			strs := re.FindStringSubmatch(s)
			if len(strs) < 3 {
				return nil, 0
			}
			lang := strs[1]
			content := strs[2]
			return &container.CodeBlock{Lang: lang, Content: content, Inline: true}, len(strs[0])
		},
	})
	RegisterBlockParser(&RegexParser{
		re: regexp.MustCompile(`^\x60(([^\\\n]|\\.)+?)\x60(\{[:]?[.]?(.*?)\})?`), // inline code block
		parse: func(re *regexp.Regexp, s string) (container.Block, int) {
			strs := re.FindStringSubmatch(s)
			if len(strs) < 4 {
				return nil, 0
			}
			lang := strs[4]
			content := strs[1]
			return &container.CodeBlock{Lang: lang, Content: content, Inline: true}, len(strs[0])
		},
	})

	RegisterBlockParser(&RegexParser{
		re: regexp.MustCompile(`^\n?#{1,6}([ \t]*)(.+)(\n|$)`), // heading 1-6
		parse: func(re *regexp.Regexp, s string) (container.Block, int) {
			return childUtil(re.FindStringSubmatch(s), 2, &container.HeadingBlock{})
		},
	})

	registerLists()
	RegisterBlockParser(&RegexParser{ // img
		re: regexp.MustCompile(`^\!(\[(?P<alt>.*)\])\((?P<path>([^'"()\\]|\\.)*?|('([^'\\]|\\.)*?')|("([^"\\]|\\.)*?")) +(?P<title>([^'"()\\]*?)|('([^']|\\.)*?')|("([^"]|\\.)*?"))?\)`),
		parse: func(re *regexp.Regexp, s string) (container.Block, int) {
			subs := utils.RegexNamedGroupMap(re.FindStringSubmatch(s), re.SubexpNames())
			if subs == nil {
				return nil, 0
			}
			subs["title"] = utils.QuoteRM(subs["title"])
			subs["alt"] = utils.QuoteRM(subs["alt"])
			subs["path"] = utils.QuoteRM(subs["path"])
			attr, n := parseAttr(s[len(subs[0]):])
			el, n0 := childUtil([]string{subs[0], subs["title"]}, 1, &container.ImageBlock{
				Src:   subs["path"],
				Alt:   subs["alt"],
				Title: subs["title"],
				Attr:  attr,
			})
			return el, n + n0
		},
	})
	RegisterBlockParser(&RegexParser{
		re: regexp.MustCompile(`^\*\*` + rem(`\*`) + `\*\*`), // bold
		parse: func(re *regexp.Regexp, s string) (container.Block, int) {
			return childUtil(re.FindStringSubmatch(s), 1, &container.InlineBlock{InlineBlockType: container.BOLD})
		},
	})

	RegisterBlockParser(&RegexParser{
		re: regexp.MustCompile(`^__` + rem("_") + `__`), // italic
		parse: func(re *regexp.Regexp, s string) (container.Block, int) {
			return childUtil(re.FindStringSubmatch(s), 1, &container.InlineBlock{InlineBlockType: container.ITALIC})
		},
	})

	RegisterBlockParser(&RegexParser{
		re: regexp.MustCompile(`^~~` + rem("~") + `~~`), // strike
		parse: func(re *regexp.Regexp, s string) (container.Block, int) {
			return childUtil(re.FindStringSubmatch(s), 1, &container.InlineBlock{InlineBlockType: container.STRIKE})
		},
	})

}

func childUtil(c_strings []string, idx int, blk container.VarBlock) (container.Block, int) {
	if len(c_strings) < idx+1 {
		return nil, 0
	}
	content := c_strings[idx]
	l := len(c_strings[0])
	blk.Append(Parse(content))
	return blk, l

}

// regexEscapeMatchHelper ignore `s` but include `\s`
func regexEscapeMatchHelper(s string) string {
	s = `(([^\\` + s + `]|\\.)+?)` // ignore <s> but include \<s> (with `\` escape)
	return s
}
