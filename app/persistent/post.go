package persistent

import (
	"kitstart_goweb/app/entity"
)

// PostPersistentInterface Post 接口
type PostPersistentInterface interface {
	SavePost(*entity.Post) error
}

// PostPersistent Post
type PostPersistent struct {
}

// SavePost - 保存 Post
func (post *PostPersistent) SavePost(entity *entity.Post) error {
	return nil
}
