To run this:
```bash
go run main.go
```

Change the `sub` value to make it fail:

```go
var data Matchable = Matchable{
	Jwt_claims: map[string]interface{}{
		"sub": "organization:zain-stacks-testing:project:zzz:stack:zst:deployment:production:operation:plan",
	},
}
```
