FROM instrumentisto/dep as server_builder
RUN mkdir -p /go/src/gitlab.com/phucvinh52/
WORKDIR /go/src/gitlab.com/phucvinh52/
COPY . /go/src/gitlab.com/phucvinh52/
RUN dep ensure && GOOS=linux go build -o /ldap-pass-webui main.go

FROM ubuntu:18.04
RUN apt update && apt-get install -y ca-certificates
COPY --from=server_builder /ldap-pass-webui /ldap-pass-webui
ENTRYPOINT [ "/ldap-pass-webui" ]