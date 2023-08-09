package Service

import (
	"errors"
	"fmt"
	"startProject/go-project-example/Repository"
	"sync"
)

type TopicInfo struct {
	Topic *Repository.Topic
	User  *Repository.User
}

type PostInfo struct {
	Post *Repository.Post
	User *Repository.User
}

type PageInfo struct {
	TopicInfo *TopicInfo
	PostList  []*PostInfo
}

type QueryPageInfoFlow struct {
	topicId  int64
	pageInfo *PageInfo
	topic    *Repository.Topic
	posts    []*Repository.Post
	userMap  map[int64]*Repository.User
}

func QueryPageInfo(topicId int64) (*PageInfo, error) {
	return NewQueryPageInfoFlow(topicId).Do()
}

func NewQueryPageInfoFlow(topId int64) *QueryPageInfoFlow {
	return &QueryPageInfoFlow{
		topicId: topId,
	}
}

func (f *QueryPageInfoFlow) Do() (*PageInfo, error) {
	if err := f.checkParams(); err != nil {
		return nil, err
	}
	return f.pageInfo, nil
}

func (f *QueryPageInfoFlow) checkParams() error {
	if f.topicId <= 0 {
		return errors.New("topic id must be larger than 0")
	}
	return nil
}

func (f *QueryPageInfoFlow) prepareInfo() error {
	var wg sync.WaitGroup
	wg.Add(2)

	var topicErr, postErr error
	go func() {
		defer wg.Done()
		topic, err := Repository.NewTopicDaoInstance().QueryTopicById(f.topicId)
		if err != nil {
			topicErr = err
			return
		}
		f.topic = topic
	}()
	go func() {
		defer wg.Done()
		posts, err := Repository.NewPostDaoInstance().QueryPostByParentId(f.topicId)
		if err != nil {
			postErr = err
			return
		}
		f.posts = posts
	}()
	wg.Wait()

	if topicErr != nil {
		return topicErr
	}
	if postErr != nil {
		return postErr
	}

	//获取用户信息
	uids := []int64{f.topicId}
	for _, post := range f.posts {
		uids = append(uids, post.Id)
	}
	userMap, err := Repository.NewUserInstance().MQueryUserById(uids)
	if err != nil {
		return err
	}
	f.userMap = userMap
	return nil
}

func (f *QueryPageInfoFlow) packPageInfo() error {
	// topic 信息
	userMap := f.userMap
	topicUser, ok := userMap[f.topic.UserId]
	if !ok {
		return errors.New("has no topic user info")
	}
	postList := make([]*PostInfo, 0)
	for _, post := range f.posts {
		postUser, ok := userMap[post.UserId]
		if !ok {
			return errors.New("has no post user info for:" + fmt.Sprint(post.UserId))
		}
		postList = append(postList, &PostInfo{
			Post: post,
			User: postUser,
		})
	}

	f.pageInfo = &PageInfo{
		TopicInfo: &TopicInfo{
			Topic: f.topic,
			User:  topicUser,
		},
		PostList: postList,
	}
	return nil
}
