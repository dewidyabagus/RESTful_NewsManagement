package response

import "RESTful/business/post"

func GetAllPostSummary(posts *[]post.Post) *[]PostSummary {
	var response = []PostSummary{}

	for _, post := range *posts {
		response = append(response, *GetOnePostSummary(&post))
	}

	if response == nil {
		response = []PostSummary{}
	}

	return &response
}
