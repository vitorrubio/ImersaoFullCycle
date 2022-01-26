# Notas da Imersão Full Cycle 5

## Minhas notas pessoais da imersão do curso Full Cycle

notas da imersão do curso Full Cycle Aulas 1 e 2

---

### Aulas
- Aula 1: Criação de processador de transações em GO
- Aula 2: Conexão do processador de transações com kafka


### Colinha
- docker-compose up -d --build   => faz o build da imagem usando o Dockerfile e roda em background
- docker exec -it [nome_do_container] bash  => chama o bash do container



# Imersão Full Stack & FullCycle 5.0 - Fincycle - Backend de processamento das transações feito com Golang

## Descrição

Repositório do back-end de processamento das transações feito com Golang

**Importante**: A aplicação do Apache Kafka deve estar rodando primeiro.

## Rodar a aplicação

### Configurar /etc/hosts

A comunicação entre as aplicações se dá de forma direta através da rede da máquina.
Para isto é necessário configurar um endereços que todos os containers Docker consigam acessar.

Acrescente no seu /etc/hosts (para Windows o caminho é C:\Windows\system32\drivers\etc\hosts):
```
127.0.0.1 host.docker.internal
```
Em todos os sistemas operacionais é necessário abrir o programa para editar o *hosts* como Administrator da máquina ou root.

Execute os comandos:

```
docker-compose up
```

Use o Apache Kafka para produzir e consumir mensagens para testar a aplicação.

Quer configurar um ambiente de desenvolvimento produtivo? Veja o vídeo: [https://www.youtube.com/watch?v=nTlM4sd-W3E](https://www.youtube.com/watch?v=nTlM4sd-W3E) 

### Para Windows 

Siga o guia rápido de instalação: [https://github.com/codeedu/wsl2-docker-quickstart](https://github.com/codeedu/wsl2-docker-quickstart) 

