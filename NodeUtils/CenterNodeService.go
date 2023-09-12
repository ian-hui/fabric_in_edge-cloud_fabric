package NodeUtils

import (
	"encoding/json"
	"fmt"

	kivik "github.com/go-kivik/kivik/v4"
)

// if exist update, if not exist upload ,
func uploadFilePosition(center_nodestru Nodestructure, msg []byte) (err error) {
	var file_position_info PositionInfo
	err = json.Unmarshal(msg, &file_position_info)
	if err != nil {
		fmt.Println(err)
	}
	_, err = Getinfo(file_position_info.FileId, center_nodestru, "position_info")
	if err != nil {
		if kivik.StatusCode(err) == 404 {
			//not exist,upload
			err = UploadPostion(file_position_info, center_nodestru)
			if err != nil {
				return
			}
		} else {
			return
		}
	}
	//exist,update
	UpdatePostion(file_position_info.FileId, center_nodestru, "Position", file_position_info.Position)
	return
}

func FileReqestToCenter(center_nodestru Nodestructure, msg []byte) {
	//receive request
	var filereqstru FileRequest
	err := json.Unmarshal(msg, &filereqstru)
	if err != nil {
		fmt.Println(err)
	}
	//check if the file exist in center
	rs, err := Getinfo(filereqstru.FileId, center_nodestru, "ciphertext_info")
	//if not exist in center, forwading to the node which has the file
	if err != nil {
		if kivik.StatusCode(err) == 404 {
			fmt.Println("<------File Not found in center------>")
			//check if the position exist in center
			rs, err := Getinfo(filereqstru.FileId, center_nodestru, "position_info")
			if err != nil {
				if kivik.StatusCode(err) == 404 {
					fmt.Println("<------Position Not found in center------>")
				} else {
					fmt.Println("db.get error:", err)
				}
				return
			}
			//send request to the node which has the file
			var posinfo FilePositionInfoDTO
			err = rs.ScanDoc(&posinfo)
			if err != nil {
				fmt.Println("scandoc error: ", err)
				return
			}
			err = ProducerAsyncSending(msg, "ReceiveFileRequestFromCenter", posinfo.Position)
			if err != nil {
				fmt.Println("receive request producer sending error ", err)
				return
			}
		} else {
			fmt.Println("db.get error:", err)
		}
		return
	}
	//if exist in center send to node linking client
	var fileif FileInfo
	err = rs.ScanDoc(&fileif)
	if err != nil {
		fmt.Println("scandoc error: ", err)
		return
	}
	data := center_nodestru.PeerNodeName + " : " + fileif.Ciphertext
	//forward to the node linking to client
	datasendinginfo := DataSend2clientInfo{
		UserId:       filereqstru.UserId,
		Data:         []byte(data),
		FileId:       filereqstru.FileId,
		TransferFlag: false,
	}
	res, err := json.Marshal(datasendinginfo)
	if err != nil {
		fmt.Printf("fail to Serialization, err:%v", err)
		return
	}
	topic := "DataForwarding" //操作名
	err = ProducerAsyncSending(res, topic, filereqstru.Kafka_addr)
	if err != nil {
		fmt.Println("producer async sending err:", err)
		return
	}
}

func UploadCiText(center_nodestru Nodestructure, msg []byte) {
	//receive ciphertext
	fmt.Println("<------upload citext---->")
	var fileinfostru FileInfo
	err := json.Unmarshal(msg, &fileinfostru)
	if err != nil {
		fmt.Println(err)
	}
	PutCipherText(fileinfostru, center_nodestru)
	//change the position info of ciphertext
	//get rev first

	//modify the position in center
	err = UpdatePostion(fileinfostru.FileId, center_nodestru, "Position", []string{center_nodestru.KafkaIp})
	if err != nil {
		fmt.Println(err)
	}

}

func UploadKeyPosition(center_nodestru Nodestructure, msg []byte) {
	//receive key position
	fmt.Println("<------upload keyposition---->")
	var key_positon KeyPostionUploadInfo
	err := json.Unmarshal(msg, &key_positon)
	if err != nil {
		fmt.Println(err)
	}

	err = UpdatePostion(key_positon.FileId, center_nodestru, "GroupIps", key_positon.GroupIps)
	if err != nil {
		fmt.Println(err)
	}
}

// 接收密钥请求，根据元数据把请求转发到各个节点
// TODO:负载均衡
func KeyReqForwarding(center_nodestru Nodestructure, msg []byte) {
	//receive request
	var freq FileRequest
	err := json.Unmarshal(msg, &freq)
	if err != nil {
		fmt.Println(err)
	}
	resultSet, err := Getinfo(freq.FileId, center_nodestru, "position_info")
	if err != nil {
		if kivik.StatusCode(err) == 404 {
			fmt.Println("<------Not found in center------>")
		} else {
			fmt.Println("db.get error:", err)
		}
		return
	}
	var position_info PositionInfo
	err = resultSet.ScanDoc(&position_info)
	if err != nil {
		fmt.Println("scandoc error :", err)
		return
	}

	for _, kafka_addr := range position_info.GroupIps {
		topic := "ReceiveKeyReq" //操作名
		err = ProducerAsyncSending(msg, topic, kafka_addr)
		if err != nil {
			fmt.Println("producer async sending err:", err)
			break
		}
	}
}
