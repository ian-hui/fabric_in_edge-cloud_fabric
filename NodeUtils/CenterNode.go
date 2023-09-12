package NodeUtils

import (
	"fabric-go-sdk/clients"
	"fmt"
	"sync"

	"github.com/Shopify/sarama"
)

// 初始化center-consumer
//
//topic:uploadposition,ReceiveRequest,UploadCiText
func InitCenterNode(topics []string, center_nodestru Nodestructure) {
	//init couchdb
	clients.InitCouchdb(center_nodestru.Couchdb_addr)
	//init kafka producer
	clients.InitProducer(center_nodestru.KafkaIp)
	//create db in couchdb
	Create_position_info(center_nodestru)
	Create_ciphertext_info(center_nodestru)
	var wg sync.WaitGroup
	wg.Add(5)
	consumer1, err := clients.InitConsumer(center_nodestru.KafkaIp)
	if err != nil {
		fmt.Printf("fail to start consumer, err:%v\n", err)
		return
	}
	go consumeUploadPosition(consumer1, center_nodestru, &wg)
	go consumeUploadCiText(consumer1, center_nodestru, &wg)
	go consumeKeyReqForwarding(consumer1, center_nodestru, &wg)
	fmt.Println(center_nodestru.KafkaIp, "init center-consumer1 begin")
	consumer2, err := clients.InitConsumer(center_nodestru.KafkaIp)
	if err != nil {
		fmt.Printf("fail to start consumer, err:%v\n", err)
		return
	}
	go consumeFileReqestToCenter(consumer2, center_nodestru, &wg)
	go consumeUploadKeyPosition(consumer2, center_nodestru, &wg)
	fmt.Println(center_nodestru.KafkaIp, "init center-consumer2 begin")
	wg.Wait()

}

func consumeUploadPosition(consumer sarama.Consumer, center_nodestru Nodestructure, wg *sync.WaitGroup) {
	partitonConsumer, err := consumer.ConsumePartition("uploadposition", 0, sarama.OffsetNewest)
	if err != nil {
		panic(err)
	}
	wg.Done()
	for {
		select {
		case msg := <-partitonConsumer.Messages():
			uploadFilePosition(center_nodestru, msg.Value)
		case <-partitonConsumer.Errors():
			fmt.Println("consumerUploadPosition error")
		}
	}
}

func consumeUploadCiText(consumer sarama.Consumer, center_nodestru Nodestructure, wg *sync.WaitGroup) {
	partitonConsumer, err := consumer.ConsumePartition("UploadCiText", 0, sarama.OffsetNewest)
	if err != nil {
		panic(err)
	}
	wg.Done()
	for {
		select {
		case msg := <-partitonConsumer.Messages():
			UploadCiText(center_nodestru, msg.Value)
		case <-partitonConsumer.Errors():
			fmt.Println("consumerUploadCiText error")
		}
	}
}

func consumeKeyReqForwarding(consumer sarama.Consumer, center_nodestru Nodestructure, wg *sync.WaitGroup) {
	partitonConsumer, err := consumer.ConsumePartition("KeyReqForwarding", 0, sarama.OffsetNewest)
	if err != nil {
		panic(err)
	}
	wg.Done()
	for {
		select {
		case msg := <-partitonConsumer.Messages():
			KeyReqForwarding(center_nodestru, msg.Value)
		case <-partitonConsumer.Errors():
			fmt.Println("consumerkeyreqforwarding error")
		}
	}
}

func consumeFileReqestToCenter(consumer sarama.Consumer, center_nodestru Nodestructure, wg *sync.WaitGroup) {
	partitonConsumer, err := consumer.ConsumePartition("FileReqestToCenter", 0, sarama.OffsetNewest)
	if err != nil {
		panic(err)
	}
	wg.Done()
	for {
		select {
		case msg := <-partitonConsumer.Messages():
			FileReqestToCenter(center_nodestru, msg.Value)
		case <-partitonConsumer.Errors():
			fmt.Println("consumeFileReqestToCenter error")
		}
	}
}

func consumeUploadKeyPosition(consumer sarama.Consumer, center_nodestru Nodestructure, wg *sync.WaitGroup) {
	partitonConsumer, err := consumer.ConsumePartition("UploadKeyPosition", 0, sarama.OffsetNewest)
	if err != nil {
		panic(err)
	}
	wg.Done()
	for {
		select {
		case msg := <-partitonConsumer.Messages():
			UploadKeyPosition(center_nodestru, msg.Value)
		case <-partitonConsumer.Errors():
			fmt.Println("consumeUploadKeyPosition error")
		}
	}
}
