
# An introduction to cross compilation with Go 1.1

This post is a compliment to one I wrote in August of last year, updating it for Go 1.1. Since last year tools such as goxc have appeared which go a beyond a simple shell wrapper to provide a complete build and distribution solution.

Introduction

Go provides excellent support for producing binaries for foreign platforms without having to install Go on the target. This is extremely handy for testing packages that use build tags or where the target platform is not suitable for development.

Support for building a version of Go suitable for cross compilation is built into the Go build scripts; just set the GOOS, GOARCH, and possibly GOARM correctly and invoke ./make.bash in $GOROOT/src. Therefore, what follows is provided simply for convenience.

Getting started

1. Install Go from source. The instructions are well documented on the Go website, golang.org/doc/install/source. A summary for those familiar with the process follows.

% hg clone https://code.google.com/p/go
% cd go/src
% ./all.bash
2. Checkout the support scripts from Github, github.com/davecheney/golang-crosscompile

% git clone git://github.com/davecheney/golang-crosscompile.git
% source golang-crosscompile/crosscompile.bash
3. Build Go for all supported platforms

% go-crosscompile-build-all
go-crosscompile-build darwin/386
go-crosscompile-build darwin/amd64
go-crosscompile-build freebsd/386
go-crosscompile-build freebsd/amd64
go-crosscompile-build linux/386
go-crosscompile-build linux/amd64
go-crosscompile-build linux/arm
go-crosscompile-build windows/386
go-crosscompile-build windows/amd64
This will compile the Go runtime and standard library for each platform. You can see these packages if you look in go/pkg.

% ls -1 go/pkg
darwin_386
darwin_amd64
freebsd_386
freebsd_amd64
linux_386
linux_amd64
linux_arm
obj
tool
windows_386
windows_amd64
Using your cross compilation environment

Sourcing crosscompile.bash provides a go-$GOOS-$GOARCH function for each platform, you can use these as you would the standard go tool. For example, to compile a program to run on linux/arm.

% cd $GOPATH/github.com/davecheney/gmx/gmxc
% go-linux-arm build
% file ./gmxc
./gmxc: ELF 32-bit LSB executable, ARM, version 1 (SYSV),
statically linked, not stripped
This file is not executable on the host system (darwin/amd64), but will work on linux/arm.

Some caveats

Cross compiled binaries, not a cross compiled Go installation

This post describes how to produce an environment that will build Go programs for your target environment, it will not however build a Go environment for your target. For that, you must build Go directly on the target platform. For most platforms this means installing from source, or using a version of Go provided by your operating systems packaging system.
If you are using

No cgo in cross platform builds

It is currently not possible to produce a cgo enabled binary when cross compiling from one operating system to another. This is because packages that use cgo invoke the C compiler directly as part of the build process to compile their C code and produce the C to Go trampoline functions. At the moment the name of the C compiler is hard coded to gcc, which assumes the system default gcc compiler even if a cross compiler is installed.

In Go 1.1 this restriction was reinforced further by making CGO_ENABLED default to 0 (off) when any cross compilation was attempted.

GOARM flag needed for cross compiling to linux/arm.

Because some arm platforms lack a hardware floating point unit the GOARM value is used to tell the linker to use hardware or software floating point code. Depending on the specifics of the target machine you are building for, you may need to supply this environment value when building.

% GOARM=5 go-linux-arm build
As of e4b20018f797 you will at least get a nice error telling you which GOARM value to use.

$ ./gmxc
runtime: this CPU has no floating point hardware, so it cannot
run this GOARM=7 binary. Recompile using GOARM=5.
By default, Go assumes a hardware floating point unit if no GOARM value is supplied. You can read more about Go on linux/arm on the Go Language Community Wiki.
