# Desafio Multithreading - APIs de CEP

Este projeto tem como objetivo demonstrar o uso de **multithreading** para buscar dados de CEP em duas APIs diferentes simultaneamente, retornando a resposta da API que for mais rápida. Além disso, é implementado um timeout de 1 segundo para garantir que o programa não espere indefinidamente por uma resposta.

## Tecnologias Utilizadas

- [Go](https://golang.org/) - Linguagem de programação utilizada no projeto
- [BrasilAPI](https://brasilapi.com.br/) - API utilizada para buscar informações de CEP
- [ViaCEP](https://viacep.com.br/) - Outra API utilizada para buscar informações de CEP

## Estrutura do Projeto

O projeto segue uma estrutura simples de pastas e arquivos. 
```plaintext
/multithreading-challenge 
│ 
├── /dto # Arquivos de definição das structs (DTOs) para deserialização das respostas das APIs 
    │ 
    ├── brasil_api.go # Structs de resposta da BrasilAPI 
    │ 
    └── via_cep.go # Structs de resposta da ViaCEP 
│ 
├── main.go # Arquivo principal que executa o desafio 
└── README.md # Este arquivo
```


## Descrição do Desafio

Neste desafio, o programa realiza duas requisições simultâneas para as APIs:

- **BrasilAPI**: `https://brasilapi.com.br/api/cep/v1/{cep}`
- **ViaCEP**: `http://viacep.com.br/ws/{cep}/json/`

A requisição que retornar primeiro será exibida no terminal. A outra resposta será descartada. Se ambas as requisições demorarem mais de 1 segundo, o programa exibe uma mensagem de timeout.

## Requisitos

- Exibir o resultado da requisição mais rápida no terminal, incluindo os dados do endereço e a origem da API (BrasilAPI ou ViaCEP).
- Se o tempo de resposta for maior que 1 segundo, exibir uma mensagem de timeout.
- Tratar erros de conexão e problemas de resposta, como resposta não sendo um JSON válido.

## Como Executar o Projeto

### Pré-requisitos

- Instale o [Go](https://golang.org/dl/)

### Passos para Execução

1. Clone este repositório:

```bash
   git clone https://github.com/leonardo-gmuller/multithreading-challenge.git
```

2. Navegue até o diretório do projeto:
```bash
   cd multithreading-challenge
```

3. Execute o projeto:
```bash
   go run main.go
```

## Estrutura do Código

O arquivo main.go contém duas funções principais que realizam as requisições HTTP:

- `SearchCepBrasilApi`(ch chan dto.BrasilApiResponse): Faz a requisição para a BrasilAPI e envia a resposta para o canal.

- `SearchCepViaCep`(ch chan dto.ViaCepResponse): Faz a requisição para a ViaCEP e envia a resposta para o canal.

Ambas as funções são executadas em goroutines para permitir a execução simultânea. O select no main aguarda a resposta mais rápida ou um timeout de 1 segundo.

### DTOs
As structs usadas para deserializar as respostas das APIs estão no pacote `dto`, separadas em dois arquivos:

- `brasil_api.go`: Define a struct BrasilApiResponse que corresponde à resposta da BrasilAPI.

- `via_cep.go`: Define a struct ViaCepResponse que corresponde à resposta da ViaCEP.

### Exemplo de Saída
Se a **BrasilAPI** for a mais rápida:
```json
{
  "cep": "01153-000",
  "state": "SP",
  "city": "São Paulo",
  "neighborhood": "Barra Funda",
  "street": "Rua Vitorino Carmilo",
  "service": "brasilapi"
}
```

Se a **ViaCEP** for a mais rápida:
```json
{
  "cep": "01153-000",
  "logradouro": "Rua Vitorino Carmilo",
  "complemento": "",
  "bairro": "Barra Funda",
  "localidade": "São Paulo",
  "uf": "SP",
  "ibge": "3550308",
  "gia": "1004",
  "ddd": "11",
  "siafi": "7107"
}
```

Se ambas demorarem mais de 1 segundo:
```bash
timeout
```