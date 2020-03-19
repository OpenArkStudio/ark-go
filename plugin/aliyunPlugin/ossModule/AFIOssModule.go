package ossModule

import (
	ark "github.com/ArkNX/ark-go/interface"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"reflect"
)

var (
	ModuleName   string
	ModuleType   reflect.Type
	ModuleUpdate string
)

type Config struct {
	AccessKey string `env:"ALIYUN_AK"`
	SecretKey string `env:"ALIYUN_SK"`
	Endpoint  string `env:"ALIYUN_ENDPOINT"`
	Bucket    string `env:"ALIYUN_BUCKET"`
}

type AFIOssModule interface {
	ark.AFIModule
	Connect(cfg Config) error
	Bucket() (*oss.Bucket, error)
	Client() (*oss.Client, error)
	PutObjectFromFile(remotePath, localPath string) error
	GetObjectToFile(remotePath, localPath string) error
	GetObject(path string) ([]byte, error)
	PutObject(objectKey, data string) error
}
