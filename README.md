# GOST12SUM(2)
[![ISC License](http://img.shields.io/badge/license-ISC-blue.svg)](https://github.com/pedroalbanese/gost12sum/blob/master/LICENSE.md) 
[![GoDoc](https://godoc.org/github.com/pedroalbanese/gost12sum?status.png)](http://godoc.org/github.com/pedroalbanese/gost12sum)
[![Go Report Card](https://goreportcard.com/badge/github.com/pedroalbanese/gost12sum)](https://goreportcard.com/report/github.com/pedroalbanese/gost12sum)
### GOST R 34.11-2012 Streebog256/512 Hashsum Tool
<pre>
Usage of gost12sum:
gost12sum [-v] [-c &lt;hash.g12&gt;] [-r] [-l] -t &lt;file.ext&gt;

  -c string
        Check hashsum file.
  -l    Use 512 bit hash (default 256-bit)
  -r    Process directories recursively.
  -t string
        Target file/wildcard to generate hashsum list.
  -v    Verbose mode. (for CHECK command)</pre>

### Examples:

#### Generate hashsum list:
```sh
$ ./gost12sum [-r] -t "*.*" > hash.g12
```
##### Always works in binary mode. 

#### Check hashsum file:
```sh
$ ./gost12sum [-v] -c hash.g12
```
##### Exit code is always 0 in vebose mode. 

## License

This project is licensed under the ISC License.
##### Copyright (c) 2020-2021 Pedro Albanese - ALBANESE Research Lab.
