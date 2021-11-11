# How to use proto schema with pulsar go client library

- Define a .proto file
- Generate Code using below command:

    ```shell
    protoc -I=. -I=$GOPATH/src -I=$GOPATH/src/github.com/gogo/protobuf/protobuf --gogofast_out=. --gogofast_opt=paths=source_relative your_file.proto
    ```

- Use Generated types while creating producer/consumer.
