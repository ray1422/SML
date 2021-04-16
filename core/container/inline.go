package container

const (
	BOLD      inlineBlockT = iota
	ITALIC    inlineBlockT = iota
	STRIKE    inlineBlockT = iota
	UNDERLINE inlineBlockT = iota
)

var (
	inlineBlockT2Str = map[inlineBlockT]string{
		BOLD:      "BOLD",
		ITALIC:    "ITALIC",
		STRIKE:    "STRIKE",
		UNDERLINE: "UNDERLINE",
	}
)

type InlineBlock struct {
	BaseBlock
	InlineBlockType inlineBlockT
}
type TextBlock struct {
	RenderGeneralAttr
	Content string
}

// Clear clear all children
func (blk *TextBlock) Children() []Block {
	return []Block{blk}
}
