package response

import (
	"RESTful/business/topic"
)

func GetAllTopic(topics *[]topic.Topic) *[]TopicDetail {
	var response = []TopicDetail{}

	for _, item := range *topics {
		response = append(response, *GetOneTopic(&item))
	}

	if response == nil {
		response = []TopicDetail{}
	}

	return &response
}
