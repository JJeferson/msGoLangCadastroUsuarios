# msGoLangCadastroUsuarios
## Teste de Microserviço.

<b>
Cadastro de usuarios num banco dynamo. </br>
Testavel com LocalStack </br>
Integração com SQS e Dynamo. </br>
Acionamento por httpRest Post. </br>
</b>
<b>
Fluxo:</br>
1)Via Collection feito acionamento. </br>
2)Aplicação vai salvar numa fila SQS. </br>
3)Aplicação vai consumir a fila e salvar no dynamo. </br>
4)Depois ela vai produzir para a outra fila avisando que deu certo. </br>
</b>

## Como testar:
Instale o LocalStack.
```
docker pull localstack/localstack

```
Execute-o:
```
docker run --rm -it -p 4566:4566 -p 4571:4571 localstack/localstack
```
Agora Inicie as filas SQS:
```
aws --endpoint-url=http://localhost:4566 sqs create-queue --queue-name acionamento_lambda_cadastro_usuario


aws --endpoint-url=http://localhost:4566 sqs create-queue --queue-name acionamento_proxima_etapa_cadastro

```
Agora o Dynamo:
```
aws --endpoint-url=http://localhost:4566 dynamodb create-table \
    --table-name cadastro_cliente \
    --attribute-definitions AttributeName=id,AttributeType=S \
    --key-schema AttributeName=id,KeyType=HASH \
    --provisioned-throughput ReadCapacityUnits=5,WriteCapacityUnits=5

```
Então execute a aplicação. </br>
Dentro da pasta app rode:
```
go mod init project-root
go get github.com/aws/aws-sdk-go
go get github.com/gorilla/mux
go run main.go
```
Usando um collection produza para o endpoint localhost:3030/produzir [Post]

Para verificar se deu certo no gitbash rode o comando:
```
aws --endpoint-url=http://localhost:4566 sqs receive-message --queue-url http://localhost:4566/000000000000/acionamento_proxima_etapa_cadastro
```

Ali dentro deve ter a resposta com true dizendo que deu tudo certo.