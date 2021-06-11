#Build with one stage
#FROM golang:1.16-alpine3.13 
#WORKDIR /app
#COPY . .
#RUN go build -o bookapi main.go
#EXPOSE 8000
#CMD [ "/app/bookapi" ]



#Build Multi-Stage 
FROM golang:1.16-alpine3.13 AS builder

WORKDIR /app
COPY . .

RUN go build -o bookapi main.go


#Run stage

FROM alpine:3.13
WORKDIR /app

COPY --from=builder /app/bookapi .


EXPOSE 8000

CMD [ "/app/bookapi" ]