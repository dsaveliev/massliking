package instabot

type GetUserFeedResponse struct {
	Items               []*Item `json:"items"`
	NumResults          int     `json:"num_results"`
	MoreAvailable       bool    `json:"more_available"`
	NextMaxID           string  `json:"next_max_id"`
	AutoLoadMoreEnabled bool    `json:"auto_load_more_enabled"`
	Status              string  `json:"status"`
}

type GetUserFollowersResponse struct {
	Users     []*User `json:"users"`
	BigList   bool    `json:"big_list"`
	PageSize  int     `json:"page_size"`
	NextMaxID string  `json:"next_max_id"`
	Status    string  `json:"status"`
}

type GetUserFollowingsResponse struct {
	Users     []*User `json:"users"`
	BigList   bool    `json:"big_list"`
	PageSize  int     `json:"page_size"`
	NextMaxID string  `json:"next_max_id"`
	Status    string  `json:"status"`
}

type GetHashtagFeedResponse struct {
	Items               []*HashtagItem `json:"items"`
	NumResults          int            `json:"num_results"`
	NextMaxID           string         `json:"next_max_id"`
	MoreAvailable       bool           `json:"more_available"`
	AutoLoadMoreEnabled bool           `json:"auto_load_more_enabled"`
	Status              string         `json:"status"`
}

type GetMediaCommentsResponse struct {
	CommentLikesEnabled bool `json:"comment_likes_enabled"`
	Comments            []struct {
		Pk               int64  `json:"pk"`
		UserID           int    `json:"user_id"`
		Text             string `json:"text"`
		Type             int    `json:"type"`
		CreatedAt        int    `json:"created_at"`
		CreatedAtUtc     int    `json:"created_at_utc"`
		ContentType      string `json:"content_type"`
		Status           string `json:"status"`
		BitFlags         int    `json:"bit_flags"`
		User             *User  `json:"user"`
		HasLikedComment  bool   `json:"has_liked_comment"`
		CommentLikeCount int    `json:"comment_like_count"`
	} `json:"comments"`
	CommentCount int `json:"comment_count"`
	Caption      struct {
		Pk           int64  `json:"pk"`
		UserID       int    `json:"user_id"`
		Text         string `json:"text"`
		Type         int    `json:"type"`
		CreatedAt    int    `json:"created_at"`
		CreatedAtUtc int    `json:"created_at_utc"`
		ContentType  string `json:"content_type"`
		Status       string `json:"status"`
		BitFlags     int    `json:"bit_flags"`
		User         *User  `json:"user"`
	} `json:"caption"`
	CaptionIsEdited         bool `json:"caption_is_edited"`
	HasMoreComments         bool `json:"has_more_comments"`
	HasMoreHeadloadComments bool `json:"has_more_headload_comments"`
	PreviewComments         []struct {
		Pk           int64  `json:"pk"`
		UserID       int    `json:"user_id"`
		Text         string `json:"text"`
		Type         int    `json:"type"`
		CreatedAt    int    `json:"created_at"`
		CreatedAtUtc int    `json:"created_at_utc"`
		ContentType  string `json:"content_type"`
		Status       string `json:"status"`
		BitFlags     int    `json:"bit_flags"`
		User         *User  `json:"user"`
	} `json:"preview_comments"`
	NextMaxID string `json:"next_max_id"`
	Status    string `json:"status"`
}

type GetMediaLikersResponse struct {
	Users     []*User `json:"users"`
	UserCount int     `json:"user_count"`
	Status    string  `json:"status"`
}

type LikeResponse struct {
	Status string `json:"status"`
}

type FollowResponse struct {
	*FriendshipStatus `json:"friendship_status"`
	Status            string `json:"status"`
}

type UnfollowResponse struct {
	LoggedInUser struct {
		Pk                         int64  `json:"pk"`
		Username                   string `json:"username"`
		FullName                   string `json:"full_name"`
		IsPrivate                  bool   `json:"is_private"`
		ProfilePicURL              string `json:"profile_pic_url"`
		IsVerified                 bool   `json:"is_verified"`
		HasAnonymousProfilePicture bool   `json:"has_anonymous_profile_picture"`
		AllowContactsSync          bool   `json:"allow_contacts_sync"`
	} `json:"logged_in_user"`
	Status string `json:"status"`
}

type CommentResponse struct {
	Comment struct {
		ContentType  string  `json:"content_type"`
		User         *User   `json:"user"`
		Pk           int     `json:"pk"`
		Text         string  `json:"text"`
		Type         int     `json:"type"`
		CreatedAt    float64 `json:"created_at"`
		CreatedAtUtc int     `json:"created_at_utc"`
		MediaID      int64   `json:"media_id"`
		Status       string  `json:"status"`
	} `json:"comment"`
	Status string `json:"status"`
}

type LoginResponse struct {
	*LoggedInUser `json:"logged_in_user"`
	Status        string `json:"status"`
}

type SearchUsernameResponse struct {
	*UserInfo `json:"user"`
	Status    string `json:"status"`
}
