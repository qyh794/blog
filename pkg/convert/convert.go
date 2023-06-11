package convert

import (
	"blog/models"
	"encoding/json"
)

func ConvertToPostDetailList(cachedData []string) []*models.PostDetail {
	postDetails := make([]*models.PostDetail, len(cachedData))
	for i, data := range cachedData {
		postDetail := &models.PostDetail{}
		err := json.Unmarshal([]byte(data), postDetail)
		if err != nil {
			// 处理解析错误
			continue
		}
		postDetails[i] = postDetail
	}
	return postDetails
}
