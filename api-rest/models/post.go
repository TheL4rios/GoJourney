package models

import "time"

type Post struct {
	Id          string 		`json:id`
	PostContent string 		`json:post_content`
	UserId      string 		`json:user_id`
	CreatedAt 	time.Time 	`json:created_at`
}