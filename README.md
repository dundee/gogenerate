
# go:generate

Source code for the [talk "go:generate"](https://www.youtube.com/watch?v=ClW_g1iDGi4) given on Prague Golang Meetup on 28th June 2022.

The talk demonstrared how to use protobuf's custom options for declaring how the message should be marshalled for logging.
We looked on possibilities how to process the proto message metadata and how to generate the code responsible for doing the marshalling.

## Example

Having a protobuf message:

```go
message User {
    uint64 id = 1 [(log.key) = "user.id"];
    string name = 2 [(log.key) = "user.name"];
    string email = 3;
    string phone = 4 [(log.key) = "user.phone"];
    string password = 5;
}
```

will generate method:

```go
func (u *User) MarshalLog() log.Fields {
	return log.Fields{
		"user.id":   u.Id,
		"user.name":   u.Name,
		"user.phone":   u.Phone,
	}
}
```

which will in the end create JSON log entry:

```json
{"level":"info","msg":"Received user","time":"2022-06-29T13:53:04+02:00","user.id":1,"user.name":"John Doe","user.phone":"555"}
```
