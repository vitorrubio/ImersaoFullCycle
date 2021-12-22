# Imersão Curso FullCycle
minhas notas pessoais do curso full cycle

## Vídeos e recursos

### Link do curso e comunidade
- [Link do curso](https://imersao.fullcycle.com.br/evento/)
- [Comunidade Imersão](https://discord.com/invite/NRmPwv44RW)
- [Comunidade do Curso](https://discord.com/channels/831167847698923570/831167847698923574)


### Esquenta
- [Ambiente produtivo com WSL 2 e Docker](https://www.youtube.com/watch?v=nTlM4sd-W3E)
- [Docker do Zero na prática](https://www.youtube.com/watch?v=zyO-XHBzi-c)
- [Desenvolvimento de APIs com Nest.js](https://www.youtube.com/watch?v=CmJGmNkY6FU)
- [Desenvolvimento de aplicações escaláveis com Go Lang](https://www.youtube.com/watch?v=nTmZlzwTErM)

- [Server Side Rendering com Next.js](https://www.youtube.com/watch?v=T9IAlagTcJk)
- [Apache kafka do Zero](https://www.youtube.com/watch?v=2uCEww7x4rs)
- [Kafka com Nest.js no mundo dos Microsserviço](https://www.youtube.com/watch?v=gMHn-itg3zk)
- [Clean Architecture](https://www.youtube.com/watch?v=YaGVURjB33I)
- [Next.js e API Routes para integrações com outras APIs](https://www.youtube.com/watch?v=mhnoux6J2g0)
- [Orquestrando containers com Kubernetes do Zero](https://www.youtube.com/watch?v=tlqWYH13nOg)


### Conteudos adicionais
- [Como funciona o RabbitMQ](https://www.youtube.com/watch?v=2YWHtbZJ0QI)
- [Sistemas Monolíticos vs Microsserviços](https://www.youtube.com/watch?v=UgmRiqL0sbA)
- [3 erros que você comete com DDD](https://www.youtube.com/watch?v=mKqGrivdsuQ)
- [Teste de carga com Kubernetes](https://www.youtube.com/watch?v=AYAE1GeO6Cs)
- [Introdução ao Apache Kafka](https://www.youtube.com/watch?v=o5yviW6QSrE&list=PL5aY_NrL1rjt_AZxj11kQjiTNLGg4ZaZA&index=3)

### Lives 
- [Mercado Livre: Tecnologias, Microsserviços e Arquitetura a prova de falhas](https://www.youtube.com/watch?v=MiukifE3fSc)
- [Do monolito aos microsserviços no Itaú: Erros e acertos](https://www.youtube.com/watch?v=xwQn9O1xWFw&feature=emb_imp_woyt)
- [API Gateways na LuizaLabs](https://www.youtube.com/watch?v=xwQn9O1xWFw&feature=emb_imp_woyt)
- [Case Full Cycle: Desafios da construção de um microsserviço](https://www.youtube.com/watch?v=3aTb8a_F4dQ&feature=emb_imp_woyt)
- [RabbitMQ na Prática](https://www.youtube.com/watch?v=3aTb8a_F4dQ)
- [LuizaLabs: Documentando arquiteturas na prática com C4 model](https://www.youtube.com/watch?v=aJZPKyElP6A&feature=emb_imp_woyt)
- [Testes automatizados do jeito certo](https://www.youtube.com/watch?v=M46XisdGXJA&feature=emb_imp_woyt)
- [Docker: 3 Formas de colocar sua aplicação em produção](https://www.youtube.com/watch?v=2Rc7S1ixrl0)
- [Observabilidade com Elastic, Prometheus e Grafana](https://www.youtube.com/watch?v=MdkkPYZD86Y)

### Aulas 
- [Aula 01 - Criando uma aplicação GO para processar transações](https://imersao.fullcycle.com.br/aula/aula-1/)
- [Aula 02 - Apache Kafka e comunicação assíncrona](https://imersao.fullcycle.com.br/aula/apache-kafka-e-comunicacao-assincrona/) - conecta a aplicação GO da aula 1 com kafka
- [Aula 03 - microsserviço nestjs e comunicação assíncrona com o kafka](https://imersao.fullcycle.com.br/aula/microsservico-nestjs-apis-e-publicacao-assincrona/)
- [Aula 04 - frontend no nextjs](https://imersao.fullcycle.com.br/aula/frontend-moderno-e-realtime-com-nextjs/)
- [Aula 05 - orquestração de containers com kubernetes](https://imersao.fullcycle.com.br/aula/orquestracao-de-containers-com-kubernetes/)





## Diretórios e arquivos desse repositírio


### Esquenta

- esquentafc-golang 
Aplicação em golang que simula um processador de transações de um e-commerce e grava os dados em um banco sqlite3. Usa conceitos de Clean Architeture

- esquentafc-nestjs
Aplicação / API em nestjs.

- notas-k8s
Notas de kubernetes

### Aulas
- aula1-golang 
Conteúdo das aulas 1 e 2 de golang onde criamos o programa processador de transações em GO que vai receber as mensagens do kafka e gravar no banco sqlite3.

- aula3-nestjs
Aplicação / API em nestjs que vai receber uma transação do frontend e enviar para o kafka. Do kafka ela vai para o processador em golang.

- aula4-nextjs
Frontend em nextjs que vai mandar ou listar os dados de transações usando a api em nestjs.


### Desafios
- desafio1 
programa em golang que salve dados em um banco sqlite3

- desafio2
api em nestjs que processe uma transação e mande para um banco de dados sqlite3 ou mysql

- desafio3 
fazer os arquivos deployment.yaml e service.yaml da aula de kubernetes