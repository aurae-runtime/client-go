# client-go

Go gRPC client for Aurae

## Installation

```shell
go get github.com/aurae-runtime/client-go
```

## Usage

### Connecting with TLS

```go
tlsConfig := &tls.Config{} // Provide your TLS configuration

client, err := aurae.NewClient("localhost:8080", credentials.NewTLS(tlsConfig))
if err != nil {
panic(err)
}
```

### Interacting with auraed

```go
exec := &runtime.Executable{
    Command: "tail -f /dev/null",
    Comment: "my-executable",
}
status, err := client.Runtime().Exec(context.TODO(), exec)
if err != nil {
    panic(err)
}

fmt.Printf("Executable running with pid %d\n", status.Proc.Pid)
```