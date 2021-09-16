# FLiP-EdgeAI

Edge AI with Jetson Nano + Sensors


# 


# FLiP-SQL

Streaming Analytics with Apache Pulsar and Apache Flink SQL


## IoT Data Publishing into MQTT


I installed the Eclipse Paho MQTT Pyton3 client for NVIDIA XAVIER NX, Raspberry Pi 4 and Mac PowerBook.


```
pip3 install paho-mqtt
```

## Run this yourself

* You could run standalone https://pulsar.apache.org/docs/en/standalone/  https://ci.apache.org/projects/flink/flink-docs-release-1.13//docs/try-flink/local_installation/
* Run on StreamNative in the Cloud easy https://streamnative.io/en/cloud/managed/


## Links

* More Source Code https://github.com/tspannhw/FLiP-IoT
* Slides https://noti.st/tspannhw/pjnmzO/continuous-sql-with-apache-streaming-flank-and-flip
* https://github.com/tspannhw/rpi-picamera-mqtt-nifi
* https://hub.streamnative.io/protocol-handlers/mop/0.2.0
* https://github.com/eclipse/paho.mqtt-spy/releases/tag/1.0.0
* https://www.hivemq.com/mqtt-toolbox/
* https://pulsar.apache.org/docs/en/reference-cli-tools/
* https://github.com/streamnative/mop/releases/tag/v2.8.0.10
* https://pulsar.apache.org/docs/en/reference-cli-tools/
* https://pulsar.apache.org/docs/en/admin-api-topics/#create
* http://www.steves-internet-guide.com/into-mqtt-python-client/
* https://pulsar.apache.org/docs/en/sql-getting-started/
* https://github.com/morsapaes/flink-sql-pulsar
* https://flink.apache.org/2021/01/07/pulsar-flink-connector-270.html
* https://github.com/morsapaes/flink-sql-pulsar/
* https://github.com/streamnative/pulsar-flink
* https://ci.apache.org/projects/flink/flink-docs-release-1.13/docs/dev/table/sqlclient/
* https://streamnative.io/en/blog/release/2021-04-20-flink-sql-on-streamnative-cloud
* https://www.youtube.com/watch?v=0BxXjEqoJlU
* https://www.youtube.com/watch?v=9ojajM7Zt0M&t=2105s
* https://docs.streamnative.io/cloud/stable/compute/flink-sql
* https://docs.streamnative.io/cloud/stable/compute/flink-sql-cookbook


## Local Pulsar Cluster REST End Points

* http://localhost:8080/admin/v2/persistent/public/default
* http://localhost:8080/admin/v2/persistent/public/default/nvidia-sensor/stats
* http://localhost:8080/admin/v2/persistent/public/default/nvidia-sensor/internalStats
* http://localhost:8080/admin/v2/persistent/public/default/nvidia-sensor/subscription/test-sub/position/10



## Create topic
```

bin/pulsar-admin topics create persistent://public/default/nvidia-sensor

```


##  Verify creation
```
bin/pulsar-admin topics list public/default
```

## Consume Data

```

bin/pulsar-client consume "persistent://public/default/nvidia-sensor-partition-0" -s "nano2gbgo" -n 0

```

## Example Run

```
----- got message -----
key:[null], properties:[], content:{"ir": "259.0", "id": "20210914142805_cfc3c4ef-6639-46cb-bb16-1fa7ee6424d9", "end": "1631629727.25", "uuid": "nano_uuid_xmm_20210914142805", "lux": "0", "gputemp": "16.5", "cputemp": "16.5", "te": "41.2743530273", "systemtime": "09/14/2021 10:28:47", "hum": "63.78", "memory": 21.8, "gas": "29830.0", "pressure": "1015.67", "host": "nano2gb-desktop", "diskusage": "32617.2 MB", "ipaddress": "192.168.1.217", "macaddress": "1c:bf:ce:1a:7f:a0", "temp": "16.39", "uv": "0.03", "gputempf": "62.0", "host_name": "nano2gb-desktop", "runtime": "41.0", "cpu": 0.0, "cputempf": "62.0"}
```

## Clean up when done

Delete topics
```
bin/pulsar-admin topics delete persistent://public/default/nvidia-sensor

```


## jetson sensor table

```
CREATE TABLE default_catalog.default_database.scada2
(
  uuid STRING, systemtime STRING,
  ipaddress STRING, host STRING, host_name STRING, 
  macaddress STRING, endtime STRING, runtime STRING, starttime STRING,
 cpu_temp STRING, diskusage STRING, id STRING, 
 temperature STRING, adjtemp STRING, adjtempf STRING, temperaturef STRING, 
  proximity INT, gasko STRING,
  publishTime TIMESTAMP(3) METADATA,
  WATERMARK FOR publishTime AS publishTime - INTERVAL '5' SECOND
) WITH (
  'connector' = 'pulsar',
  'topic' = 'persistent://public/default/mqtt-2',
  'value.format' = 'json',
  'service-url' = 'pulsar://localhost:6650',
  'admin-url' = 'http://localhost:8080',
  'scan.startup.mode' = 'earliest'
);
```



## Example Architecture


![Apache Pulsar + Apache Flink](https://streamnative.io/uploads/images/blogs/flinksql/4.png)

## Running Flink Cluster

```
./bin/start-cluster.sh
./bin/sql-client.sh embedded --library /Users/tspann/Documents/servers/flink-1.13.2/sqllib -e /Users/tspann/Documents/servers/flink-1.13.2/sql-client.yaml
```

