# go-unidecode

Port of https://metacpan.org/pod/Text::Unidecode in go

```go
unidecode := `úñícøðé`
ascii := unidecode.Decode(unicode) // => "unicode"
```

# internals

It uses `git@github.com:xuender/unidecode.git`, parses the resources an build the `Xxxx.go` files from there using `cmd/main.go`
