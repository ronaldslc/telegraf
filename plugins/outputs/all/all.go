package all

import (
	_ "github.com/ronaldslc/telegraf/plugins/outputs/amon"
	_ "github.com/ronaldslc/telegraf/plugins/outputs/amqp"
	_ "github.com/ronaldslc/telegraf/plugins/outputs/datadog"
	_ "github.com/ronaldslc/telegraf/plugins/outputs/influxdb"
	_ "github.com/ronaldslc/telegraf/plugins/outputs/kafka"
	_ "github.com/ronaldslc/telegraf/plugins/outputs/kinesis"
	_ "github.com/ronaldslc/telegraf/plugins/outputs/librato"
	_ "github.com/ronaldslc/telegraf/plugins/outputs/mqtt"
	_ "github.com/ronaldslc/telegraf/plugins/outputs/nsq"
	_ "github.com/ronaldslc/telegraf/plugins/outputs/opentsdb"
	_ "github.com/ronaldslc/telegraf/plugins/outputs/prometheus_client"
	_ "github.com/ronaldslc/telegraf/plugins/outputs/riemann"
)
