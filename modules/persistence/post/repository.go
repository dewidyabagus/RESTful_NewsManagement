package post

import (
	"errors"
	"time"

	"github.com/lib/pq"
	"gorm.io/gorm"

	"RESTful/business"
	"RESTful/business/post"
	"RESTful/modules/persistence/topic"
)

type Topic topic.Topic

type Post struct {
	ID          string `gorm:"id;type:uuid;primaryKey"`
	TopicID     string `gorm:"topic_id;type:uuid"`
	Topic       Topic
	Title       string         `gorm:"title;type:varchar(100);not null"`
	Slug        string         `gorm:"slug;type:varchar(50);unique:posts_slug_uniq;not null"`
	Excerpt     string         `gorm:"excerpt;type:text;not null"`
	Body        string         `gorm:"body;type:text;not null"`
	Tags        pq.StringArray `gorm:"tags;type:varchar(30)[]"`
	Published   bool           `gorm:"published;type:boolean;default:false"`
	PublishedAt time.Time      `gorm:"published_at;type:timestamp;not null"`
	CreatedAt   time.Time      `gorm:"created_at;type:timestamp;not null"`
	UpdatedAt   time.Time      `gorm:"updated_at;type:timestamp;not null"`
	DeletedAt   time.Time      `gorm:"deleted_at;type:timestamp"`
}

func (p *Post) toBusinessPost() *post.Post {
	return &post.Post{
		ID:              p.ID,
		TopicId:         p.TopicID,
		TopicName:       p.Topic.Name,
		TopicDesciption: p.Topic.Description,
		Title:           p.Title,
		Slug:            p.Slug,
		Excerpt:         p.Excerpt,
		Body:            p.Body,
		Tags:            p.Tags,
		Published:       p.Published,
		PublishedAt:     p.PublishedAt,
		CreatedAt:       p.CreatedAt,
		UpdatedAt:       p.UpdatedAt,
	}
}

func toAllBusinessPost(posts *[]Post) *[]post.Post {
	var response = []post.Post{}

	for _, post := range *posts {
		response = append(response, *post.toBusinessPost())
	}

	return &response
}

func toInsertPost(p *post.Post) *Post {
	return &Post{
		ID:        p.ID,
		TopicID:   p.TopicId,
		Title:     p.Title,
		Slug:      p.Slug,
		Excerpt:   p.Excerpt,
		Body:      p.Body,
		Tags:      p.Tags,
		CreatedAt: p.CreatedAt,
		UpdatedAt: p.UpdatedAt,
	}
}

type Repository struct {
	DB *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db}
}

func (r *Repository) InsertPost(post *post.Post) error {
	return r.DB.Create(toInsertPost(post)).Error
}

func (r *Repository) FindPostBySlug(slug *string) (*post.Post, error) {
	var post = new(Post)

	err := r.DB.Preload("Topic").First(post, "slug = ? and (to_char(deleted_at, 'YYYY') = '0001' or deleted_at is null)", slug).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, business.ErrDataNotFound
		}
		return nil, err
	}

	return post.toBusinessPost(), nil
}

func (r *Repository) FindPostById(id *string) (*post.Post, error) {
	var result = new(Post)

	err := r.DB.First(result, "(to_char(deleted_at, 'YYYY') = '0001' or deleted_at is null) and id = ?", id).Error
	if err != nil {
		return nil, err
	}

	return result.toBusinessPost(), nil
}

func (r *Repository) FindPostByTopicId(topicId *string) (*[]post.Post, error) {
	var posts = new([]Post)

	err := r.DB.Preload("Topic").
		Where("(to_char(deleted_at, 'YYYY') = '0001' or deleted_at is null) and topic_id = ?", topicId).
		Order("created_at asc").Find(posts).Error
	if err != nil {
		return nil, err
	}

	return toAllBusinessPost(posts), nil
}

func (r *Repository) FindAllPost(status *string) (*[]post.Post, error) {
	var posts = new([]Post)

	rs := r.DB.Preload("Topic")
	if *status == "deleted" {
		rs.Where("to_char(deleted_at, 'YYYY') != '0001'")

	} else {
		rs.Where("(to_char(deleted_at, 'YYYY') = '0001' or deleted_at is null)")

		if *status == "publish" {
			rs.Where("published = true") // data has been published
		} else {
			rs.Where("published = false") // data draft
		}
	}

	if err := rs.Find(posts).Error; err != nil {
		return nil, err
	}

	return toAllBusinessPost(posts), nil
}

func (r *Repository) PublishPost(id *string, published time.Time) error {
	var result = new(Post)

	if err := r.DB.First(result, "id = ?", id).Error; err != nil {
		return err
	}

	return r.DB.Model(result).Updates(Post{Published: true, PublishedAt: published}).Error
}

func (r *Repository) UpdatePost(id *string, p *post.Post) error {
	var result = new(Post)

	if err := r.DB.First(result, "id = ?", id).Error; err != nil {
		return err
	}

	return r.DB.Model(result).Updates(Post{
		TopicID:   p.TopicId,
		Title:     p.Title,
		Slug:      p.Slug,
		Excerpt:   p.Excerpt,
		Body:      p.Body,
		Tags:      p.Tags,
		UpdatedAt: p.UpdatedAt,
	}).Error
}
