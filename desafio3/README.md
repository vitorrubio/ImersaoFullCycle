# Desafio 3

## Cluster Kubernetes com kind e imagem de nginx

### Colinhas

- Criar um cluster do kubernetes dando nome e especificando versão da imagem
`kind create cluster --name desafio3 --image kindest/node:v1.17.17 --config cluster-config.yaml`

- ver informações do cluster
`kubectl cluster-info --context kind-desafio3`

- listar os nodes 
`kubectl get nodes`

- aplicar um arquivo de configuração
`kubectl apply -f deployment.yaml`

- listar os pods
`kubectl get pods`

- listar os services
`kubectl get services`

- listar os replicaset
`kubectl get rs`

- listar os deployments
`kubectl get deployments`

- redirecionar porta de pod
`kubectl port-forward pod/[nome pod] 8080:80`

- redirecionar porta de service
`kubectl port-forward svc/[nome pod] 8080:80`
tipo: `kubectl port-forward svc/desafio3-service 80:80`

mas ... no wsl2 sem root, tem que mandar um sudo -E porque não tem permissão pra servir nas portas baixar reservadas < 1024:

`sudo -E kubectl port-forward svc/desafio3-service 80:80`