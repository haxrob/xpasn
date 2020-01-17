# xpasn
Expands an autonomous system (AS) number to it's prefixes or individual IP addresses.

The main intention of the tool is to use it to chain into other tools that accept hosts via standard input, such as [httprobe](https://github.com/tomnomnom/httprobe) or vulnerability scanners.

# Installation

If you have a Go environment ready to go:

```
go get github.com/x1sec/xpasn
```

Precompiled executables for Windows, Linux and MacOS can be downloaded [here](https://github.com/x1sec/xpasn/releases)

# Usage

Running with no parameters will return a list of prefixes/subnets that belong to the AS. For example [Tesla Motors](https://ipinfo.io/AS394161)
```
$ xpasn AS394161
8.45.124.0/24
192.95.64.0/24
199.66.9.0/24
...
```
The `-e` flag will output the host IP addresses for each prefix.
```
$ xpasn -e AS394161
8.45.124.1
8.45.124.2
8.45.124.3
...
```
Example using httprobe:
```
$ xpasn -e AS394161 | httprobe 
```
