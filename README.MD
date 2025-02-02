## ProtoBuf and grpc (Typescript & Go) - Go implementation.


for running the protobuf file you need to install the protobuf compiler

### for mac
```bash
brew install protobuf

```

### for linux
```bash
sudo apt-get update
sudo apt-get install protobuf-compiler

or manually you can download using curl and then extract the file and add the bin directory to the path.
```

### for windows 

```bash
 figure it out yourself

```


in order to run the protobuf file to respective language you need to install the respective language compiler for the protobuf. protobuf comes with default support for the c++ ,ruby,php,java,rust and python. for other languages you need to install the compiler plugin.



for this example i will be using the typescript and go language.so typescript and go compiler extension need to be installed.

```bash
npm install -g ts-protoc-gen
npm install -g ts-proto

go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
```


make sure after installing the compiler you can run the following command to check the version of the compiler.

```bash
protoc --version
```


 for compiling the protobuf for golang you need to run the following command.

```bash
protoc --go_out=./payment --go_opt=paths=source_relative payment.proto
```

 --go_out : this flag is used to specify the output directory for the generated go file. if you don't want directory you can use . to generate the file in the current directory.

--go_opt=paths=source_relative : this flag is used to specify the path of the generated file relative to the source file.

payment.proto : this is the input file for the compiler. if you want to compile multiple files once you can use wildcard character. *.proto


for running the typescript file you need to run the following command (script modify as per you requirement).

```bash
./proto-script.sh

```

run the task file (a alternative to makefile)

## install task file in mac
```bash
brew install go-task

```

## install task file in linux
```bash
sudo snap install task --classic
  or temporarily
curl -sL https://taskfile.dev/install.sh | sh
```

## install task file in windows
```bash
choco install go-task
```
[for more info](https://taskfile.dev/#/installation)

for go you need to run the following command to run the go file.

```bash
go mod tidy
```


> [!NOTE]
> make sure you have protoc-gen-go & proto-gen-go-grpc installed in > your system. if not then install it using the following command.

[for more info](https://grpc.io/docs/languages/go/quickstart/)


```bash
## run the grpc server
```bash
go run server.go
```

## run the grpc client
```bash
go run client.go

``` 