package ifapm

import (
	"context"
	"github.com/sirupsen/logrus"
)

type Log struct {
}

func init() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
}

var Logger = &Log{}

func (l *Log) Info(ctx context.Context, action string, kv map[string]interface{}) {
	kv["action"] = action
	logrus.WithFields(kv).Infoln(ctx)
}

func (l *Log) Error(ctx context.Context, action string, kv map[string]interface{}, err error) {
	kv["action"] = action
	logrus.WithFields(kv).Errorln(err)
}

func (l *Log) Debug(ctx context.Context, action string, kv map[string]interface{}) {
	kv["action"] = action
	logrus.WithFields(kv).Debugln(ctx)
}

func (l *Log) Warn(ctx context.Context, action string, kv map[string]interface{}) {
	kv["action"] = action
	logrus.WithFields(kv).Warnln(ctx)
}
