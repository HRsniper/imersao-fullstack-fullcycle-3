# Imersão Full Stack & FullCycle 3.0 - Codebank - Back-end da loja

## Descrição

Repositório do back-end das faturas (Code Invoice) feito com Nest.js

**Importante**: A aplicação do Apache Kafka, Golang (codebank) deve estar rodando primeiro.

## Configurar _/etc/hosts_

A comunicação entre as aplicações se dá de forma direta através da rede da máquina.
Para isto é necessário configurar um endereços que todos os containers Docker consigam acessar.

Acrescente no seu _/etc/hosts_ (para Windows o caminho é _C:\Windows\system32\drivers\etc\hosts_):

```
127.0.0.1 host.docker.internal
```

Em todos os sistemas operacionais é necessário abrir o programa para editar o _hosts_ como Administrator da máquina ou root.

## Rodar a aplicação

Execute os comandos:

```
docker-compose up
```

Quando parar os containers do Nest, lembre-se antes de rodar o `docker-compose up`, rodar o `docker-compose down` para limpar o armazenamento, senão lançará erro ao subir novamente.

Acessar http://localhost:3002/credit-cards.
