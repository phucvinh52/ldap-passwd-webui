FROM instrumentisto/dep as server_builder
RUN mkdir -p /go/src/gitlab.com/phucvinh52/
WORKDIR /go/src/gitlab.com/phucvinh52/
COPY Gopkg.toml /go/src/gitlab.com/phucvinh52/
COPY main.go /go/src/gitlab.com/phucvinh52/
RUN dep ensure 
COPY . /go/src/gitlab.com/phucvinh52/
RUN GOOS=linux go build -o /ldap-pass-webui main.go

FROM ubuntu:18.04
RUN apt update && apt-get install -y ca-certificates
RUN mkdir -p /myapp
COPY --from=server_builder /ldap-pass-webui /ldap-pass-webui
COPY . /myapp
WORKDIR /myapp
ENTRYPOINT [ "/ldap-pass-webui" ]