package request

import "RESTful/business/post"

type Post struct {
	TopicId string   `json:"topic_id"`
	Title   string   `json:"title"`
	Slug    string   `json:"slug"`
	Excerpt string   `json:"excerpt"`
	Body    string   `json:"body"`
	Tags    []string `json:"tags"`
}

func (p *Post) ToBusinessPostSpec() *post.PostSpec {
	return &post.PostSpec{
		TopicId: p.TopicId,
		Title:   p.Title,
		Slug:    p.Slug,
		Excerpt: p.Excerpt,
		Body:    p.Body,
		Tags:    p.Tags,
	}
}
