# simple-store

Simply Store your file
Made to make your work easier as for saving files.

## Features

- [x] Chunking file (in one server)
- [ ] Replication (between servers)
- [ ] Centeral database for saving server stats for replication
- [ ] Dockerizing and adding installation methods

## Made out of

- [github.com/labstack/echo](https://github.com/labstack/echo)
- [github.com/sirupsen/logrus](https://github.com/sirupsen/logrus)
- [github.com/spf13/cobra](https://github.com/spf13/cobra)
- [github.com/spf13/pflag](https://github.com/spf13/cobra)
- [github.com/swaggo/echo-swagger](https://github.com/swaggo/echo-swagger)
- [github.com/etcd-io/bbolt](https://github.com/etcd-io/bbolt)

## Usage

run project with `--help` option

All the flags can be set using Environment Variables (if you prefer) Here is the naming convention for them

All of them are starting with `SS_` and the Option (replace `-` with `_`) and make it Uppercase

Also, prometheus exporter will start at the same address in `/metrics` address

and swagger will start at `/swagger` address

## Development

### Swagger

After making changes in api run the command below to generate openapi spec

> make sure swaggo is installed `go install github.com/swaggo/swag/cmd/swag@latest`

```shell
cd api
swag init -g api.go
```

### Package Structure

- As you can see utilities are in internal package with sub packages
- also api related things are inside api package (docs package is auto generated by swaggo)
- put all of your types inside types package
- cmd package is just for initializing project using command line
