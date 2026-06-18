package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/segmentio/kafka-go"
)

type TelemetryEvent struct {
	DriverID  string    `json:"driver_id"`
	Latitude  float64   `json:"latitude"`
	Longitude float64   `json:"longitude"`
	Timestamp time.Time `json:"timestamp"`
}

const (
	kafkaURL = "localhost:9092"
	topic    = "driver-telemetry"

	// Códigos de cores para o terminal bombar no GIF
	colorReset  = "\033[0m"
	colorGreen  = "\033[32m"
	colorBlue   = "\033[34m"
	colorYellow = "\033[33m"
)

func main() {
	log.Println("🚀 Inicializando motor de processamento logístico...")
	ctx := context.Background()

	// Simulando 3 entregadores ao mesmo tempo para mostrar concorrência no GIF!
	go startDriverSimulator(ctx, "ENTREGADOR-01")
	go startDriverSimulator(ctx, "ENTREGADOR-02")
	go startDriverSimulator(ctx, "ENTREGADOR-03")

	startTelemetryProcessor(ctx)
}

func startDriverSimulator(ctx context.Context, driverID string) {
	writer := &kafka.Writer{
		Addr:     kafka.TCP(kafkaURL),
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	}
	defer writer.Close()

	// Tempo de início aleatório para não mandarem todos no mesmíssimo milissegundo
	rand.Seed(time.Now().UnixNano())

	for {
		event := TelemetryEvent{
			DriverID:  driverID,
			Latitude:  -23.550520 + (rand.Float64() * 0.005),
			Longitude: -46.633308 + (rand.Float64() * 0.005),
			Timestamp: time.Now(),
		}

		body, _ := json.Marshal(event)

		_ = writer.WriteMessages(ctx, kafka.Message{
			Key:   []byte(event.DriverID),
			Value: body,
		})

		// Log em VERDE para o envio
		fmt.Printf("%s📤 [%s] Enviou GPS (Lat: %.4f)%s\n", colorGreen, event.DriverID, event.Latitude, colorReset)

		// Cada entregador envia em um tempo ligeiramente diferente (entre 1 e 3 segundos)
		time.Sleep(time.Duration(1+rand.Intn(2)) * time.Second)
	}
}

func startTelemetryProcessor(ctx context.Context) {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{kafkaURL},
		Topic:   topic,
		GroupID: "processor-group",
	})
	defer reader.Close()

	for {
		msg, err := reader.ReadMessage(ctx)
		if err != nil {
			continue
		}

		var event TelemetryEvent
		_ = json.Unmarshal(msg.Value, &event)

		// Log em AZUL para o processamento
		fmt.Printf("%s📥 [PROCESSADOR] Lendo dados da %s...%s\n", colorBlue, event.DriverID, colorReset)

		// Regra de negócio visual em AMARELO
		if event.Latitude > -23.5490 {
			fmt.Printf("%s🔔 [ALERTA] %s acabou de entrar no raio do restaurante!%s\n", colorYellow, event.DriverID, colorReset)
		}
	}
}
