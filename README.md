# CEP Weather API 2

## Visão Geral
Este projeto consiste em dois microserviços em Go, organizados seguindo a Clean Architecture e instrumentados com OpenTelemetry + Zipkin para tracing distribuído:

- **Serviço A**: Responsável por receber um CEP via POST, validar seu formato e repassar para o Serviço B.
- **Serviço B**: Recebe o CEP, consulta o ViaCEP para obter a cidade e busca a temperatura atual na API WeatherAPI, retornando os valores em Celsius, Fahrenheit e Kelvin.

## Arquitetura
Cada serviço possui:
- Um `main.go` responsável pela inicialização (configuração de tracer OTEL, middleware Gin e start do servidor).
- Diretórios `cmd`, `internal` e `cfg` para separar rotas, lógica de negócio, clientes HTTP, modelos de dados e carregamento de variáveis de ambiente.
- Instrumentação OTEL com middleware Gin para capturar spans de entrada e saída.


## Configuração
Os dois serviços e os componentes de observabilidade (OpenTelemetry Collector e Zipkin) são orquestrados via Docker Compose. Variáveis principais:
- `SERVICE_A` e `SERVICE_B`: portas e URLs internas.
- `FREEWEATHER_API_KEY`: chave da WeatherAPI.
- `OTEL_ENDPOINT`: endpoint do collector (geralmente `otel-collector:4317`).
- `OTEL_SERVICE_NAME`: nome do serviço para identificação no Zipkin.

## Como rodar
1. Certifique-se de ter Docker e Docker Compose instalados.
2. No arquivo `docker-compose.yml`, ajuste `WEATHER_API_KEY` em `service-b`.
3. Execute no terminal, no diretório raiz do projeto:
   ```bash
   docker-compose up --build
   ```
4. Aguarde até que todos os containers (collector, Zipkin, service-a, service-b) estejam em estado `healthy` e sem erros.

## Endpoints
- **Serviço A**  
  - `POST /cep`  
    - Envia JSON com `{ "cep": "XXXXXXXX" }`.  
    - Retorna 422 se formato inválido, ou os dados de cidade e temperatura obtidos pelo Serviço B.

- **Serviço B**  
  - `GET /weather/:cep`  
    - Recebe CEP como parâmetro de rota.  
    - Retorna 422 se formato inválido, 404 se não encontrá-lo no ViaCEP, ou JSON com cidade e temperaturas (C, F, K).

## Observabilidade
- Acesse Zipkin em `http://localhost:9411/zipkin/`.
- No UI do Zipkin:
  1. Selecione o **Service Name** (service-a ou service-b).
  2. Ajuste o **Time Range** para cobrir suas requisições.
  3. Clique em **Find Traces** para listar traces e spans.
- Cada operação-chave (validação no Serviço A, lookup de CEP, fetch de temperatura) aparece como span, permitindo análise de latência e falhas.

## Swagger

- Acesse a documentação Swagger do Serviço A em `http://localhost:8080/swagger/index.html`.
- Acesse a documentação Swagger do Serviço B em `http://localhost:9090/swagger/index.html`.
