package parser

import (
	"fmt"
	"testing"

	"github.com/ray1422/SML/utils"
	"github.com/stretchr/testify/assert"
)

// all those test should fail to parse as attr
func TestParseAttrNil(t *testing.T) {
	t.Parallel()
	var s string = ""
	s = `{test: a, test: b`
	var a utils.Dict
	var n int
	a, n = parseAttr(s)
	assert.Nil(t, a)
	assert.Equal(t, 0, n)

	s = `{test: a, test: "b`
	a, n = parseAttr(s)
	assert.Nil(t, a)
	assert.Equal(t, 0, n)

	s = `{test: a, test: "b,`
	a, n = parseAttr(s)
	assert.Nil(t, a)
	assert.Equal(t, 0, n)

	s = `{test: a, test: "}`
	a, n = parseAttr(s)
	assert.Nil(t, a)
	assert.Equal(t, 0, n)

	s = `{test: a, test: "b
	
	
	
	
	`
	a, n = parseAttr(s)
	assert.Nil(t, a)
	assert.Equal(t, 0, n)

	s = `{test: a, test: "b}`
	a, n = parseAttr(s)
	assert.Nil(t, a)
	assert.Equal(t, 0, n)
	s = `test: a, test: "b}`
	a, n = parseAttr(s)
	assert.Nil(t, a)
	assert.Equal(t, 0, n)
	s = `{test: a, test: "b",:www:www,}`
	a, n = parseAttr(s)
	assert.Nil(t, a)
	assert.Equal(t, 0, n)

}

func TestParseAttrSimple(t *testing.T) {
	t.Parallel()
	var s string = ""
	s = `{testA: a, testB: b}`
	var a utils.Dict
	var n int
	a, n = parseAttr(s)
	if assert.NotNil(t, a) {
		assert.Equal(t, "a", a["testA"])
		assert.Equal(t, "b", a["testB"])
		assert.Equal(t, len(s), n)
	}
}

func TestParseAttrNested(t *testing.T) {
	t.Parallel()
	var s string = ""
	s = `{
	testA: a,
	testB: {
		I: 1,
		II: 2,},
	testC: {
		I: 1,
		II: 2
	},
	testD: {
		I: 1,
		II: 2},
	testE: {
		I: 1,
		II: 2,
	}
}
`
	var a utils.Dict
	var n int
	a, n = parseAttr(s)
	if assert.NotNil(t, a) {
		assert.Equal(t, "a", a["testA"])
		assert.Equal(t, "1", a.V("testB", "I"))
		assert.Equal(t, "2", a.V("testB", "II"))
		assert.Equal(t, "1", a.V("testC", "I"))
		assert.Equal(t, "2", a.V("testC", "II"))
		assert.Equal(t, "1", a.V("testD", "I"))
		assert.Equal(t, "2", a.V("testD", "II"))
		assert.Equal(t, "1", a.V("testE", "I"))
		assert.Equal(t, "2", a.V("testE", "II"))
		assert.Equal(t, len(s)-1, n)
	}
}

func TestParseAttrQuote(t *testing.T) {
	t.Parallel()
	var s string = ""
	s = `{"testA": a, a"test": b, a"test"b: c, "  test": d, t1: "test"t,
	t2: t "test" t}w`
	var a utils.Dict
	var n int
	a, n = parseAttr(s)
	if assert.NotNil(t, a) {
		assert.Equal(t, "a", a.V("testA"))
		assert.Equal(t, "b", a.V(`a"test"`))
		assert.Equal(t, "c", a.V(`a"test"b`))
		assert.Equal(t, "d", a.V("  test"))
		assert.Equal(t, `"test"t`, a.V("t1"))
		assert.Equal(t, `t "test" t`, a.V("t2"))
		assert.Equal(t, "w", s[n:])
	}
}

func TestParseAttrUnicode(t *testing.T) {
	s := `{"中文1": 中文啊啊啊啊, 中文！"喔喔喔": 恩？, 中文: 中文, "  中文": 中文2, }其他中文`
	a, n := parseAttr(s)
	if assert.NotNil(t, a) {
		assert.Equal(t, "中文啊啊啊啊", a.V("中文1"))
		assert.Equal(t, "恩？", a.V(`中文！"喔喔喔"`))
		assert.Equal(t, "中文", a.V(`中文`))
		assert.Equal(t, "中文2", a.V("  中文"))
		assert.Equal(t, "其他中文", s[n:])
	}
	a.Dump(0)
}

func TestParseTMP(t *testing.T) {
	t.Parallel()
	var s string = ""
	s = `{  "testA" 	: a "w" www}w`
	var a utils.Dict
	var n int
	a, n = parseAttr(s)
	a.Dump(0)
	fmt.Println(a == nil, n)
}
