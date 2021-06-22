# Authorizer API

## Arquitetura:
  - Golang como opção de linguagem para se ter algo desempenho no processamento das autorizacões;
  - Código clean organizado de forma elegante e com alto desempenho, organizado em três principais pacotes no projeto:
    - `application`: Define a camada de entrada da aplicação, subdividida em `controller`, nesse caso apenas com o operationController para receber a entrada das mensages via arquivo ou stdin (Obs. O main faz um pouco desse papel também). `presentation`, onde ficam os dtos de entrada e saída. E `service`, onde fica a lógica de negócio da aplicação, como a criação de accounts e a autorização de uma transação na conta.
    - `domain`: Define os modelos, ou seja, caso o projeto evolua, eles seriam as entidades que poderíamos persistir numa base de dados por exemplo. Além disso, eles representam os objetos que são manipulamos na camada de serviço.
    - `infraestructure`: Além de um pacote para manter a organização e manutenção do projeto, pode servir para adicionarmos outros estruturas ao projeto como clients, bigqueue, exceções customizadas, swagger (caso evolua para um REST), utilities, log, etc.
    - Os testes unitários seguem a conversão `nomeDoArquivo_test.go` e deve ficar dentro do mesmo pacote. Fizemos testes das principais funcionalidades e validaçoes do fluxo de autorização.
    - Está disponível um arquivo `operations.txt` no diretório principal do projeto que pode ser usado para entrada das operações.

## Instruções de compilação e execução:
  - Certifique-se de ter o Golang instalado e configurado [Go Install](https://golang.org/doc/install);
  - Dentro do diretório principal do projeto execute o comando `go run main.go < operations.txt` caso queira usar o arquivo como entrada das operações. Caso prefira passar as operações via stdin, execute o comando `go run main.go` com todas as operações que deseje, por exemplo: 
```yaml
{"account": {"active-card": true, "available-limit": 1000}}
{"transaction": {"merchant": "Nike", "amount": 800, "time": "2019-02-13T11:01:01.000Z"}}
{"transaction": {"merchant": "Uber", "amount": 80, "time": "2019-02-13T11:01:11.000Z"}}
```
...e digite `comand` + `enter` para encerrar a execução. Em ambos os caso lhe será apresentado a saída padrão:
```yaml
{"account":{"active-card":true,"available-limit":1000},"violations":[]}
{"account":{"active-card":true,"available-limit":200},"violations":[]}
{"account":{"active-card":true,"available-limit":120},"violations":[]}
```
