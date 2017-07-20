package instabot

type UserInfo struct {
	IsBusiness          bool   `json:"is_business"`
	ProfilePicURL       string `json:"profile_pic_url"`
	HdProfilePicURLInfo struct {
		Width  int    `json:"width"`
		Height int    `json:"height"`
		URL    string `json:"url"`
	} `json:"hd_profile_pic_url_info"`
	UsertagsCount              int    `json:"usertags_count"`
	ExternalLynxURL            string `json:"external_lynx_url"`
	FollowingCount             int    `json:"following_count"`
	HasAnonymousProfilePicture bool   `json:"has_anonymous_profile_picture"`
	GeoMediaCount              int    `json:"geo_media_count"`
	ExternalURL                string `json:"external_url"`
	Username                   string `json:"username"`
	Biography                  string `json:"biography"`
	HasChaining                bool   `json:"has_chaining"`
	FullName                   string `json:"full_name"`
	IsPrivate                  bool   `json:"is_private"`
	PK                         int    `json:"pk"`
	FollowerCount              int    `json:"follower_count"`
	ProfilePicID               string `json:"profile_pic_id"`
	IsVerified                 bool   `json:"is_verified"`
	HdProfilePicVersions       []struct {
		Width  int    `json:"width"`
		Height int    `json:"height"`
		URL    string `json:"url"`
	} `json:"hd_profile_pic_versions"`
	MediaCount int  `json:"media_count"`
	IsFavorite bool `json:"is_favorite"`
}

type User struct {
	PK                         int               `json:"pk"`
	Username                   string            `json:"username"`
	FullName                   string            `json:"full_name"`
	IsPrivate                  bool              `json:"is_private"`
	IsVerified                 bool              `json:"is_verified"`
	ProfilePicURL              string            `json:"profile_pic_url"`
	ProfilePicID               string            `json:"profile_pic_id,omitempty"`
	HasAnonymousProfilePicture bool              `json:"has_anonymous_profile_picture"`
	IsFavorite                 bool              `json:"is_favorite"`
	IsUnpublished              bool              `json:"is_unpublished"`
	FriendshipStatus           *FriendshipStatus `json:"friendship_status"`
}

type FriendshipStatus struct {
	IncomingRequest bool `json:"incoming_request"`
	IsPrivate       bool `json:"is_private"`
	Following       bool `json:"following"`
	OutgoingRequest bool `json:"outgoing_request"`
	FollowedBy      bool `json:"followed_by"`
	Blocking        bool `json:"blocking"`
}

type LoggedInUser struct {
	AllowContactsSync bool   `json:"allow_contacts_sync"`
	FullName          string `json:"full_name"`
	IsPrivate         bool   `json:"is_private"`
	IsVerified        bool   `json:"is_verified"`
	ProfilePic        string `json:"profile_pic"`
	Username          string `json:"username"`
	AnonProfilePic    bool   `json:"has_anonymous_profile_picture"`
	PK                int    `json:"pk"`
}

type Item struct {
	TakenAt         int    `json:"taken_at"`
	PK              int    `json:"pk"`
	ID              string `json:"id"`
	DeviceTimestamp int    `json:"device_timestamp"`
	MediaType       int    `json:"media_type"`
	Code            string `json:"code"`
	ClientCacheKey  string `json:"client_cache_key"`
	FilterType      int    `json:"filter_type"`
	ImageVersions2  struct {
		Candidates []struct {
			Width  int    `json:"width"`
			Height int    `json:"height"`
			URL    string `json:"url"`
		} `json:"candidates"`
	} `json:"image_versions2"`
	OriginalWidth                int           `json:"original_width"`
	OriginalHeight               int           `json:"original_height"`
	User                         *User         `json:"user"`
	OrganicTrackingToken         string        `json:"organic_tracking_token"`
	LikeCount                    int           `json:"like_count"`
	TopLikers                    []interface{} `json:"top_likers"`
	HasLiked                     bool          `json:"has_liked"`
	CommentLikesEnabled          bool          `json:"comment_likes_enabled"`
	HasMoreComments              bool          `json:"has_more_comments"`
	MaxNumVisiblePreviewComments int           `json:"max_num_visible_preview_comments"`
	PreviewComments              []interface{} `json:"preview_comments"`
	Comments                     []interface{} `json:"comments"`
	CommentCount                 int           `json:"comment_count"`
	Caption                      struct {
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
		MediaID      int64  `json:"media_id"`
	} `json:"caption"`
	CaptionIsEdited bool  `json:"caption_is_edited"`
	PhotoOfYou      bool  `json:"photo_of_you"`
	NextMaxID       int64 `json:"next_max_id,omitempty"`
	Location        struct {
		Pk               int     `json:"pk"`
		Name             string  `json:"name"`
		Address          string  `json:"address"`
		City             string  `json:"city"`
		Lng              float64 `json:"lng"`
		Lat              float64 `json:"lat"`
		ExternalSource   string  `json:"external_source"`
		FacebookPlacesID int64   `json:"facebook_places_id"`
	} `json:"location,omitempty"`
	Lat           float64 `json:"lat,omitempty"`
	Lng           float64 `json:"lng,omitempty"`
	ViewCount     float64 `json:"view_count,omitempty"`
	VideoVersions []struct {
		Type   int    `json:"type"`
		Width  int    `json:"width"`
		Height int    `json:"height"`
		URL    string `json:"url"`
	} `json:"video_versions,omitempty"`
	HasAudio      bool    `json:"has_audio,omitempty"`
	VideoDuration float64 `json:"video_duration,omitempty"`
}

type HashtagItem struct {
	TakenAt         int    `json:"taken_at"`
	PK              int    `json:"pk"`
	ID              string `json:"id"`
	DeviceTimestamp int    `json:"device_timestamp"`
	MediaType       int    `json:"media_type"`
	Code            string `json:"code"`
	ClientCacheKey  string `json:"client_cache_key"`
	FilterType      int    `json:"filter_type"`
	ImageVersions2  struct {
		Candidates []struct {
			Width  int    `json:"width"`
			Height int    `json:"height"`
			URL    string `json:"url"`
		} `json:"candidates"`
	} `json:"image_versions2"`
	OriginalWidth                int           `json:"original_width"`
	OriginalHeight               int           `json:"original_height"`
	User                         *User         `json:"user"`
	OrganicTrackingToken         string        `json:"organic_tracking_token"`
	LikeCount                    int           `json:"like_count"`
	HasLiked                     bool          `json:"has_liked"`
	CommentLikesEnabled          bool          `json:"comment_likes_enabled"`
	HasMoreComments              bool          `json:"has_more_comments"`
	MaxNumVisiblePreviewComments int           `json:"max_num_visible_preview_comments"`
	PreviewComments              []interface{} `json:"preview_comments"`
	Comments                     []interface{} `json:"comments"`
	CommentCount                 int           `json:"comment_count"`
	Caption                      struct {
		Pk           int64  `json:"pk"`
		UserID       int64  `json:"user_id"`
		Text         string `json:"text"`
		Type         int    `json:"type"`
		CreatedAt    int    `json:"created_at"`
		CreatedAtUtc int    `json:"created_at_utc"`
		ContentType  string `json:"content_type"`
		Status       string `json:"status"`
		BitFlags     int    `json:"bit_flags"`
		User         *User  `json:"user"`
		MediaID      int64  `json:"media_id"`
	} `json:"caption"`
	CaptionIsEdited bool  `json:"caption_is_edited"`
	PhotoOfYou      bool  `json:"photo_of_you"`
	NextMaxID       int64 `json:"next_max_id,omitempty"`
	Usertags        struct {
		In []interface{} `json:"in"`
	} `json:"usertags,omitempty"`
}
