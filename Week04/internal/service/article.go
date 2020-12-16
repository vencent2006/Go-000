/**
 * @Author: vincent
 * @Description:
 * @File:  article
 * @Version: 1.0.0
 * @Date: 2020/12/16 18:30
 */

package service

import (
	"context"
	pb "week04/api/article/v1"
	"week04/internal/biz"
)

const (
	uid   = 8888
	uname = "vincent"
)

type ArticleService struct {
	u *biz.ArticleUseCase
	pb.UnimplementedArticleServer
}

func NewArticleService(u *biz.ArticleUseCase) *ArticleService {
	return &ArticleService{u: u}
}

func (s *ArticleService) PostArticle(ctx context.Context, r *pb.PostArticleRequest) (*pb.PostArticleResponse, error) {
	// DTO -> DO
	a := &biz.Article{Title: r.Title, Content: r.Content}

	// call biz
	s.u.SaveArticle(a)

	// response
	return &pb.PostArticleResponse{Id: a.ID}, nil
}
