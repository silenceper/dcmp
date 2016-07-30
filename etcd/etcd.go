package etcd

import (
	"log"
	"time"

	"golang.org/x/net/context"

	"github.com/coreos/etcd/client"
)

var globalEndpoints []string
var globalKapi client.KeysAPI

func InitConfig(endpoints []string) {
	globalEndpoints = endpoints
}

func getKapi() client.KeysAPI {
	if globalKapi == nil {
		cfg := client.Config{
			Endpoints: globalEndpoints,
			Transport: client.DefaultTransport,
			// set timeout per request to fail fast when the target endpoint is unavailable
			HeaderTimeoutPerRequest: time.Second,
		}
		c, err := client.New(cfg)
		//client.EnablecURLDebug()
		if err != nil {
			log.Fatal(err)
		}
		globalKapi = client.NewKeysAPI(c)
	}
	return globalKapi
}

func GetKeyList(path string) (*client.Response, error) {
	kapi := getKapi()

	opts := &client.GetOptions{}
	opts.Recursive = true
	return kapi.Get(context.Background(), path, opts)
}

func UpdateKey(key, value string) (*client.Response, error) {
	kapi := getKapi()
	return kapi.Update(context.Background(), key, value)
}

func SetKey(key, value string, opts *client.SetOptions) (*client.Response, error) {
	kapi := getKapi()
	return kapi.Set(context.Background(), key, value, opts)
}

func DeleteKey(key string, opts *client.DeleteOptions) (*client.Response, error) {
	kapi := getKapi()
	return kapi.Delete(context.Background(), key, opts)
}
