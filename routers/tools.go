package routers

import (
	"encoding/json"
	"fmt"
	"github.com/Tianyi822/easy_blog_server/logger"
	"github.com/Tianyi822/easy_blog_server/model/dto"
	"github.com/gin-gonic/gin"
)

func GetArticleDataFromRawData(ctx *gin.Context) (*dto.ArticleDto, error) {
	// 获取并解析 RawData 数据
	rowData, _ := ctx.GetRawData()
	articleDto := &dto.ArticleDto{}

	// 解析数据
	err := json.Unmarshal(rowData, articleDto)
	if err != nil {
		msg := fmt.Sprintf("请求解析失败，请检查数据格式是否正确: %v", err)
		logger.Error(msg)
		BadRequest(ctx, msg, nil)
		return nil, err
	}

	return articleDto, nil
}
