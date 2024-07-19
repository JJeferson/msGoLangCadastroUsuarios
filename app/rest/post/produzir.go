package post

import (
	"encoding/json"
	"net/http"
	"project-root/dynamo"
	"project-root/exceptions"
	"project-root/producers"
	"project-root/structs"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

var SQSClient = sqs.New(session.Must(session.NewSession()))

func Produzir(w http.ResponseWriter, r *http.Request) {
	var usuario structs.Usuario
	if err := json.NewDecoder(r.Body).Decode(&usuario); err != nil {
		exceptions.HandleError(w, "Invalid request payload", err)
		return
	}

	// Produzir para a fila SQS "acionamento_lambda_cadastro_usuario"
	if err := producers.ProduzirCadastroUsuario(usuario); err != nil {
		exceptions.HandleError(w, "Failed to produce to cadastro_usuario queue", err)
		return
	}

	// Consumir a fila SQS "acionamento_lambda_cadastro_usuario"
	messages, err := SQSClient.ReceiveMessage(&sqs.ReceiveMessageInput{
		QueueUrl:            aws.String(producers.QueueURL),
		MaxNumberOfMessages: aws.Int64(1),
	})
	if err != nil {
		exceptions.HandleError(w, "Failed to receive message from SQS", err)
		return
	}

	if len(messages.Messages) > 0 {
		// Gravar no DynamoDB
		idDynamo, err := dynamo.GravarUsuario(messages.Messages[0])
		if err != nil {
			exceptions.HandleError(w, "Failed to save to DynamoDB", err)
			return
		}

		// Produzir para a fila SQS "acionamento_proxima_etapa_cadastro"
		proximaEtapaPayload := structs.ProximaEtapaPayload{
			IdDynamo: idDynamo,
			Nome:     usuario.Nome,
			Sucesso:  true,
		}
		if err := producers.ProduzirProximaEtapa(proximaEtapaPayload); err != nil {
			exceptions.HandleError(w, "Failed to produce to proxima_etapa queue", err)
			return
		}
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "Success"}`))
}
