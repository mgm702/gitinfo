gitinfo
===
[![Build Status](https://travis-ci.org/mgm702/gitinfo.svg?branch=master)](https://travis-ci.org/mgm702/gitinfo)

gitinfo is a simple, and fast package for getting information on your favorite Github Repositories.
The goal was to create a simple CLI tool that any developer could use to retrieve various parts of info related to the Github Repository they searched for.

## Installation:

Make sure you have a working Go environment.  Go version 1.2+ is supported.  [See
the install instructions for Go](http://golang.org/doc/install.html).

To install gitinfo, simply run:

```
$ go get github.com/mgm702/gitinfo
```

Make sure your `PATH` includes the `$GOPATH/bin` directory so your commands can
be easily used:

```
export PATH=$PATH:$GOPATH/bin
```

## Commands:
Now you should be able to use the following commands in order to gather specific info you desire:
```
     url, u          Prints the Urls of a project
     issues, i       Issues of a project
     stars, s        Stars of a project
     forks, f        Forks for a project
     size, k         Size of a project
     description, d  Description of a project
     help, h         Shows a list of commands or help for one command

```

## Global Options:

```
   --help, -h     show help
   --version, -v  print the version
```
