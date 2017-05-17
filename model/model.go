package model

import "time"

type (

	User struct {
		UserID         int64     `db:"user_id"`
		FullName       string    `db:"full_name"`
		UserName       string    `db:"username"`
		Bio            string    `db:"bio"`
		Mailaddress    string    `db:"mailaddress"`
		ProfilePicture string    `db:"profile_picture"`
		CreatedTime    time.Time `db:"create_time"`
		PrivateFlg     int64     `db:"private_flg"`
		Token          string    `db:"token"`
	}

	UserResponse struct {
		UserID         int64     `json:"id"`
		FullName       string    `json:"full_name"`
		UserName       string    `json:"username"`
		Bio            string    `json:"bio"`
		Mailaddress    string    `json:"mailaddress"`
		ProfilePicture string    `json:"profile_picture"`
		CreatedTime    string 	 `json:"created_time"`
		PrivateFlg     int64     `json:"private_flg"`
		Password          string    `json:"password"`
	}

	TimelineResponse struct {
		MediaID     int64	`json:"media_id"`
		UserID      int64	`json:"user_id"`
		CreatedTime string	`json:"created_time"`
		Picture     string	`json:"img_path"`
		Body        string	`json:"caption"`
		LikeCount int64 `json:"like_count"`
		User []UserResponse `json:"user"`
		LikeCounts int `json:"like_counts"`
		IsLiked bool `json:"is_liked"`
	}

	LikesResponse struct {
		Counts int `json:"counts"`
	}

	UserMediaResponse struct {
		MediaID     int64	`json:"id"`
		CreatedTime string	`json:"created_time"`
		Picture     string	`json:"img_path"`
		Body        string	`json:"caption"`
		LikeCount int64 `json:"like_count"`
		User []UserResponse `json:"user"`
		LikeCounts int `json:"like_counts"`
		IsLiked bool `json:"is_liked"`
	}

	Userinfo struct {
		ID        int    `db:"id"`
		Email     string `db:"email"`
		Firstname string `db:"first_name"`
		Lastname  string `db:"last_name"`
	}

	UserinfoJSON struct {
		ID        int    `json:"id"`
		Email     string `json:"email"`
		Firstname string `json:"firstName"`
		Lastname  string `json:"lastName"`
	}

	ResponseData struct {
		//User  []userinfo `json:"users"`
		Users []UserResponse `json:"user"`
		Timeline []TimelineResponse `json:"timeline"`
	}
)
