package controller

import (
	"context"

	"problemanalysis-backend/apiv1"

	"github.com/gogf/gf/v2/frame/g"
)

var (
	Upload = cUpload{}
)

type cUpload struct{}

func (h *cHello) Upload(ctx context.Context, req *apiv1.UploadReq) (res *apiv1.UploadRes, err error) {
	file := g.RequestFromCtx(ctx).GetUploadFile("upload-file")
	tableName, err := file.Save("/tmp/")
	if err != nil {
		return nil, err
	}
	res = &apiv1.UploadRes{
		TableName: tableName,
	}
	return
}
