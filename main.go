package main

import (
	"fabric-go-sdk/NodeUtils"
	"fabric-go-sdk/clients"
	"fabric-go-sdk/sdkInit"
)

var center = NodeUtils.Nodestructure{
	KafkaIp:      "0.0.0.0:9091",
	Couchdb_addr: "http://admin:123456@0.0.0.0:5984",
	AreaId:       "0",
}

var peertopics = []string{"register", "upload", "filereq", "KeyUpload", "ReceiveKeyUpload", "ReceiveKeyReq", "DataForwarding", "ReceiveFileRequestFromCenter"}
var centertopics = []string{"uploadposition", "UploadCiText", "KeyReqForwarding", "FileReqestToCenter", "UploadKeyPosition"}

func main() {

	// sdk setup
	clients.InitFabric("./cfg/org1conf.yaml", &clients.UserinfoChannel_info)
	defer clients.UserinfoChannel_info.EvClient.Unregister(sdkInit.BlockListener(clients.UserinfoChannel_info.EvClient))
	defer clients.UserinfoChannel_info.EvClient.Unregister(sdkInit.ChainCodeEventListener(clients.UserinfoChannel_info.EvClient, clients.UserinfoChannel_info.ChaincodeID))

	// clients.InitFabric("./clients/config2.yaml", &clients.Channel2_info, &clients.RegionAAccessApp)
	// defer clients.Channel2_info.EvClient.Unregister(sdkInit.BlockListener(clients.Channel2_info.EvClient))
	// defer clients.Channel2_info.EvClient.Unregister(sdkInit.ChainCodeEventListener(clients.Channel2_info.EvClient, clients.Channel2_info.ChaincodeID))

	// clients.InitFabric("./clients/config3.yaml", &clients.Channel3_info, &clients.RegionBUserApp)
	// defer clients.Channel3_info.EvClient.Unregister(sdkInit.BlockListener(clients.Channel3_info.EvClient))
	// defer clients.Channel3_info.EvClient.Unregister(sdkInit.ChainCodeEventListener(clients.Channel3_info.EvClient, clients.Channel3_info.ChaincodeID))

	// clients.InitFabric("./clients/config4.yaml", &clients.Channel4_info, &clients.RegionBAccessApp)
	// defer clients.Channel4_info.EvClient.Unregister(sdkInit.BlockListener(clients.Channel4_info.EvClient))
	// defer clients.Channel4_info.EvClient.Unregister(sdkInit.ChainCodeEventListener(clients.Channel4_info.EvClient, clients.Channel4_info.ChaincodeID))

	// //启动peer消费者
	// NodeUtils.InitPeerNode(peertopics, peer0_org1)
	// NodeUtils.InitPeerNode(peertopics, peer0_org2)
	// NodeUtils.InitPeerNode(peertopics, peer1_org1)
	// NodeUtils.InitPeerNode(peertopics, peer1_org2)
	// NodeUtils.InitPeerNode(peertopics, peer0_org3)
	// NodeUtils.InitPeerNode(peertopics, peer0_org4)
	// NodeUtils.InitPeerNode(peertopics, peer1_org3)
	// NodeUtils.InitPeerNode(peertopics, peer1_org4)
	//启动center消费者
	// NodeUtils.InitCenterNode(centertopics, center)

	//启动zookeeper
	// clients.InitZookeeper()
	//以下为gin框架以及kafka的代码

}
