# Redis Go CLI

Easy to use, configurable CLI application to test connectivity, view data/configuration in Redis built with Go.

Options available:

- `-data`: Prints all the existing data on Redis.
- `-config`: Prints all the current Redis config.
- `-testReadWrite`: Tests the Redis connection by setting and deleting sample data based on the `-count` flag.
- `-insertSampleData`: Inserts sample data into Redis based on the `-count` flag.
- `-deleteAllData`: Deletes all existing data in the Redis instance.
- `-count`: Number of records to insert and/or delete(default is 5 records). Applicable to `-insertSampleData` and `-testReadWrite`.

Settings:

- `-hostname`: Sets the hostname of the Redis instance to connect(default is localhost).
- `-port`: Sets the port of the Redis instance to connect(default is 6379).
- `-username`: Sets the username of the Redis instance to connect(default is empty).
- `-password`: Sets the password of the Redis instance to connect(default is empty).

## How to use

1. To build the executable use the following command,

```
$ go build -o bin/rgcli_version
```

2. Run the executable file with the required flags,

```
// to test connectivity
$ rgcli_version

// to connect with custom parameters like hostname, port, etc.
$ rgcli_version -hostname=redis-xyz.com -port=4999 -username=redis123 -password=1234

// to get existing data
$ rgcli_version -data

// to get existing data with custom parameters
$ rgcli_version -hostname=redis-xyz.com -port=4999 -username=redis123 -password=1234 -data

// to get data and config
$ rgcli_version -data -config

// inserts sample data based on the count argument
$ rgcli_version -insertSampleData -count=10

// to test read and write to Redis with the count argument
$ rgcli_version -testReadWrite -count=10
```

3. (Optional) To compile for a specific OS use the `GOOS` env variable, if `direnv` is used modify the `.envrc` file.

```
// to compile binary for linux
export GOOS=linux
```
