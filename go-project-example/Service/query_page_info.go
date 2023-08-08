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
