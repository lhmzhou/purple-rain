# purple-rain

`purple-rain` is a simple TCP listener that returns what it receives on the TCP session. `purple-rain` can be used as a dummy backend for any systems which require testing TCP echo requests and response functionality.

## Build from CLI

To start listening simply execute `purple-rain` with the following parameters.

```sh
$ purple-rain -p 0042
```

Verbosity can be added with the `-v` flag.

```sh
$ purple-rain -v -p 0042
```

To listen on a specific IP interface simply use the `-i` flag.

```sh
$ purple-rain -v -i 0.0.0.0 -p 0042
```  

## Build with Docker

This repository includes a simple `Dockerfile`. To build and run this container simply execute the following.

```sh
$ sudo docker build -t purple-rain .
$ sudo docker run -p 0042:0042 --name purple-rain purple-rain
```
