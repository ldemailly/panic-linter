[![Go Report Card](https://goreportcard.com/badge/github.com/ldemailly/panic-linter)](https://goreportcard.com/report/github.com/ldemailly/panic-linter)
[![GoDoc](https://godoc.org/github.com/ldemailly/panic-linter?status.svg)](https://pkg.go.dev/github.com/ldemailly/panic-linter)
[![codecov](https://codecov.io/gh/ldemailly/panic-linter/branch/main/graph/badge.svg)](https://codecov.io/gh/ldemailly/panic-linter)
[![Maintainability](https://api.codeclimate.com/v1/badges/bf83c496d49b169cd744/maintainability)](https://codeclimate.com/github/ldemailly/panic-linter/maintainability)
[![CI Checks](https://github.com/ldemailly/panic-linter/actions/workflows/include.yml/badge.svg)](https://github.com/ldemailly/panic-linter/actions/workflows/include.yml)

# panic-linter
`paniccheck` is a golang linter that flags panic() call that don't have a comment explaining why panic.


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
