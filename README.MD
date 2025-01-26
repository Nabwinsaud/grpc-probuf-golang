# This is the protobuf explorer . it contains the basic information about the protobuf and how to use it with typescript and go (mainly go examples is highly focused and tested).


## for running the protobuf file you need to install the protobuf compiler

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


## in order to run the protobuf file to respective language you need to install the respective language compiler for the protobuf. protobuf comes with default support for the c++ ,Ruby,Php,Java and python. for other languages you need to install the compiler plugin.



### for this example i will be using the typescript and go language.so typescript and go compiler extension need to be installed.

```bash
npm install -g ts-protoc-gen
npm install -g ts-proto

go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
```


## make sure after installing the compiler you can run the following command to check the version of the compiler.

```bash
protoc --version
```


## for compiling the protobuf to golang you need to run the following command.

```bash
protoc --go_out=./payment --go_opt=paths=source_relative payment.proto
```

### --go_out : this flag is used to specify the output directory for the generated go file. if you don't want directory you can use . to generate the file in the current directory.
### --go_opt=paths=source_relative : this flag is used to specify the path of the generated file relative to the source file.
### payment.proto : this is the input file for the compiler. if you want to compile multiple files once you can use wildcard character. *.proto


### for running the typescript file you need to run the following command (script modify as per you requirement).

```bash
./proto-script.sh

```