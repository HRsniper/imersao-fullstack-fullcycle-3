# Imersão Full Stack & FullCycle 3.0 - Codebank

## Descrição

Repositório do CodeBack Go (Backend)

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

Quando parar os containers do Go, lembre-se antes de rodar o `docker-compose up`, rodar o `docker-compose down` para limpar o armazenamento, senão lançará erro ao subir novamente.
