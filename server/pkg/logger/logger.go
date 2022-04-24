package logger

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

var (
	defaultWriter io.Writer = os.Stdout
	logger                  = &Logger{out: defaultWriter}
)

type Logger struct {
	out       io.Writer
	formatter LogFormatter
	timer     time.Time
	response  response
	Params    LogFormatterParams
}

type LoggerConfig struct {
	Formatter LogFormatter
	Output    io.Writer
}

type LogFormatter func(params LogFormatterParams) string

type LogFormatterParams struct {
	// TimeStamp shows the time after the server returns a response.
	TimeStamp time.Time
	// StatusCode is HTTP response code.
	StatusCode int
	// Latency is how much time the server cost to process a certain request.
	Latency time.Duration
	// ClientIP equals Context's ClientIP method.
	ClientIP string
	// Method is the HTTP method given to the request.
	Method string
	// Path is a path the client requests.
	Path string
	// ErrorMessage is set if error has occurred in processing the request.
	ErrorMessage string
	// BodySize is the size of the Response Body
	BodySize int
}

var defaultLogFormatter = func(param LogFormatterParams) string {
	if param.Latency > time.Minute {
		// Truncate in a golang < 1.8 safe way
		param.Latency = param.Latency - param.Latency%time.Second
	}

	return fmt.Sprintf("[LOG] %v | %3d | %13v | %15s | %-7s  %#v\n%s",
		param.TimeStamp.Format("06/01/02 - 15:04:05"),
		param.StatusCode,
		param.Latency,
		param.ClientIP,
		param.Method,
		param.Path,
		param.ErrorMessage,
	)
}

func DefaultLogger() *Logger {
	return LoggerWithConfig(LoggerConfig{})
}

func LoggerWithFormatter(f LogFormatter) *Logger {
	return LoggerWithConfig(LoggerConfig{Formatter: f})
}

func LoggerWithWriter(out io.Writer) *Logger {
	return LoggerWithConfig(LoggerConfig{Output: out})
}

func LoggerWithConfig(conf LoggerConfig) *Logger {
	formatter := conf.Formatter
	if formatter == nil {
		formatter = defaultLogFormatter
	}

	out := conf.Output
	if out == nil {
		out = defaultWriter
	}

	logger.out = out
	logger.formatter = formatter
	return logger
}

func (l *Logger) HandleRequest(r *http.Request) http.ResponseWriter {
	l.timer = time.Now()
	path := r.URL.Path
	query := r.URL.RawQuery

	if query != "" {
		path = path + "?" + query
	}

	l.Params.Path = path

	return &l.response
}

func (l *Logger) HandleResponse(r *http.Request, w http.ResponseWriter) {
	l.Params.TimeStamp = time.Now()
	l.Params.Latency = l.Params.TimeStamp.Sub(l.timer)

	l.Params.Method = r.Method
	l.Params.StatusCode = l.response.statusCode
	l.Params.BodySize = l.response.body.Len()
	// l.Params.ClientIP
	// l.Params.ErrorMessage

	// fmt.Println("Response body:", l.response.body.String())
}

func (l *Logger) LogRequestResponse() {
	fmt.Fprint(l.out, l.formatter(l.Params))
}
