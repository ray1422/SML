package container

type ListBlock struct {
	Ordered bool
	BaseBlock
}

type ListGroupBlock struct {
	BaseBlock
}
type HeadingBlock struct {
	InlineBlock
}
type Frame struct {
	Direction directionT
	Ratio     []float64
	VarBlock
}
