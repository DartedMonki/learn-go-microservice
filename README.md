# Installation
> Download and install Go
> https://go.dev/doc/install

> Install Protobuf
> https://grpc.io/docs/protoc-installation/

> Run go mod
```bash
cd {service}

go mod init github.com/dartedmonki/{service}

go work init ./{service}
```
> Run Makefile on ./common
```bash
make
```

> To run the services locally using air

```bash
go install github.com/cosmtrek/air@latest
```

> Run inside the service folder

```bash
air
```