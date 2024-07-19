package producers

import (
	"encoding/json"
	"project-root/exceptions"
	"project-root/structs"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sqs"
)

var ProximaEtapaQueueURL = "http://localhost:4566/000000000000/acionamento_proxima_etapa_cadastro"

func ProduzirProximaEtapa(payload structs.ProximaEtapaPayload) error {
	messageBody, err := json.Marshal(payload)
	if err != nil {
		return exceptions.NewAppError("Failed to marshal next step payload", err)
	}

	_, err = SQSClient.SendMessage(&sqs.SendMessageInput{
		MessageBody: aws.String(string(messageBody)),
		QueueUrl:    aws.String(ProximaEtapaQueueURL),
	})
	if err != nil {
		return exceptions.NewAppError("Failed to send message to SQS", err)
	}

	return nil
}
