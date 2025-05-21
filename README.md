# Projeto Load Tester

Este projeto em Go é uma aplicação de linha de comando para realizar testes de carga em serviços web. Ele permite definir a quantidade total de requisições, o nível de concorrência e exibe estatísticas de desempenho ao final da execução.

## Estrutura de Diretórios

```
load-tester/
├── Dockerfile
├── go.mod
├── LICENSE
├── main.go
├── README.md
├── report/
│   └── report.go
└── tester/
    └── tester.go
```

## Requisitos

- Go 1.21 ou superior
- Docker (opcional, para execução em container)
- Acesso a um endpoint HTTP para testes

## Instalação

Clone o repositório e acesse o diretório do projeto:

```bash
git clone <url-do-repositorio>
cd load-tester
```

Baixe as dependências:

```bash
go mod tidy
```

## Execução

### Executando Localmente

```bash
go run main.go -url http://localhost:8080 -requests 100 -concurrency 10
```

### Executando com Docker

```bash
docker build -t loadtester .
docker run --rm loadtester -url http://host.docker.internal:8080 -requests 100 -concurrency 10
```

## Parâmetros

- `-url`: URL do serviço a ser testado (obrigatório)
- `-requests`: Número total de requisições (padrão: 1)
- `-concurrency`: Número de chamadas simultâneas (padrão: 1)

## Saída Esperada

```bash
Resumo da execução:
➡️  Total de requisições: 10000
✅ Requisições bem-sucedidas: 9975
❌ Requisições com falha: 25
⏱️  Tempo total: 2.314214471s
📊 Status HTTP por código:
   - 200: 9975 respostas
   - 0: 25 respostas
```

- Código `0` indica falha de conexão ou erro de rede.
- Códigos HTTP reais (como `200`, `500`, `404`) são agrupados e contados.

---

## Possíveis Erros e Soluções

### Erro: `connection refused` ou `no such host`

- Verifique se a URL fornecida está correta e acessível.
- Se estiver usando Docker, use `host.docker.internal` para acessar serviços na máquina host.

### Erro ao buildar com Docker

- Verifique se você adicionou corretamente os novos campos (`SuccessCount`, `StatusCodeCounts`) na struct `TestResult` em `tester/tester.go` e também no pacote `report`.

## Considerações Finais

Este projeto é uma ferramenta simples e útil para validar o desempenho de serviços web sob carga. Pode ser extendido com suporte a diferentes métodos HTTP, autenticação, headers personalizados, geração de relatórios em CSV ou JSON, e muito mais.