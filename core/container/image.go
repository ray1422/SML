package container

import "github.com/ray1422/SML/utils"

type ImageBlock struct {
	BaseBlock
	Title string
	Src   string
	Alt   string
	Attr  utils.Dict
}
