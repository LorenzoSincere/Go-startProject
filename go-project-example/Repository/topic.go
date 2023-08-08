package Repository

import (
	"startProject/go-project-example/Util"
	"sync"
	"time"
)

type Topic struct {
	Id        int64     `gorm:"column:id"`
	UserId    int64     `gorm:"column:user_id"`
	Title     string    `gorm:"column:title"`
	Content   string    `gorm:"column:content"`
	CreatedAt time.Time `gorm:"column:created_at"`
}

func (Topic) TableName() string {
	return "topic"
}

type TopicDao struct {
}

var topicDao *TopicDao
var topicOnce sync.Once

func NewTopicDaoInstance() *TopicDao {
	topicOnce.Do(
		func() {
			topicDao = &TopicDao{}
		})
	return topicDao
}

func (*TopicDao) QueryTopicById(id int64) (*Topic, error) {
	var topic Topic
	err := db.Where("id = ?", id).Find(&topic).Error
	if err != nil {
		Util.Logger.Error("find topic by id err:" + err.Error())
	}
	return &topic, err
}
