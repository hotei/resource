
##  resource

(c) 2015 David Rook - all rights reserved. Released under BSD 2-Clause License,
the details of which are at the end of this document.

### What Is It?

Yet another resource generating program.  There are a BUNCH of others but I
wanted something very simple.  This fits my requirements.

## Installation

If you have a working go installation:

> ```go get github.com/hotei/resource```

If you don't yet have go installed you might find it easiest to download 
the zip version and unpack locally.

## Features

* simple
* can specify package name or default to main
* can specify line length (defaults to 80)
* produces readable go code
* resources are stored in base64, not compressed
	* the resource file will be about 50% larger than the original source.
	* if compression is desirable the change is easy (it's on my TODO list)
	* my current application is mostly for images so it would not benefit from compression

## Usage

A resource is normally created as part of a makefile or "go generate" process something like :

```go	
	resource -source="loki.jpg" -rc="loki.go" -var="lokiJpgBites"
	resource -source="thor.jpg" -rc="thor.go" -var="thorJpgBites"
	go build
```
The end result is to make two global variables (lokiJpgBites,thorJpgBites []byte)
available to the program being built. 

Since the decode process happens at initialzation the slices are available without
the programmer doing anything extra.  Typical use after that is to attach a 
reader to the byte slice and pass the reader on to the program as if it had
come from opening a file.

### Limitations

* >>> Warning <<< The file designated as output by -rc="file" will be overwritten without warning.
* More of a thing to note than a limitation but each resource comes contained
in its own file.go If you have a large number of resources that could be a pain to manage.
While I don't need it (yet), the code could easily be refactored to take multiple
-source arguments and put the results into a single file.  It's a TODO.
* Since it's done at compile time there is a cost to compile the resources.  If
the files are many megabytes it may be noticable.  One way to mitigate that is
to make the resource file a make target that's only created when the source
changes. This makes the single file per resource more attractive. For instance:

```
loki.go:	loki.jpg
	resource -source="loki.jpg" -rc="loki.go" -var="lokiJpgBites"
	
all: loki.go
	go build
	
```

In this case "make all" will only create loki.go if loki.jpg changes.

### TODO

NOTE:  "higher" relative priority is at top of list

1. Allow for multiple source files.
1. Compression option
1. Get timing data for compilation of large resources

### BUGS

1.  as needed (none active at the moment)

### References

* [go language reference] [1] 
* [go standard library package docs] [2]
* [Source for resource program on github] [3]

#### Some similar Packages/Programs in alpha order (as of 2015-08-15)

* [akavel/embd-go] [6]
* [jteeuwen/go-bindata] [5]
* [tv42/becky] [4]

[1]: http://golang.org/ref/spec/ "go reference spec"
[2]: http://golang.org/pkg/ "go package docs"
[3]: http://github.com/hotei/resource "github.com/hotei/resource"
[4]: http://github.com/tv42/becky "github.com/tv42/becky"
[5]: http://github.com/jteeuwen/go-bindata "github.com/jteeuwen/go-bindata"
[6]: http://github.com/akavel/embd-go "github.com/akavel/embd-go"

Comments can be sent to David Rook  <hotei1352@gmail.com>  Issues can be filed
on github if needed.

License
-------
The 'resource' go program(s) are distributed under the Simplified BSD License:

> Copyright (c) 2015 David Rook. All rights reserved.
> 
> Redistribution and use in source and binary forms, with or without modification, are
> permitted provided that the following conditions are met:
> 
>    1. Redistributions of source code must retain the above copyright notice, this list of
>       conditions and the following disclaimer.
> 
>    2. Redistributions in binary form must reproduce the above copyright notice, this list
>       of conditions and the following disclaimer in the documentation and/or other materials
>       provided with the distribution.
> 
> THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDER ``AS IS'' AND ANY EXPRESS OR IMPLIED
> WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND
> FITNESS FOR A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL <COPYRIGHT HOLDER> OR
> CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR
> CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
> SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON
> ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING
> NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF
> ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

DO NOT EDIT BELOW - Generated by [godoc2md](http://godoc.org/github.com/davecheney/godoc2md)

# .

resource.go (c) 2015 David Rook - all rights reserved

tool to convert files to go code - see also go generate for other ways








- - -
DO NOT EDIT ABOVE - Generated by [godoc2md](http://godoc.org/github.com/davecheney/godoc2md)

github.com/hotei/resource imports 
```
	bytes
	encoding/base64
	errors
	flag
	fmt
	io
	io/ioutil
	log
	math
	os
	path/filepath
	reflect
	runtime
	sort
	strconv
	strings
	sync
	sync/atomic
	syscall
	time
	unicode
	unicode/utf8
	unsafe
```
```
deadcode results:

```
