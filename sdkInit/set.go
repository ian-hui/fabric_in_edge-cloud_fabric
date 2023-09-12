package sdkInit

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
)

func (t *Application) Set(usrinfo UserInfo) (string, error) {
	b, err := json.Marshal(usrinfo)
	if err != nil {
		return "", fmt.Errorf("指定的userinfo对象序列化时发生错误")
	}
	request := channel.Request{ChaincodeID: t.SdkEnvInfo.ChaincodeID, Fcn: "set", Args: [][]byte{[]byte(usrinfo.UserId), b}}
	response, err := t.SdkEnvInfo.ChClient.Execute(request)
	if err != nil {
		// set失败
		return "", err
	}

	//fmt.Println("============== response:",response)

	return string(response.TransactionID), nil
}

func (t *Application) SetAccess(fileaccessinfo FileAccessInfo) (string, error) {
	b, err := json.Marshal(fileaccessinfo)
	if err != nil {
		return "", fmt.Errorf("指定的fileaccessinfo对象序列化时发生错误")
	}

	request := channel.Request{ChaincodeID: t.SdkEnvInfo.ChaincodeID, Fcn: "set", Args: [][]byte{[]byte(fileaccessinfo.FileId), b}}
	response, err := t.SdkEnvInfo.ChClient.Execute(request)
	if err != nil {
		// set失败
		return "", err
	}

	//fmt.Println("============== response:",response)

	return string(response.TransactionID), nil
}
