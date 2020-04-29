# PRIVY ID MINI E-WALLET

## Instalation
```bash
go get -v github.com/oktopriima/mini-e-wallet
```

## How to use

Make sure you has been installed `rubenv/sql-migrate`.
If not, you can following this documentation [link](https://github.com/rubenv/sql-migrate)

Move into your GOPATH directory
```bash
cd $GOPATH/src/github.com/oktopriima/mini-e-wallet
cp example-env.yaml env.yaml
cp dbconfig-example.yml dbconfig.yml
sql-migrate up --env=local
go build
./mini-e-wallet
```