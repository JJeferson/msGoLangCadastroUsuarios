package dynamo

import (
	"encoding/json"
	"fmt"
	"project-root/exceptions"
	"project-root/structs"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/sqs"
)

var DynamoClient = dynamodb.New(session.Must(session.NewSession()))
var TableName = "cadastro_cliente"

func GravarUsuario(message *sqs.Message) (string, error) {
	var usuario structs.Usuario
	if err := json.Unmarshal([]byte(*message.Body), &usuario); err != nil {
		return "", exceptions.NewAppError("Failed to unmarshal message body", err)
	}

	id := fmt.Sprintf("%v", message.MessageId)

	item := map[string]*dynamodb.AttributeValue{
		"id": {
			S: aws.String(id),
		},
		"nome": {
			S: aws.String(usuario.Nome),
		},
		"idade": {
			N: aws.String(fmt.Sprintf("%d", usuario.Idade)),
		},
	}

	for i, endereco := range usuario.Enderecos {
		item[fmt.Sprintf("endereco%d", i)] = &dynamodb.AttributeValue{
			S: aws.String(endereco.NomeEndereco),
		}
	}

	input := &dynamodb.PutItemInput{
		Item:      item,
		TableName: aws.String(TableName),
	}

	_, err := DynamoClient.PutItem(input)
	if err != nil {
		return "", exceptions.NewAppError("Failed to put item in DynamoDB", err)
	}

	return id, nil
}
