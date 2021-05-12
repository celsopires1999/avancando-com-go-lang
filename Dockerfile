FROM golang:latest

WORKDIR /go/src

#COPY . .

ENV PATH="/go/bin:${PATH}"

RUN apt-get update

# Inicializando o módulo
#RUN go mod init celsopires/learning-golang 

#Gera o binário - somente para mandar para produção
#RUN go install 

CMD ["tail", "-f", "/dev/null"]