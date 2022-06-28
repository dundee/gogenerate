package log

import (
	"github.com/sirupsen/logrus"
)

type Fields map[string]interface{}

type LogMarshaler interface {
	MarshalLog() Fields
}

func Info(msg string, args ...LogMarshaler) {
	f := logrus.Fields{}

	for _, arg := range args {
		for k, v := range arg.MarshalLog() {
			f[k] = v
		}
	}

	logrus.WithFields(f).Info(msg)
}
