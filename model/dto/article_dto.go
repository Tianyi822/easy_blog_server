package dto

type ArticleDto struct {
	Title      string `json:"title"`
	BriefIntro string `json:"brief_intro"`
	Url        string `json:"url"`
}
