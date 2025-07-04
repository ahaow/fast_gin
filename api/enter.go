package api

import (
	"fast_gin/api/file_api"
	"fast_gin/api/image_api"
	"fast_gin/api/user_api"
)

type Api struct {
	UserApi   user_api.UserApi
	ImagesApi image_api.ImagesApi
	FilesApi  file_api.FileApi
}

var App = new(Api)
