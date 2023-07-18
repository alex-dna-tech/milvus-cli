# Golang Milvus CLI
## Used libraries
- [Cobra - Modern CLI framework](https://github.com/spf13/cobra)
- [Viper - Module to save configuration](https://github.com/spf13/viper)
- [PromptUI - Interactive prompt for CLI](https://github.com/spf13/viper)

## Install
- Get executable file from [releases](https://github.com/alex-dna-tech/milvus-cli/releases)
- Use [docker container](https://hub.docker.com/r/dnat4/milvus-cli)

## TODO (use golang [sdk](https://github.com/milvus-io/milvus-sdk-go) or [protobuf](https://github.com/milvus-io/milvus-proto/tree/master/go-api) client)
[Python example)(https://milvus.io/docs/cli_commands.md)
### Calc command
- [ ] calc distance
- [ ] calc mkts_from_hybridts
- [ ] calc mkts_from_unixtime
- [ ] calc hybridts_to_unixtime

### Create command
- [ ] create user
- [ ] create alias
- [ ] create collection
- [ ] create partition
- [ ] create index

### Delete command
- [ ] delete user
- [ ] delete alias
- [ ] delete collection
- [ ] delete entities
- [ ] delete partition
- [ ] delete index

### Describe command
- [ ] describe collection
- [ ] describe partition
- [ ] describe index

### List command
- [ ] list users
- [ ] list collections
- [ ] list indexes
- [ ] list partitions

### Show command
- [ ] show connection
- [ ] show index_progress
- [ ] show loading_progress
- [ ] show query_segment

### Other command
- [ ] load
- [ ] load_balance
- [ ] query
- [ ] release
- [ ] search
- [ ] import
- [x] version
- [ ] clear
- ~~connect~~
- ~~exit~~
- ~~help~~

## Combine [Milvus backup](https://github.com/zilliztech/milvus-backup) REST API example

- Use [ToolJet](https://github.com/ToolJet/ToolJet) for WebUI builder for REST API
- Use [Attu Node Example](https://github.com/zilliztech/attu) to interact
