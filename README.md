<h1 align="center">
    <img style="width: 50%" alt="iDev" src="https://user-images.githubusercontent.com/18368947/88479901-f2f4ea80-cf28-11ea-866e-377b8e2d7a49.png" />
</h1>

## 🔖  Sobre o desfio

- iDev API, desenvolvido na liguagem [Go](https://golang.org/ "Go"), no qual tem por função a leitura de aquivo [JSON](https://www.json.org/json-en.html "JSON") e efetuar os cálculos de média, moda e tendência de utilização, para os itens descritos no arquivo de leitura.

## 💻 Tecnologias
- <img width="20px" src="https://img.icons8.com/color/2x/golang.png" /> [Go](https://golang.org/ "Go")
- <img width="20px" src="https://img.icons8.com/material/2x/harambe-the-gorilla.png" /> [Mux](https://www.gorillatoolkit.org/pkg/mux "Mux") - Roteamento
- <img width="20px" src="https://insomnia.rest/images/icon-small.png" /> [Insomnia](https://insomnia.rest/ "Insomnia") - Efetuando testes das requisições

## ▶️ Getting Started

- Passo 1: executar a instalação do [Go](https://golang.org/ "Go")
- Passo 2: git clone do projeto [iDev](https://github.com/rafaelsanzio/iDev "iDev")

```bash
# Navegando até a pasta do projeto
$ cd iDev
# Instalando todas as depêndencias necessárias
$ go get .
# Starting o backend da aplicação
$ go run main.go
```

## 🔨 Casos de Teste
 - Requisição de informações do projeto
 	- **Rota:** http://localhost:8080/
 	- **Response**
```json
{
    "enterprise": "iDev Soluções",
    "about": "API developed to read a JSON file showing (avarage, trend and usage trend) of servers",
    "projectLink": "github.com/rafaelsanzio/iDev",
    "creator": "Rafael Sanzio"
}
```

- Requisição de informações de um determinado server
	- **Rota:** http://localhost:8080/server-info/:serverName
	- **Param:** (serverName) - nome do servidor que deseja buscar as informações
	- **Response:**
```json
{
  "serverName": "server0",
  "CPU": {
    "avarage": "0.50 %",
    "mode": [
      -1
    ],
    "usageTrend": "0.26 %"
  },
  "memory": {
    "avarage": "4.96 GB",
    "mode": [
      4.313038509465582,
      4.118241302627751
    ],
    "usageTrend": "5.55 GB"
  },
  "disk": {
    "avarage": "25.77 GB",
    "mode": [
      -1
    ],
    "usageTrend": "17.88 GB"
  },
  "occurrences": 283
}
```

**Obeservações**: Quando a moda retorna (-1) quer dizer que tal atributo não tem uma moda definida, caso retorne mais de um, indica que o atributo tem mais de uma moda.

<h1 align="center">
	<img alt="Desafio iDev API" src="https://insomnia.rest/images/run.svg" />
</h1>


## ㊗️ Considerações
- Projeto desenvolvido por:

	- <img width="20px" src="https://img.icons8.com/fluent/96/github.png" /> [Rafael Sanzio](https://github.com/rafaelsanzio "Rafael Sanzio")
 
	- <img width="20px" src="https://img.icons8.com/color/2x/linkedin.png" /> [Rafael Sanzio](https://www.linkedin.com/in/rafael-sanzio-012778143/ "Rafael Sanzio")
