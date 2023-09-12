package NodeUtils

// //kivik_client 连接池
// type kivik_client_map struct {
// 	*sync.Map
// }
// type kafka_producer_map struct {
// 	*sync.Map
// }

//get一个kivikclient的pool，如果没有就创建一个新的pool，new方法已经定义好了，返回一个client
// func (p *kivik_client_map) GetKivikClient(kivik_addr string) (*kivik.Client, error) {
// 	pool, ok := p.Load(kivik_addr)
// 	if !ok {
// 		// 如果不存在，创建一个新的 clientPool
// 		newpool := &sync.Pool{
// 			New: func() interface{} {
// 				// 创建一个新的 kivik.client 对象
// 				client, err := kivik.New("couch", kivik_addr)
// 				if err != nil {
// 					panic(err)
// 				}
// 				return client
// 			},
// 		}
// 		p.Store(kivik_addr, newpool)
// 		return newpool.Get().(*kivik.Client), nil
// 	}
// 	return pool.(*sync.Pool).Get().(*kivik.Client), nil
// }

// func (p *kivik_client_map) ReleaseKivikClient(kivik_addr string, client *kivik.Client) error {
// 	pool, ok := p.Load(kivik_addr)
// 	if ok {
// 		pool.(*sync.Pool).Put(client)
// 		return nil
// 	} else {
// 		return fmt.Errorf("pool doesnot exist")
// 	}
// }

// func (p *kafka_producer_map) GetProducerClient(kafka_addr string) (*sarama.AsyncProducer, error) {
// 	pool, ok := p.Load(kafka_addr)
// 	if !ok {
// 		newpool := &sync.Pool{
// 			New: func() interface{} {
// 				//创建kafka producer
// 				producerconfig := sarama.NewConfig()
// 				producerconfig.Producer.RequiredAcks = sarama.WaitForAll          // 发送完数据需要leader和follow都确认
// 				producerconfig.Producer.Partitioner = sarama.NewRandomPartitioner // 新选出一个partition，只有一个分区，所以只能是0
// 				producerconfig.Producer.Return.Successes = true                   // 成功交付的消息将在success channel返回
// 				client, err := sarama.NewAsyncProducer([]string{kafka_addr}, producerconfig)
// 				if err != nil {
// 					panic(err)
// 				}
// 				return &client
// 			},
// 		}
// 		p.Store(kafka_addr, newpool)
// 		return newpool.Get().(*sarama.AsyncProducer), nil
// 	}
// 	return pool.(*sync.Pool).Get().(*sarama.AsyncProducer), nil
// }

// func (p *kafka_producer_map) ReleaseProducerClient(kafka_addr string, client *sarama.AsyncProducer) error {
// 	pool, ok := p.Load(kafka_addr)
// 	if ok {
// 		pool.(*sync.Pool).Put(client)
// 		return nil
// 	} else {
// 		return fmt.Errorf("pool doesnot exist")
// 	}
// }
