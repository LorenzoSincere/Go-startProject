package Repository

import (
	"errors"
	"gorm.io/gorm"
	"startProject/go-project-example/Util"
	"sync"
	"time"
)

type Post struct {
	Id        int64     `gorm:"column:id"`
	ParentId  int64     `gorm:"column:parent_id"`
	UserId    int64     `gorm:"column:user_id"`
	Content   string    `gorm:"column:content"`
	DiggCount int       `gorm:"column:digg_count"`
	CreatedAt time.Time `gorm:"column:created"`
}

func (Post) TableName() string {
	return "post"
}

type PostDao struct {
}

var postDao *PostDao
var postOnce sync.Once

func NewPostDaoInstance() *PostDao {
	postOnce.Do(
		func() {
			postDao = &PostDao{}
		})
	return postDao
}

func (*PostDao) QueryPostById(id int64) (*Post, error) {
	var post Post
	err := db.Where("id = ?", id).Find(&post).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if err != nil {
		Util.Logger.Error("find post by id err:" + err.Error())
		return nil, nil
	}
	return &post, nil
}

func (*PostDao) QueryPostByParentId(parentId int64) ([]*Post, error) {
	var posts []*Post
	err := db.Where("parent_id = ?", parentId).Find(&posts).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if err != nil {
		Util.Logger.Error("find post by parent id err:" + err.Error())
	}
	return posts, nil
}

func (*PostDao) CreatePost(post *Post) error {
	if err := db.Create(post).Error; err != nil {
		Util.Logger.Error("insert post err:" + err.Error())
		return err
	}
	return nil
}
