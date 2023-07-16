# account-details
Comprehensive project exploring end to end usage of protobuff with go.

## Notes

### `Windows`

- I recommend you use powershell (try to update: [see](https://github.com/PowerShell/PowerShell/releases)) for following this course.
- I recommend you use [Chocolatey](https://chocolatey.org/) as package installer (see [Install](https://chocolatey.org/install))


### Build

#### `Prerequisites`

```shell
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

#### `Linux/MacOS`

```shell
go mod tidy
make
```

#### `Windows - Chocolatey`
```shell
choco install make
go mod tidy
make
```

#### `Windows - Without Chocolatey`

```shell
protoc -Iproto --go_opt=module=github.com/Clement-Jean/proto-go-course --go_out=. proto/*.proto

go mod tidy
go build -o account-details.exe .
```

## Run

```
./account-details
```

or

```
./account-details.exe
```
