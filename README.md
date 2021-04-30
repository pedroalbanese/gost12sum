# GOST12SUM(2)

<pre>
GOST R 34.11-2012 Streebog256/512 Hashsum Tool - ALBANESE Lab (c) 2020-2021

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
<pre>
$ ./gost12sum [-r] -t "*.*" > hash.txt
</pre>
##### Always works in binary mode. 

#### Check hashsum file:
<pre>
$ ./gost12sum [-v] -c hash.txt
</pre>
##### Exit code is always 0 in vebose mode. 
