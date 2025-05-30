package utils

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
)

const (
	// Default log format will output [INFO]: 2006-01-02T15:04:05Z07:00 - Log message
	defaultLogFormat       = "[%time%] [%lvl%] %caller% %fields% %msg%\n"
	defaultTimestampFormat = time.RFC3339
)

// Formatter implements logrus.Formatter interface.
type Formatter struct {
	// Timestamp format
	TimestampFormat string
	// Available standard keys: time, msg, lvl
	// Also can include custom fields but limited to strings.
	// All of fields need to be wrapped inside %% i.e %time% %msg%
	LogFormat string
}

// Format building log message.
func (f *Formatter) Format(entry *logrus.Entry) ([]byte, error) {
	output := f.LogFormat
	if output == "" {
		output = defaultLogFormat
	}

	timestampFormat := f.TimestampFormat
	if timestampFormat == "" {
		timestampFormat = defaultTimestampFormat
	}

	output = strings.Replace(output, "%time%", entry.Time.Format(timestampFormat), 1)

	output = strings.Replace(output, "%msg%", entry.Message, 1)

	level := strings.ToUpper(entry.Level.String())
	output = strings.Replace(output, "%lvl%", level, 1)

	caller := ""

	if entry.Caller != nil {
		m1 := regexp.MustCompile(`.*/`)
		s := m1.ReplaceAllString(entry.Caller.Func.Name(), "")
		s2 := m1.ReplaceAllString(entry.Caller.File, "") + ":" + strconv.Itoa(entry.Caller.Line)
		caller = "[" + s + "][" + s2 + "]"
	}

	output = strings.Replace(output, "%caller%", caller, 1)

	fields := ""
	for _, val := range entry.Data {
		switch v := val.(type) {
		case string:
			fields += "[" + v + "]"
			// output = strings.Replace(output, "%"+k+"%", v, 1)
		case int:
			s := strconv.Itoa(v)
			fields += "[" + s + "]"
			// output = strings.Replace(output, "%"+k+"%", s, 1)
		case bool:
			s := strconv.FormatBool(v)
			fields += "[" + s + "]"
			// output = strings.Replace(output, "%"+k+"%", s, 1)
		case error:
			s := v.Error()
			fields += "[" + s + "]"
		default:
			s := fmt.Sprint(v)
			fields += "[" + s + "]"
		}
	}

	output = strings.Replace(output, "%fields%", fields, 1)

	return []byte(output), nil
}
