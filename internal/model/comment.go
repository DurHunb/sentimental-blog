package model

type Comment struct {
	*Model
	CommentUid int    `json:"comment_uid"`
	Content    string `json:"content"`
	LikeNum    string `json:"like_num"`
	ArticleUid int    `json:"article_uid"`
	ArticleId  int    `json:"article_id"`
	Ip         string `json:"ip"`
	IpLoc      string `json:"ip_loc"`
	Status     int    `json:"status"`
}
