Desafio: Cliente-Servidor em Go com Cotação do Dólar

Objetivo:
  Criar um sistema cliente-servidor em Go que:

Consuma uma API externa para obter a cotação do dólar.
Armazene as cotações em um banco de dados SQLite.
Retorne a cotação mais recente para um cliente.
Utilize contextos para gerenciar timeouts.
Persista a cotação atual em um arquivo de texto.
Estrutura do Projeto:

client.go:
  Realiza uma requisição HTTP GET para o servidor.
  Recebe a cotação do dólar.
  Salva a cotação em um arquivo de texto.
  Utiliza um contexto com timeout de 300ms.
  
server.go:
  Inicia um servidor HTTP na porta 8080.
  Consome a API https://economia.awesomeapi.com.br/json/last/USD-BRL.
  Armazena a cotação no banco de dados SQLite usando um contexto com timeout de 10ms.
  Retorna a cotação mais recente em formato JSON para o cliente.
  Utiliza um contexto com timeout de 200ms para chamar a API.
  Exporta um endpoint /cotacao.

  
Requisitos:

  Go: Instalar o Go e configurar o ambiente de desenvolvimento.
  Banco de dados: Criar um banco de dados SQLite para armazenar as cotações.
  Pacotes: Utilizar os pacotes net/http, context, database/sql, encoding/json e os.
  
Instruções de Uso:

  Clonar o repositório:
    Bash
    git clone https://seu-repositorio.git
    Use o código com cuidado.
  
  Instalar as dependências:
    Bash
    go mod tidy
    Use o código com cuidado.
  
  Criar o banco de dados:
    Opção 1: Criar um banco de dados SQLite vazio com o nome especificado no código.
    Opção 2: Criar uma tabela no banco de dados para armazenar as cotações (id, cotacao, data).
      Executar o servidor:
      Bash
      go run server.go
      Use o código com cuidado.
  
  Executar o cliente:
    Bash
    go run client.go
    Use o código com cuidado.

Estrutura de Arquivos (Exemplo):

.
├── client.go
├── server.go
├── cotacao.txt
└── cotacoes.db
Considerações:

Tratamento de erros: Implementar um tratamento de erros robusto para lidar com falhas na conexão com a API, no banco de dados e em outras operações.
Log: Adicionar logs para facilitar a depuração e o monitoramento do sistema.
Testes: Criar testes unitários para garantir a qualidade do código.
Melhorias:

Cache: Implementar um cache para evitar chamadas desnecessárias à API.
Escalabilidade: Considerar o uso de um pool de conexões com o banco de dados para melhorar o desempenho.

Segurança: Validar as entradas e proteger o sistema contra ataques.

Contribuições: Contribuições são bem-vindas! Abra um issue para discutir novas funcionalidades ou melhorias.

Licença:
[Especificar a licença utilizada]

Autor:
[Ivan Oliveira]

Este README fornece uma base sólida para o seu projeto. Adapte-o para refletir as especificidades do seu código e adicionar mais detalhes conforme necessário.
