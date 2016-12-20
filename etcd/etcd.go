package etcd

import (
	"log"
	"net/http"
	"os"
	"time"

	"golang.org/x/net/context"

	"github.com/coreos/etcd/client"
	"github.com/coreos/etcd/pkg/transport"
)

var globalEndpoints []string
var globalCaFile string
var globalCertFile string
var globalKeyFile string

var globalKapi client.KeysAPI

func InitConfig(endpoints []string, caFile string, certFile string, keyFile string) {
	globalEndpoints = endpoints
	globalCaFile = caFile
	globalCertFile = certFile
	globalKeyFile = keyFile
}

func getKapi() client.KeysAPI {
	if globalKapi == nil {
		transport, err := getTransport()
		if err != nil {
			log.Fatal(err)
		}
		cfg := client.Config{
			Endpoints: globalEndpoints,
			Transport: transport,
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

func getTransport() (*http.Transport, error) {
	if globalCaFile == "" {
		globalCaFile = os.Getenv("DCMP_CA_FILE")
	}
	if globalCertFile == "" {
		globalCertFile = os.Getenv("DCMP_CERT_FILE")
	}
	if globalKeyFile == "" {
		globalKeyFile = os.Getenv("DCMP_KEY_FILE")
	}

	tls := transport.TLSInfo{
		CAFile:   globalCaFile,
		CertFile: globalCertFile,
		KeyFile:  globalKeyFile,
	}

	return transport.NewTransport(tls, 30*time.Second)
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
