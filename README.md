# How to use proto schema with pulsar go client library

- Define a .proto file
- Generate Code using below command:

    ```shell
    protoc -I=. -I=$GOPATH/src -I=$GOPATH/src/github.com/gogo/protobuf/protobuf --gogofast_out=. --gogofast_opt=paths=source_relative your_file.proto
    ```

- Attach schema to the topic

    1. Generate a JSON string of your schema.

    2. Define schema definition in a json file. See example schema definition below:

        ```json
        {
            "type": "PROTOBUF",
            "schema": "{\"fields\":[{\"default\":\"\",\"name\":\"name\",\"type\":\"string\"},{\"default\":0,\"name\":\"age\",\"type\":\"int\"}],\"name\":\"Person\",\"namespace\":\"person\",\"type\":\"record\"}",
            "properties": {}
        }
        ```

    3. Upload schema to the pulsar topic

        ```shell
        pulsar-admin schemas upload --filename < schemafile > < topicname >
        ```

- Use Generated types while creating producer/consumer. See producer/producer_test.go for examples.
