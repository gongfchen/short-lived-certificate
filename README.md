# An applicable Short-Lived Certificates Framework

This framework aims to provide an effiecient method for revoking compromised or misissued Public Key Certificates. It avoids online revocation checking and contructs the certificate as a caching mechanism.

To learn more about the design, check out our [blog post](https://gongfchen.github.io/files/blog.pdf).

## Setup

```
$ pip3 install merklelib
$ cd CA/script
$ go mod download
```

## Building

```
$ cd <CA or middle-daemon or website-daemon>/script
$ go build
```

## Usage

### CA

To start running CA:

```
$ ./CA
```

### Middle daemon

To start running Midle Daemon:

```
./Middle-daemon
```

To request short-lived certificates for a domain (default to be 365 certificates). Need domain public keys available:

```
request [domain_name]
```

To request daily decryption key for a domain. Need to provide the day number:

```
key [domain_name] [day_number]
```

### Website daemon: 

To verify a certificate, run the following:

```
./Website-daemon
...
Listening to request...
[domain_name] [num_of_day]
```
