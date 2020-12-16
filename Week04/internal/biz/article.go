/**
 * @Author: vincent
 * @Description:
 * @File:  article
 * @Version: 1.0.0
 * @Date: 2020/12/16 17:01
 */

package biz

type Article struct {
	ID      int64
	Title   string
	Uid     int64
	Author  string
	Content string
}

// article repo
type ArticleRepo interface {
	Post(*Article) int64 // post article
}

// Article UserCase
type ArticleUseCase struct {
	repo ArticleRepo
}

func NewArticleUseCase(repo ArticleRepo) *ArticleUseCase {
	return &ArticleUseCase{repo: repo}
}

func (u *ArticleUseCase) SaveArticle(a *Article) {
	id := u.repo.Post(a)
	a.ID = id
}
