# Load Balancer Go

- Projeto de estudo criado com o objetivo de aprofundar conhecimentos em Go. 

## Padrões de Projeto e Arquitetura

Este projeto implementa um **load balancer** em Go utilizando o **Strategy Pattern** e **Simple Factory**.  
As estratégias possíveis são:  

- **Round-Robin**  
- **Least Connections**  
- **Weighted Round-Robin**  

O LB implementado segue aberto para extensão (novos algoritmos) e fechada para modificação.

![Diagrama de Classes](assets/lb.jpg)


## Integração Contínua (CI)

O projeto possui **pipeline de CI** que roda em cada push ou pull request na branch `main`.  
Ele realiza:

- Checkout do código  
- Configuração do Go 1.22 com cache  
- Verificação da versão do Go  
- Execução de linter (`golangci-lint`)  
- Build do projeto  

Isso garante qualidade de código e detecção precoce de erros.

## Docker

O Dockerfile utiliza **duas etapas** para otimizar tamanho e segurança da imagem final:

1. **Build** – compila o binário Go:
   - Base: Alpine com Go 1.22  
   - Define diretório de trabalho `/app`  
   - Copia todo o código-fonte  
   - Compila e gera o binário  

2. **Prod** – cria a imagem de produção leve:
   - Base: Alpine limpa (descarta código-fonte e cache de build)  
   - Copia apenas o binário da etapa de build  
   - Expõe a aplicação na porta 8080  
   - Inicia o container executando o binário  

**Resultado:** imagem final contendo apenas o binário, reduzindo tamanho total.

## Inicialização

O load balancer pode ser inicializado a partir de um arquivo **JSON de configuração**, que define:  

- Porta da aplicação  
- Pool de backends  
- Estratégia de balanceamento  

### Exemplo 1: Least Connections

```json
{
  "balancingStrategy": "least-connections",
  "backend": [
    {"url": "http://service1.example.com"},
    {"url": "http://service2.example.com"},
    {"url": "http://service3.example.com"}
  ]
}
````

### Exemplo 2: Weighted Round Robin

```json
{
  "balancingStrategy": "weighted-round-robin",
  "backend": [
    {"url": "http://service1.example.com", "weight": 4},
    {"url": "http://service2.example.com", "weight": 2},
    {"url": "http://service3.example.com", "weight": 1}
  ]
}
```

### Exemplo 3: Round Robin

```json
{
  "balancingStrategy": "round-robin",
  "backend": [
    {"url": "http://service1.example.com"},
    {"url": "http://service2.example.com"},
    {"url": "http://service3.example.com"}
  ]
}
```
