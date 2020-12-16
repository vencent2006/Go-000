/**
 * @Author: vincent
 * @Description:
 * @File:  article
 * @Version: 1.0.0
 * @Date: 2020/12/16 17:21
 */

package data

import (
	"log"
	"time"
	"week04/internal/biz"
)

const (
	uid    = 8888
	author = "vincent"
)

func NewArticleRepo() biz.ArticleRepo {
	return &articleRepo{}
}

// articleRepo implementation
type articleRepo struct{}

func (repo *articleRepo) Post(article *biz.Article) int64 {
	// todo: get uid and author(uname) from context
	article.Uid = uid
	article.Author = author

	// todo: post article code
	log.Printf("post article: author(%d, %s),title(%s),content(%s)\n",
		article.Uid,
		article.Author,
		article.Title,
		article.Content)

	// todo: article id
	articleId := time.Now().Unix()

	return articleId
}
