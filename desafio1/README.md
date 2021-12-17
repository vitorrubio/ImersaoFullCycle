# imersao fullcycle desafio1
primeiro desafio das aulas da imersão full cycle

## Colinhas

### Colinha de Dockerfile

Executar um programa e depois continuar executando sem fechar o container

`CMD ./desafio && tail -f /dev/null`

---

Continuar executando depois de rodar uma imagem que não faz nada 

`CMD ["tail","-f","/dev/null"]`


### Colinha de comandos Docker

Compilar uma imagem
`docker build -t [suatag]/[nome]:latest .`

Rodar uma imagem 
`docker run --rm --name=[nome] [nome da imagem]:versao`

Remover todos os containeres
`docker rm -f $(docker ps -a -q)`

Remover todos os volumes
`docker volume rm $(docker volume ls -q)`

Deleta todas as imagens "dangling" (não tagueadas ou associadas a um container)
`docker system prune`

Deleta todas as imagens que não estão rodando
`docker system prune -a` 

### Referências 
- [https://stackoverflow.com/questions/25135897/how-to-automatically-start-a-service-when-running-a-docker-container](https://stackoverflow.com/questions/25135897/how-to-automatically-start-a-service-when-running-a-docker-container)
- [https://www.geeksforgeeks.org/docker-copy-instruction/](https://www.geeksforgeeks.org/docker-copy-instruction/)
- [https://www.ctl.io/developers/blog/post/dockerfile-add-vs-copy/](https://www.ctl.io/developers/blog/post/dockerfile-add-vs-copy/)
- [https://phoenixnap.com/kb/docker-add-vs-copy](https://phoenixnap.com/kb/docker-add-vs-copy)

-(https://docs.tibco.com/pub/mash-local/4.1.1/doc/html/docker/GUID-BD850566-5B79-4915-987E-430FC38DAAE4.html)[https://docs.tibco.com/pub/mash-local/4.1.1/doc/html/docker/GUID-BD850566-5B79-4915-987E-430FC38DAAE4.html]
-(https://www.digitalocean.com/community/tutorials/how-to-remove-docker-images-containers-and-volumes)[https://www.digitalocean.com/community/tutorials/how-to-remove-docker-images-containers-and-volumes]