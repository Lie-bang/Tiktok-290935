package Service

import (
	"context"
	"douyin/cmd/video/dal/db"
	"douyin/kitex_gen/douyinvideo"
	"log"
)

type UpdateCommentCountService struct {
	ctx context.Context
}

func NewUpdateCommentCountService(ctx context.Context) *UpdateCommentCountService {
	return &UpdateCommentCountService{ctx: ctx}
}

func (f *UpdateCommentCountService) UpdateCommentCount(req *douyinvideo.DouyinCommentCountRequest) error {

	err := db.UpdateCommentCount(f.ctx, req)
	if err != nil {
		log.Print(err)
		return err
	}
	return nil

}
