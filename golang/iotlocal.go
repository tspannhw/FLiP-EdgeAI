// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package main

import (
        "context"
        "fmt"
        "log"
        "time"

        "github.com/apache/pulsar-client-go/pulsar"
	"github.com/hpcloud/tail"
)

// customized by Tim Spann
// 2021 Sept 13
func main() {

	client, err := pulsar.NewClient(pulsar.ClientOptions{
        URL:               "pulsar://192.168.1.181:6650",
        OperationTimeout:  30 * time.Second,
        ConnectionTimeout: 30 * time.Second,
        })

	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	producer, err := client.CreateProducer(pulsar.ProducerOptions{
		Topic: "nvidia-sensor",
	})

	if err != nil {
    		log.Fatal(err)
	}

	t, err := tail.TailFile("/opt/demo/logs/sensors.log", tail.Config{Follow:true})
        for line := range t.Lines {
		if msgId, err := producer.Send(context.Background(), &pulsar.ProducerMessage{
			Payload: []byte(line.Text),
		}); err != nil {
			log.Fatal(err)
		} else {
			fmt.Printf("jetsonsensor:Published message: %v-%s \n", msgId,line.Text)
		}
	}
}
