package requestStruct

import "mime/multipart"

type InfluencerFeedback struct {
	Name       string                `form:"name" binding:"required"`
	LinkedIn   string                `form:"linkedin" binding:"required"`
	Post       *multipart.FileHeader `form:"post,omitempty"`
	Company    string                `form:"company" binding:"required"`
	Feedback   string                `form:"feedback" binding:"required"`
	Influencer interface{}
	ImageLink  string
}
