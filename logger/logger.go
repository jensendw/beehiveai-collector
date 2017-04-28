package logger

import "github.com/op/go-logging"

var format = logging.MustStringFormatter(
	`%{color}%{time:15:04:05.000} %{shortfunc} ▶ %{level:.4s} %{id:03x}%{color:reset} %{message}`,
)

var Logger = logging.MustGetLogger("Beehiveai-Collector")
