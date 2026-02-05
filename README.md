# passgen

Simple password generator written in Go. (crypto-rand)

## Install

```bash
go install github.com/kyouko17/cmd/passgen/passgen@latest
```
## Usage
The command generates a password of 4-32 characters, which contains at least 1 special character, 1 uppercase letter, 1 lowercase letter, 1 number

```bash
passgen
```

-length
```bash
passgen -length 16
