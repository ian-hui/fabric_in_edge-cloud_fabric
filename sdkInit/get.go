package sdkInit

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
)

func (t *Application) Get(userid string, endpoint string) (UserInfo, error) {
	var userif UserInfo
	response, err := t.SdkEnvInfo.ChClient.Query(channel.Request{ChaincodeID: t.SdkEnvInfo.ChaincodeID, Fcn: "get", Args: [][]byte{[]byte(userid)}},
		channel.WithTargetEndpoints(endpoint))
	if err != nil {
		return userif, fmt.Errorf("failed to query: %v", err)
	}
	// 对查询到的状态进行反序列化
	err = json.Unmarshal(response.Payload, &userif)
	if err != nil {
		return userif, err
	}
	return userif, nil
}

func (t *Application) GetAccess(fileid string, endpoint string) (FileAccessInfo, error) {
	var fileaccessinfo FileAccessInfo
	response, err := t.SdkEnvInfo.ChClient.Query(channel.Request{ChaincodeID: t.SdkEnvInfo.ChaincodeID, Fcn: "get", Args: [][]byte{[]byte(fileid)}},
		channel.WithTargetEndpoints(endpoint))
	if err != nil {
		return fileaccessinfo, fmt.Errorf("failed to query: %v", err)
	}
	// 对查询到的状态进行反序列化
	err = json.Unmarshal(response.Payload, &fileaccessinfo)
	if err != nil {
		return fileaccessinfo, err
	}
	return fileaccessinfo, nil
}
