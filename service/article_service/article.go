package article_service

import "oldboymiaosha/models"

type Article struct {
	ID            int
	TagID         int
	Title         string
	Desc          string
	Content       string
	CoverImageUrl string
	State         int
	CreateBy      string
	ModifiedBy    string
	PageNum       int
	PageSize      int
}

func (a *Article) Add() error {
	article := map[string]interface{}{
		"tag_id":          a.TagID,
		"title":           a.Title,
		"desc":            a.Desc,
		"content":         a.Content,
		"created_by":      a.CreateBy,
		"cover_image_url": a.CoverImageUrl,
		"state":           a.State,
	}
	if err := models.AddArticle(article); err != nil {
		return err
	}

	return nil

}
