package producers

import (
	"encoding/json"
	"project-root/exceptions"
	"project-root/structs"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sqs"
)

var SQSClient *sqs.SQS
var QueueURL = "http://localhost:4566/000000000000/acionamento_lambda_cadastro_usuario"

func ProduzirCadastroUsuario(usuario structs.Usuario) error {
	messageBody, err := json.Marshal(usuario)
	if err != nil {
		return exceptions.NewAppError("Failed to marshal user payload", err)
	}

	_, err = SQSClient.SendMessage(&sqs.SendMessageInput{
		MessageBody: aws.String(string(messageBody)),
		QueueUrl:    aws.String(QueueURL),
	})
	if err != nil {
		return exceptions.NewAppError("Failed to send message to SQS", err)
	}

	return nil
}
