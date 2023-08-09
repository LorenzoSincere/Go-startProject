package Service

import (
	"errors"
	"startProject/go-project-example/Repository"
	"time"
	"unicode/utf8"
)

type PublishPostFlow struct {
	userId  int64
	content string
	topicId int64
	postId  int64
}

func PublishPost(topicId, userId int64, content string) (int64, error) {
	return NewPublishPostFlow(topicId, userId, content).Do()
}

func NewPublishPostFlow(topicId, userId int64, content string) *PublishPostFlow {
	return &PublishPostFlow{
		userId:  userId,
		content: content,
		topicId: topicId,
	}
}

func (f *PublishPostFlow) Do() (int64, error) {
	if err := f.checkParam(); err != nil {
		return 0, err
	}
	return f.postId, nil
}

func (f *PublishPostFlow) checkParam() error {
	if f.userId <= 0 {
		return errors.New("userId id must be larger than zero")
	}
	if utf8.RuneCountInString(f.content) >= 500 {
		return errors.New("content length must be less than 500")
	}
	return nil
}

func (f *PublishPostFlow) publish() error {
	post := &Repository.Post{
		ParentId:  f.topicId,
		UserId:    f.userId,
		Content:   f.content,
		CreatedAt: time.Now().UTC(),
	}
	if err := Repository.NewPostDaoInstance().CreatePost(post); err != nil {
		return err
	}
	f.postId = post.Id
	return nil
}
