package src

import (
	"errors"
	"fmt"
	"github.com/ArkNX/ark-go/interface"
	"github.com/ArkNX/ark-go/plugin/aliyunPlugin/ossModule"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"io/ioutil"
	"reflect"
	"runtime"
	"strings"
)

func init() {
	ossModule.ModuleName = ark.GetName((*AFCOssModule)(nil))
	ossModule.ModuleType = ark.GetType((*AFCOssModule)(nil))
	ossModule.ModuleUpdate = runtime.FuncForPC(reflect.ValueOf((&AFCOssModule{}).Update).Pointer()).Name()
}

type AFCOssModule struct {
	ark.AFCModule
	// other data
	cfg    ossModule.Config
	client *oss.Client
	bucket *oss.Bucket
}

func (ossModule *AFCOssModule) Init() error {
	return nil
}
func (ossModule *AFCOssModule) Connect(cfg ossModule.Config) error {
	cli, err := oss.New(cfg.Endpoint, cfg.AccessKey, cfg.SecretKey)
	if err != nil {
		return err
	}

	bucket, err := cli.Bucket(cfg.Bucket)
	if err != nil {
		return err
	}

	ossModule.cfg = cfg
	ossModule.client = cli
	ossModule.bucket = bucket
	return nil
}

func (ossModule *AFCOssModule) Bucket() (*oss.Bucket, error) {
	if ossModule.bucket != nil {
		return nil, errors.New("oss bucket is nil")
	}
	return ossModule.bucket, nil
}

func (ossModule *AFCOssModule) Client() (*oss.Client, error) {
	if ossModule.bucket != nil {
		return nil, errors.New("oss Client is nil")
	}
	return ossModule.client, nil
}

func (ossModule *AFCOssModule) PutObjectFromFile(remotePath, localPath string) error {
	return ossModule.bucket.PutObjectFromFile(remotePath, localPath)
}

func (ossModule *AFCOssModule) GetObjectToFile(remotePath, localPath string) error {
	return ossModule.bucket.GetObjectToFile(remotePath, localPath)
}

func (ossModule *AFCOssModule) GetObject(path string) ([]byte, error) {
	body, err := ossModule.bucket.GetObject(path)
	if err != nil {
		return nil, err
	}
	defer body.Close()

	data, err := ioutil.ReadAll(body)
	if err != nil {
		return nil, err
	}

	if len(data) == 0 {
		return nil, fmt.Errorf("file %s is empty", path)
	}

	return data, nil
}

func (ossModule *AFCOssModule) PutObject(objectKey, data string) error {
	return ossModule.bucket.PutObject(objectKey, strings.NewReader(data))
}
