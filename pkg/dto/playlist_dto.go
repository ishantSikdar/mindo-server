package dto

import "time"

type CreatePlaylistRequest struct {
	Name         string   `json:"name"         binding:"required"`
	Description  string   `json:"description"  binding:"required"`
	DomainName   string   `json:"domainName"   binding:"required"`
	ThumbnailURL string   `json:"thumbnailUrl" binding:"required"`
	IsAIGen      bool     `json:"isAIGen"`
	Topics       []string `json:"topics"       binding:"required,dive"`
}

type PlaylistDetailsDTO struct {
	ID           string          `json:"id"`
	Name         string          `json:"name"`
	Description  string          `json:"description"`
	InterestID   string          `json:"interestId"`
	ThumbnailURL string          `json:"thumbnailUrl"`
	Views        int             `json:"views"`
	Code         string          `json:"code"`
	CreatedAt    time.Time       `json:"createdAt"`
	UpdatedAt    time.Time       `json:"updatedAt"`
	UpdatedBy    string          `json:"updatedBy"`
	IsAIGen      bool            `json:"isAIGen"`
	Topics       []TopicsMiniDTO `json:"topics"`
}

type TopicsMiniDTO struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	VideoID     string `json:"videoId"`
	TopicNumber int    `json:"topicNumber"`
}

type PlaylistPreviewDTO struct {
	ID           string    `json:"id"`
	Name         string    `json:"name"`
	Description  string    `json:"description"`
	InterestID   string    `json:"interestId"`
	ThumbnailURL string    `json:"thumbnailUrl"`
	Views        int       `json:"views"`
	Code         string    `json:"code"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
	UpdatedBy    string    `json:"updatedBy"`
	IsAIGen      bool      `json:"isAIGen"`
	TopicsCount  int       `json:"topicsCount,omitempty"`
}

type VideoDataDTO struct {
	ID           string    `json:"id"`
	TopicID      string    `json:"topicId"`
	VideoID      string    `json:"videoId"`
	Title        string    `json:"title"`
	VideoDate    time.Time `json:"videoPublishedAt"`
	ChannelTitle string    `json:"channelTitle"`
	ThumbnailURL string    `json:"thumbnailUrl"`
	ExpiryAt     time.Time `json:"expiryAt"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
	UpdatedBy    string    `json:"updatedBy"`
}

type VideoMiniDTO struct {
	VideoID      string    `json:"videoId"`
	Title        string    `json:"title"`
	VideoDate    time.Time `json:"videoPublishedAt"`
	ThumbnailURL string    `json:"thumbnailUrl"`
	ChannelTitle string    `json:"channelTitle"`
}

type GroupedVideoDataResponse struct {
	Video      VideoDataDTO   `json:"video"`
	MoreVideos []VideoDataDTO `json:"moreVideos"`
}

type GeneratedPlaylist struct {
	Description string   `json:"description"`
	Title       string   `json:"subject"`
	Topics      []string `json:"syllabus"`
}

type GeneratePlaylistParams struct {
	Title string `json:"subject"`
}
