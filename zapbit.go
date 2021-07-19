package zapbit

import (
	"io"

	"github.com/zainul/ark/log"
	"github.com/zainul/zapbit/kafka"
	"github.com/zainul/zapbit/rabbitmq"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Writer struct {
	eng io.Writer
}

type Option *Writer

func WithKafka(cfg kafka.Config, o ...Option) {
	kafkaEngine, err := kafka.NewWriter(cfg)
	if err != nil {
		log.Error("Failed to init kafka")
	}

	for _, opt := range o {
		opt.eng = kafkaEngine
	}
}

func WithRabbit(cfg rabbitmq.Config, o ...Option) {
	rabbitMQ, err := rabbitmq.NewWriter(cfg, "zapbit-log")
	if err != nil {
		log.Error("Failed to init rabbit")
	}

	for _, opt := range o {
		opt.eng = rabbitMQ
	}
}

// GetCore is get zapcore of log engine
func GetCore(w io.Writer) zapcore.Core {
	highPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.ErrorLevel
	})
	lowPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl < zapcore.ErrorLevel
	})
	topicLog := zapcore.AddSync(w)
	jsonEnc := zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())

	return zapcore.NewTee(
		zapcore.NewCore(jsonEnc, topicLog, highPriority),
		zapcore.NewCore(jsonEnc, topicLog, lowPriority),
	)
}