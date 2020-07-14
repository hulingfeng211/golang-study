package main

import (
	"fmt"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/model"
	"github.com/nacos-group/nacos-sdk-go/utils"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"log"
)

const (
	ServerName  = "DEMO-GO"
	ClusterName = "a"
	DataId      = "demo-dev.yaml"
	GROUP       = "DEFAULT_GROUP"
)

//
//func SubscribeConfig()  {
//
//}
//content, err := configClient.GetConfig(vo.ConfigParam{
//DataId: "dataId",
//Group:  "group"})

func main() {

	//namingClient,err:=clients.
	RegisterService()

	//namingClient.Unsubscribe(&vo.SubscribeParam{
	//	ServiceName: SERVER_NAME,
	//	Clusters:    []string{"a"},
	//	SubscribeCallback: func(services []model.SubscribeService, err error) {
	//		log.Printf("\n\n callback return services:%s \n\n", utils.ToJsonString(services))
	//	},
	//})

	select {}
}

func RegisterService() {
	//nacos 配置信息
	serverConfigs := []constant.ServerConfig{
		{IpAddr: "localhost",
			Port: 8848},
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
		Ip:          "127.0.0.1",
		Port:        8858,
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
	content, err := configClient.GetConfig(vo.ConfigParam{
		DataId: DataId,
		Group:  GROUP})

	fmt.Println(content)

	configClient.ListenConfig(vo.ConfigParam{
		DataId: DataId,
		Group:  GROUP,
		OnChange: func(namespace, group, dataId, data string) {
			fmt.Println("group:" + group + ", dataId:" + dataId + ", data:" + data)
		},
	})
}
