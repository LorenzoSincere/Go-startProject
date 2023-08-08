package Service

import "startProject/go-project-example/Repository"

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
	return NewQueryPageInfoFlow(topicId).Do(), nil
}

func NewQueryPageInfoFlow(topId int64) *QueryPageInfoFlow {
	return &QueryPageInfoFlow{
		topicId: topId,
	}
}
