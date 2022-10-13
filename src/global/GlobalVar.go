package global

import (
	sf "github.com/bwmarrin/snowflake"
	"github.com/coocood/freecache"
	ut "github.com/go-playground/universal-translator"
	"github.com/jinzhu/gorm"
	"github.com/minio/minio-go"
	"go.uber.org/zap"
	"winterchen.com/my-site-go/src/config"
)

var (
	Configs     config.Config
	Trans       ut.Translator
	DB          *gorm.DB
	Log         *zap.Logger
	MinioClient *minio.Client
	Snowflake   *sf.Node
	Cache       *freecache.Cache
)
