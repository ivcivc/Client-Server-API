## Desafio: Cliente-Servidor em Go para Cotação do Dólar

**Objetivo:**

Desenvolva um sistema cliente-servidor em Go que:
* **Consulte** uma API externa para obter a cotação do dólar em tempo real.
* **Armazene** as cotações históricas em um banco de dados SQLite.
* **Retorne** a cotação atual para um cliente.
* **Utilize** contextos (context) para gerenciar timeouts nas operações.
* **Persista** a última cotação em um arquivo de texto.

**Estrutura do Projeto:**
.
├── client.go
├── server.go
├── cotacao.txt
└── cotacoes.db

**Funcionamento:**

1. **Servidor (server.go):**
   * Inicia um servidor HTTP na porta 8080.
   * Exporta um endpoint `/cotacao` para receber requisições do cliente.
   * Consome a API https://economia.awesomeapi.com.br/json/last/USD-BRL para obter a cotação do dólar.
   * Armazena a cotação no banco de dados SQLite utilizando contextos para garantir a persistência dentro de 10ms.
   * Retorna a cotação atual em formato JSON para o cliente.

2. **Cliente (client.go):**
   * Realiza uma requisição HTTP GET para o servidor.
   * Recebe a cotação do dólar e extrai o valor "bid".
   * Salva a cotação em um arquivo de texto `cotacao.txt`.
   * Utiliza contextos para garantir que a requisição ao servidor seja concluída em 300ms.

**Requisitos:**

* **Go:** Versão compatível com as bibliotecas utilizadas.
* **Banco de dados:** SQLite configurado para armazenar as cotações.
* **Pacotes:** `net/http`, `context`, `database/sql`, `encoding/json`, `os`.

**Como executar:**

1. **Clonar o repositório:**
   ```bash
   git clone https://Client-Server-API.git

2. **Instalar as dependências:**
   go mod tidy
   
3. **Rodar:**
   go run server.go
   
4. **Rodar:**
   go run client.go


   
