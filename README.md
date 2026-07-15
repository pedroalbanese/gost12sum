# GOST12SUM(2)
[![ISC License](http://img.shields.io/badge/license-ISC-blue.svg)](https://github.com/pedroalbanese/gost12sum/blob/master/LICENSE.md) 
[![GitHub downloads](https://img.shields.io/github/downloads/pedroalbanese/gost12sum/total.svg?logo=github&logoColor=white)](https://github.com/pedroalbanese/gost12sum/releases)
[![GoDoc](https://godoc.org/github.com/pedroalbanese/gost12sum?status.png)](http://godoc.org/github.com/pedroalbanese/gost12sum)
[![Go Report Card](https://goreportcard.com/badge/github.com/pedroalbanese/gost12sum)](https://goreportcard.com/report/github.com/pedroalbanese/gost12sum)
[![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/pedroalbanese/gost12sum)](https://golang.org)
[![GitHub release (latest by date)](https://img.shields.io/github/v/release/pedroalbanese/gost12sum)](https://github.com/pedroalbanese/gost12sum/releases)
### GOST R 34.11-2012 Streebog256/512 Hashsum Tool
```
Usage of gost12sum:
./gost12sum [-c <hash.g12>] [-r] [-l] <file.ext>
  -c string
        Check hashsum file.
  -l    Use 512 bit hash (default 256-bit)
  -r    Process directories recursively.
```

### Examples:

#### Generate hashsum list:
```sh
$ ./gost12sum [-r] [-l] "*.*" > hash.g12
```
##### Always works in binary mode. 

#### Check hashsum file:
```sh
$ ./gost12sum -c hash.g12
$ echo $?
```

## License

This project is licensed under the ISC License.
##### Copyright (c) 2020-2026 Pedro F. Albanese - ALBANESE Research Lab.
