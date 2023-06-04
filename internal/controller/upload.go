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
	files := g.RequestFromCtx(ctx).GetUploadFiles("upload-file")
	names, err := files.Save("/tmp/")
	if err != nil {
		return nil, err
	}
	g.RequestFromCtx(ctx).Response.WriteJsonExit(names)

	return
}
