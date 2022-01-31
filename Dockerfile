FROM golang:1.13 as builder1
COPY . ./src/github.com/canonical/iot-management
WORKDIR /go/src/github.com/canonical/iot-management
RUN CGO_ENABLED=1 GOOS=linux go build -a -o /go/bin/management -ldflags='-extldflags "-static"' cmd/management/main.go
RUN CGO_ENABLED=1 GOOS=linux go build -a -o /go/bin/createsuperuser -ldflags='-extldflags "-static"' cmd/createsuperuser/main.go


FROM node:8-alpine as builder2
COPY webapp .
WORKDIR /
RUN npm install
RUN npm rebuild node-sass
RUN npm run build

# Set params from the environment variables
ARG DRIVER="postgres"
ARG DATASOURCE="dbname=management sslmode=disable"
ARG HOST="management:8010"
ARG SCHEME="http"
ARG DEVICETWINAPI="http://devicetwin:8040/v1/"
ARG STOREURL="https://api.snapcraft.io/api/v1/"
ENV DRIVER="${DRIVER}"
ENV DATASOURCE="${DATASOURCE}"
ENV HOST="${HOST}"
ENV SCHEME="${SCHEME}"
ENV DEVICETWINAPI="${DEVICETWINAPI}"
ENV STOREURL="${STOREURL}"

# Copy the built applications to the docker image
FROM ubuntu:18.04
WORKDIR /root/
RUN apt-get update
RUN apt-get install -y ca-certificates
COPY --from=builder1 /go/bin/management .
COPY --from=builder1 /go/bin/createsuperuser .
COPY --from=builder1 /go/src/github.com/canonical/iot-management/static ./static
COPY --from=builder2 build/static/css ./static/css
COPY --from=builder2 build/static/js ./static/js
COPY --from=builder2 build/index.html ./static/app.html
EXPOSE 8010
ENTRYPOINT ./management
CMD ['./createsuperuser']
