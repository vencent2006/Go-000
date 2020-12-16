//+build wireinject

package main

import (
	"week04/internal/biz"
	"week04/internal/data"

	"github.com/google/wire"
)

func InitArticleUseCase() *biz.ArticleUseCase {
	wire.Build(biz.NewArticleUseCase, data.NewArticleRepo)
	return &biz.ArticleUseCase{}
}
