package dto

type ArticleDto struct {
	Aid        string `json:"aid"`
	Title      string `json:"title"`
	BriefIntro string `json:"-"`
	Url        string `json:"url"`
}
