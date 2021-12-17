FROM golang:1.17

WORKDIR /go/src

RUN apt-get update && apt-get install -y sqlite3

#para incluir o source e compilar dentro do container e rodar junto com tail -f

#COPY . /go/src/desafiosrc

#RUN go build -o ./desafio ./desafiosrc/main.go

#CMD ./desafio && tail -f /dev/null


#para copiar o binario e rodar só ele

COPY ./desafio /go/src/desafio

COPY ./banco.db /go/src/banco.db

CMD ./desafio && tail -f /dev/null
#CMD ["tail","-f","/dev/null"] 