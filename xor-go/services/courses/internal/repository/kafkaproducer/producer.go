package kafkaproducer

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/juju/zaputil/zapctx"
	"github.com/segmentio/kafka-go"
	global "go.opentelemetry.io/otel"
	"go.uber.org/zap"
	"xor-go/services/courses/internal/domain"
	"xor-go/services/courses/internal/service/adapters"
)

var _ adapters.KafkaClient = &KafkaProducer{}

type KafkaProducer struct {
	producer *kafka.Writer
}

func NewKafkaProducer(cfg *Config) *KafkaProducer {
	writer := kafka.Writer{
		Addr:     kafka.TCP(cfg.Broker),
		Topic:    cfg.Topic,
		Balancer: &kafka.LeastBytes{},
	}

	return &KafkaProducer{
		producer: &writer,
	}
}

func (kp *KafkaProducer) SendMessage(ctx context.Context, message string) error {
	logger := zapctx.Logger(ctx)

	tr := global.Tracer(domain.ServiceName)
	newCtx, span := tr.Start(ctx, "courses/repository/kafka.SendMessage")
	defer span.End()

	messageJSON, err := json.Marshal(message)
	if err != nil {
		logger.Error("failed to marshal Command to message:", zap.Error(err))
		// TODO AppError
		return fmt.Errorf("failed to marshal Command to message: %w", domain.ErrInternal)
	}

	err = kp.SendMessageWithKaKafka(newCtx, messageJSON)
	if err != nil {
		logger.Error("failed write message to kafka:", zap.Error(err))
		// TODO AppError
		return fmt.Errorf("failed to send message to Kafka: %w", domain.ErrInternal)
	}
	return nil
}

func (kp *KafkaProducer) SendMessageWithKaKafka(ctx context.Context, message []byte) error {
	return kp.producer.WriteMessages(ctx, kafka.Message{
		Value: message,
	})
}
