# FLiP-EdgeAI

Edge AI with Jetson Nano + Sensors

# Author

** Timothy Spann **


# FLiP-SQL

Streaming Analytics with Apache Pulsar and Apache Flink SQL from data received from NVIDIA Jetson Nano with Environmental Sensor running Python and Golang apps streaming data over Pulsar


## Run this yourself

* You could run standalone 
* https://pulsar.apache.org/docs/en/standalone/  
* https://ci.apache.org/projects/flink/flink-docs-release-1.13//docs/try-flink/local_installation/
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
* http://localhost:8080/admin/v2/persistent/public/default/nvidia-sensor-partition-0/stats
* http://localhost:8080/admin/v2/persistent/public/default/nvidia-sensor-partition-0/internalStats
* http://localhost:8080/admin/v2/persistent/public/default/nvidia-sensor-partition-0/subscription/test-sub/position/10


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
CREATE TABLE default_catalog.default_database.nvidiasensor
(
  `id` STRING, uuid STRING, ir STRING,
  `end` STRING, lux STRING, gputemp STRING, 
  cputemp STRING, `te` STRING, systemtime STRING, hum STRING,
 memory STRING, gas STRING, pressure STRING, 
 `host` STRING, diskusage STRING, ipaddress STRING, macaddress STRING, 
  gputempf STRING, host_name STRING,
    `runtime` STRING, cpu STRING,cputempf STRING,
  publishTime TIMESTAMP(3) METADATA,
  WATERMARK FOR publishTime AS publishTime - INTERVAL '5' SECOND
) WITH (
  'connector' = 'pulsar',
  'topic' = 'persistent://public/default/nvidia-sensor',
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

## Sensors on NVIDIA JETSON NANO 2GB

* https://www.waveshare.com/environment-sensor-for-jetson-nano.htm 
* https://www.waveshare.com/wiki/Environment_Sensor_for_Jetson_Nano


With Jetson Nano motherboard, it can collect the temperature and humidity, atmospheric pressure, ambient light, VOC, infrared, ultraviolet and other sensor data of the surrounding environment.


* Onboard TSL25911FN digital ambient light sensor infrared
* BME280 temperature, humidity and pressure sensor
* ICM20948 motion attitude sensor
* Si1145 red UVt sensor ultraviolet
* SGP40 gas sensor VOC
* 1.3-inch OLED onboard

Download and install Environment_sensor_fot_jetson_nano.   Includes sensors for Python2, only Python2.   

## Run python2 app

```

python2 sensors.py


```

## consume

```
bin/pulsar-client consume "persistent://public/default/mqtt-2" -s mqtt-reader
```

## admin

```
bin/pulsar-admin topics stats-internal persistent://public/default/mqtt-2
bin/pulsar-admin topics peek-messages --count 5 --subscription mqtt-reader persistent://public/default/mqtt-2
bin/pulsar-admin topics subscriptions persistent://public/default/mqtt-2
```


