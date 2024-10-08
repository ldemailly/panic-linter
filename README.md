# panic-linter
Go linter that flags panic() call that don't have a comment explaining why panic


## Why?

panic should only be used very sparingly, for catching bugs basically, and thus deserve a comment to confirm that that's indeed the case

bad:
```go
  panic("catch this")
```

good:
```go
  panic(fmt.Sprintf("bug: unexpected input=%v because...", input)) // Shouldn't happen unless we have bug
```
