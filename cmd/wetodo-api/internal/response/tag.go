package response

import (
	"wetodo/internal/storage"
)

// TagListResponse represent collection of tags
type TagListResponse struct {
	Tags []Tag `json:"tags"`
}

type Tag struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	CreatedAt int    `json:"createdAt"`
	UpdatedAt int    `json:"updatedAt"`
}

func CopyToTagListResponse(tags []storage.Tag) TagListResponse {
	ts := make([]Tag, 0)
	for _, tag := range tags {
		ts = append(ts, Tag{
			ID:        tag.ID,
			Name:      tag.Name,
			CreatedAt: int(tag.CreatedAt.Unix()),
			UpdatedAt: int(tag.UpdatedAt.Unix()),
		})
	}
	return TagListResponse{Tags: ts}
}
