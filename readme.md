# PAYMENT-READER

Serviços responsável por escutar estar eventos vindos da GLIC, persistir na collection de histórico e produzir um evento para o compilador do EDI.

## Requisitos

- Go 1.21.3

## Bibliotecas utilizadas

- Viper
- Cobra
- MongoDB
- Gin Gonic

## Modo de usar

- Alterar o arquivo `.env.example` para `.env`
- Adicionar as configurações necessárias do banco e do kafka

Para obter ajuda:

```bash
make help
```

Para subir a api:
```bash
make api
```

Para subir o worker:
```bash
make wrk
```
