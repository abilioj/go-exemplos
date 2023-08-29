# -inicia o projeto
go mod init app-exemplo

# pra roda
go run main.go

# pra compilar 
go build

# pra compilar especifico pra windows
GOOS=windows go build
-------------------------------------------------------------------------------------------------
# Instalando plugin go
go get -u github.com/golang/protobuf/protoc-gen-go

# Gerando codigo direto chamando protoc
protoc --go_out=. ./user/user.proto

# Gerando codigo usando go generate
go generate
-------------------------------------------------------------------------------------------------
docker build -t abiliojgf/myservergo:lastest .
docker run --rm -p 8080:8080 abiliojgf/myservergo:lastest