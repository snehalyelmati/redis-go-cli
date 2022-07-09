# Redis Go CLI

Easy to use CLI application to view data/configuration in Redis.

Options available:
- `-data`: Prints all the existing data on Redis. 
- `-config`: Prints all the current Redis config.
- `-testReadWrite`: Tests the Redis connection by setting and deleting sample data based on the `-count` flag.
- `-insertSampleData`: Inserts sample data into Redis based on the `-count` flag.
- `-count`: Number of records to insert and/or delete(default is 5 records).
- `-deleteAllData`: Deletes all existing data in the Redis instance.

Settings:
- `-hostname`: Sets the hostname of the Redis instance to connect(default is localhost).
- `-port`: Sets the port of the Redis instance to connect(default is 6379).
- `-password`: Sets the password of the Redis instance to connect(default is empty).

## How to use
1. To build the executable use the following command,
```
$ go build -o bin/rgcli_version
```

2. Run the executable file with the required flags,
```
$ rgcli_version -hostname=redis-xyz.com -port=4999 -password=1234 -data -config
```
