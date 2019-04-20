package model

import (
	"github.com/IBM/cloudland/web/sca/dbs"
)

type Instance struct {
	Model
	Hostname    string        `gorm:"type:varchar(128)"`
	Domain      string        `gorm:"type:varchar(128)"`
	Status      string        `gorm:"type:varchar(32)"`
	Reason      string        `gorm:"type:text"`
	FloatingIps []*FloatingIp `gorm:"foreignkey:InstanceID"`
	Interfaces  []*Interface  `gorm:"foreignkey:InstanceID"`
	FlavorID    int64
	Flavor      *Flavor `gorm:"foreignkey:FlavorID"`
	ImageID     int64
	Image       *Image `gorm:"foreignkey:ImageID"`
	Keys        []*Key `gorm:"many2many:InstanceKeys;"`
	Userdata    string `gorm:"type:text"`
	Hyper       int32  `gorm:"default:-1"`
}

func init() {
	dbs.AutoMigrate(&Instance{})
}