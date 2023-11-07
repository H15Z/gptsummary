# build stage
FROM golang:buster AS build-env
# RUN apk --no-cache add build-base git 

ADD . /src
RUN cd /src && go build -o gptsummary

# final stage
FROM debian:buster-slim

########## SETUP TLS CERTS ############
RUN apt-get update && apt-get install -y ca-certificates
########## SETUP TLS CERTS ############

WORKDIR /app

#copy binary to final container
COPY --from=build-env /src/gptsummary /app/

#copy data files
COPY data /app/data

#copy config
COPY .cmd.yml /app/.cmd.yml

ENTRYPOINT ./gptsummary stream -c 10 -t 3
