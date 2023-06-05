package apiv1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type UploadReq struct {
	g.Meta `type:"file" path:"/upload" tags:"Upload" method:"post" summary:"Upload File"`
}
type UploadRes struct {
	g.Meta    `example:"string"`
	TableName string
}
