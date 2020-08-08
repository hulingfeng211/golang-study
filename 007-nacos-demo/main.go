package main

import (
	"gopkg.in/yaml.v3"
	"log"

	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/model"
	"github.com/nacos-group/nacos-sdk-go/utils"
	"github.com/nacos-group/nacos-sdk-go/vo"
)

type Server struct {
	Name string `json:"name" yaml:"name"`
	Port uint64 `json:"port" yaml:"port"`
}

type ConfigModel struct {
	Server Server `json:"server" yaml:"server"`
	Value  string `json:"value" yaml:"value"`
}

const (
	ServerName  = "DEMO-GO"
	ClusterName = "a"
	DataId      = "demo-dev.yaml"
	GROUP       = "DEFAULT_GROUP"
)

var nacosServer = "localhost"
var nacosPort = 8848
var localServer = "localhost"
var localPort = 8858

var configModel = ConfigModel{}

func main() {
	RegisterService(nacosServer, uint64(nacosPort))

	select {}
}

/*
nacosServer：nacos服务器ip地址
nacosPort:nacos服务器端口
*/

func OnConfigRefreshHandle(namespace, group, dataId, data string) {
	err := yaml.Unmarshal([]byte(data), &configModel)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	log.Println(configModel)
}

func RegisterService(nacosServer string, nacosPort uint64) {
	//nacos 配置信息
	serverConfigs := []constant.ServerConfig{
		{IpAddr: nacosServer,
			Port: nacosPort},
	}
	clientConfig := constant.ClientConfig{
		TimeoutMs:           5000,
		ListenInterval:      10000,
		NotLoadCacheAtStart: true,
		LogDir:              "data/log",
	}
	namingClient, err := clients.CreateNamingClient(map[string]interface{}{

		constant.KEY_SERVER_CONFIGS: serverConfigs,

		constant.KEY_CLIENT_CONFIG: clientConfig,
	})
	if err != nil {
		panic(err)
	}
	configClient, err := clients.CreateConfigClient(map[string]interface{}{
		"serverConfigs": serverConfigs,
		"clientConfig":  clientConfig,
	})
	if err != nil {
		panic(err)
	}
	success, _ := namingClient.RegisterInstance(vo.RegisterInstanceParam{
		Ip:          localServer,
		Port:        uint64(localPort),
		ServiceName: ServerName,
		Weight:      10,
		ClusterName: ClusterName,
		Enable:      true,
		Healthy:     true,
		Ephemeral:   true,
	})
	if !success {
		panic("注册失败")
	}
	namingClient.Subscribe(&vo.SubscribeParam{
		ServiceName: ServerName,
		Clusters:    []string{ClusterName},
		SubscribeCallback: func(services []model.SubscribeService, err error) {
			log.Printf("\n\n callback return services:%s \n\n", utils.ToJsonString(services))
		},
	})
	//read config from nacos
	//content, err := configClient.GetConfig(vo.ConfigParam{
	//	DataId: DataId,
	//	Group:  GROUP})

	//fmt.Println(configModel)
	configClient.ListenConfig(vo.ConfigParam{
		DataId: DataId,
		Group:  GROUP,
		OnChange: func(namespace, group, dataId, data string) {
			//fmt.Println("group:" + group + ", dataId:" + dataId + ", data:" + data)
			//err = yaml.Unmarshal([]byte(data), &configModel)
			//if err != nil {
			//	log.Fatalf("error: %v", err)
			//}
			OnConfigRefreshHandle(namespace, group, dataId, data)
		},
	})
}
