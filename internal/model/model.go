package model

import (
	"time"

	"github.com/spf13/cast"
)

// BaseModel 模型基类
type BaseModel struct {
	ID uint64 `json:"id,omitempty" gorm:"column:id;primaryKey;autoIncrement;"`
}

// CommonTimestampsField 时间戳
type CommonTimestampsField struct {
	CreateAt time.Time `json:"create_at,omitempty" gorm:"create_at;index;" valid:"create_at"`
	UpdateAt time.Time `json:"update_at,omitempty" gorm:"update_at;index;" valid:"update_at"`
}

func (b BaseModel) GetStringID() string {
	return cast.ToString(b.ID)
}
