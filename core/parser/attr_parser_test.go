package parser

import (
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
