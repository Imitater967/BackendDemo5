package models

type VideoFavoriteModel struct {
	User  int64
	Video int64
}

func (*VideoFavoriteModel) TableName() string {
	return "video_favorite"
}
