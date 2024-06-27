# GO111MODULE: Controls module support. auto enables it when a go.mod file is present.

GO111MODULE is an environment variable that controls the behavior of the Go module system. The Go module system is a feature introduced in Go version 1.11 to manage dependencies and versioning of packages in a Go project.

Before the introduction of Go modules, Go relied on the GOPATH environment variable to manage dependencies. In this model, all Go packages and their dependencies were stored in a single directory called GOPATH. This approach was simple, but it led to several problems, including version conflicts, the need to manually manage dependencies, and difficulties in sharing code.

The GO111MODULE environment variable has three possible values:

1. GO111MODULE=off: This value disables the Go module system and uses the legacy GOPATH mode instead. In this mode, dependencies are downloaded and stored in the GOPATH directory, and the go command looks for packages in the directories specified by the GOPATH environment variable.

2. GO111MODULE=on: This value enables the Go module system and uses modules to manage dependencies. In this mode, the go command looks for the go.mod file in the project directory to determine the required dependencies and their versions. If the file exists, the command downloads the required dependencies and stores them in a local cache, which can be shared between projects.

3. GO111MODULE=auto: This value enables the Go module system if a go.mod file is present in the project directory. If no go.mod file is found, the legacy GOPATH mode is used.


GO111MODULE is usually set as an environment variable before running the Go command. For example, you can set it as follows:

export GO111MODULE=on
Or, you can set it only for a specific command:

GO111MODULE=on go build

# GOARCH: Architecture target for the Go compiler, amd64 specifies a 64-bit architecture.

# GOBIN: Directory where go install installs binaries. Empty means it defaults to GOPATH/bin. some packages run .exe binaries


GOBIN: This environment variable defines the directory where Go binaries (executables) are installed when you use the go install command. It should be added to your system's PATH to execute these binaries conveniently.


# GOCACHE: Directory for cached build results. 

The GOCACHE environment variable in Go specifies the directory where the Go toolchain stores cached build results. This cache significantly speeds up subsequent builds by reusing results from previous builds when possible.

Default Location: By default, GOCACHE is set to a directory under the system's cache directory, such as $HOME/.cache/go-build on Unix systems or %LOCALAPPDATA%\go-build on Windows.

Setting GOCACHE:

```
export GOCACHE=/path/to/custom/cache

```

# GOENV: Path to Go environment configuration file. 

The GOENV environment variable specifies the path to the Go environment configuration file. This file can store environment variable settings that override the default Go environment settings.

Default Location: The default is $HOME/.config/go/env on Unix systems and %USERPROFILE%\AppData\Roaming\go\env on Windows.
Setting GOENV:
```
export GOENV=/path/to/custom/env

```
# GOEXE: File extension for executables, .exe for Windows.

GOEXE: File Extension for Executables

The GOEXE environment variable specifies the file extension for executable files. This is particularly useful for cross-compiling Go programs for different operating systems.

Typical Values:
For Windows: .exe

For other systems: an empty string
Setting GOEXE:

```
export GOEXE=.exe

```

#  GOEXPERIMENT: Space-separated list of enabled/disabled experimental features.

GOEXPERIMENT: Space-Separated List of Enabled/Disabled Experimental Features
The GOEXPERIMENT environment variable controls the enabling or disabling of experimental features in the Go toolchain. This variable is used by developers to test and provide feedback on new features before they are officially released.

Example:
```
export GOEXPERIMENT=unified

```
#  GOFLAGS: Default flags for the go command.

Purpose: This variable allows you to specify additional flags for the go command.

Usage: Set GOFLAGS to include additional command-line flags when running go commands.




#  GOHOSTARCH: Host architecture, amd64 for 64-bit.

The GOHOSTARCH environment variable specifies the architecture of the host machine where the Go toolchain is running. This is used by the Go build system to generate binaries compatible with the host architecture.

Common Values:
amd64: 64-bit x86 architecture
386: 32-bit x86 architecture
arm: 32-bit ARM architecture
arm64: 64-bit ARM architecture
Example:

```
export GOHOSTARCH=amd6

```
#  GOHOSTOS: Host operating system, windows

The GOHOSTOS environment variable specifies the operating system of the host machine where the Go toolchain is running. This helps the Go build system to generate binaries compatible with the host operating system.

Common Values:
linux
windows
darwin (for macOS)
freebsd
Example:
```
export GOHOSTOS=windows

```
#  GOINSECURE: Comma-separated list of modules or hosts the go command should treat as insecure.

The GOINSECURE environment variable allows you to specify a comma-separated list of modules or hosts that the go command should treat as insecure. This can be useful when working with private modules that are not available over HTTPS or have self-signed certificates.

Example:

```
export GOINSECURE=example.com,anotherexample.com/private

```
#  GOMODCACHE: Directory for storing downloaded modules.

The GOMODCACHE environment variable specifies the directory where the Go toolchain stores downloaded modules. This directory acts as a cache for module downloads, improving build performance by reusing previously downloaded modules.

Default Location: The default location is $GOPATH/pkg/mod.

Setting GOMODCACHE:

```
export GOMODCACHE=/path/to/custom/modcache

```

#  GONOPROXY: Comma-separated list of module patterns to exclude from the proxy.

This variable specifies the Go module proxy server used for downloading module dependencies. It can be set to a custom proxy or one of the publicly available ones like "https://proxy.golang.org.

Proxy Control: GOPROXY helps manage the source of Go module downloads, enhancing security and reliability.


#  GONOSUMDB: Comma-separated list of module patterns to exclude from sumdb.

The GONOSUMDB environment variable specifies a comma-separated list of module patterns that should be excluded from the Go checksum database (sumdb). The Go checksum database is used to ensure the integrity of downloaded modules by verifying their checksums against a trusted source. By excluding certain modules from sumdb, you can bypass this verification, which may be necessary for private or internal modules.

Example:

```
export GONOSUMDB=example.com,anotherexample.com/private

```

#  GOOS: Target operating system, windows.

The GOOS environment variable specifies the target operating system for which the Go toolchain should build the binary. This is particularly useful for cross-compiling programs to run on different operating systems.

Common Values:
linux
windows
darwin (for macOS)
freebsd
netbsd
openbsd
plan9
solaris

Example:

```
export GOOS=windows

```

#  GOPATH: Workspace directory for Go projects and dependencies.

```
The GOPATH environment variable specifies the location of your workspace.
If no GOPATH is set, it is assumed to be $HOME/go on Unix systems and %USERPROFILE%\go on Windows. If you want to use a custom location as your workspace, you can set the GOPATH environment variable. 

```

# GOPRIVATE: Comma-separated list of module patterns that are private.

The GOPRIVATE environment variable controls which modules the go command considers to be private (not available publicly) and should therefore not use the proxy or checksum database. The variable is a comma-separated list of glob patterns (in the syntax of Go's path.Match) of module path prefixes. For example,

```
GOPRIVATE=*.corp.example.com,rsc.io/private

```

causes the go command to treat as private any module with a path prefix matching either pattern, including git.corp.example.com/xyzzy, rsc.io/private, and rsc.io/private/quux.

if your private repos belonged to a organization, such as a github organization org_name, just type the organization name:

```
GOPRIVATE=github.com/org_name

```
# GOPROXY: URL of the Go module proxy.

The GOPROXY environment variable specifies the URL of the Go module proxy, which the go command uses to download modules. Using a proxy can improve download speeds and reliability by caching and distributing modules.

Example:

```
export GOPROXY=https://proxy.golang.org

```

# GOROOT: Root directory where Go is installed.

The GOROOT environment variable specifies the root directory where the Go toolchain is installed. This directory contains the Go binary, standard library, and other tools.

Example:

```
export GOROOT=/usr/local/go

```
# GOSUMDB: URL of the Go checksum database.

The GOSUMDB environment variable specifies the URL of the Go checksum database, which is used to verify the integrity of modules.

Example:

```
export GOSUMDB=sum.golang.org

```
# GOTMPDIR: Directory for temporary files. Empty uses the system default.

The GOTMPDIR environment variable specifies the directory for temporary files created by the Go toolchain. If empty, the system's default temporary directory is used.

Example:

```
export GOTMPDIR=/path/to/tempdir

```

# GOTOOLCHAIN: Go toolchain to use.

The GOTOOLCHAIN environment variable specifies which Go toolchain to use. This can be useful for testing different versions of the Go toolchain.

Example:

```
export GOTOOLCHAIN=go1.18

```

# GOTOOLDIR: Directory containing Go tools.

The GOTOOLDIR environment variable specifies the directory containing Go tools, such as the compiler and linker.

Example:

```
export GOTOOLDIR=/usr/local/go/pkg/tool

```
# GOVCS: Comma-separated list of version control commands.

The GOVCS environment variable specifies a comma-separated list of version control commands that the go command should use when interacting with version control systems.

Example:

```
export GOVCS=git:https,svn:https

```

# GOVERSION: Installed Go version.

The GOVERSION environment variable specifies the version of the installed Go toolchain.

Example:

```
export GOVERSION=go1.18

```
# GCCGO: Command for the gccgo compiler.

The GCCGO environment variable specifies the command for the gccgo compiler, an alternative Go compiler that uses the GCC backend.

Example:

```
export GCCGO=gccgo

```
# GOAMD64: Optimizations for the AMD64 architecture.

The GOAMD64 environment variable specifies optimization options for the AMD64 architecture.

Example:

```
export GOAMD64=v3

```

# AR: Command for the archiver.

The AR environment variable specifies the command for the archiver used by the Go toolchain.

Example:

```
export AR=ar

```

# CC: Command for the C compiler.

The CC environment variable specifies the command for the C compiler used by the Go toolchain.

Example:

```
export CC=gcc

```

# CXX: Command for the C++ compiler.

The CXX environment variable specifies the command for the C++ compiler used by the Go toolchain.

Example:

```
export CXX=g++

```
# CGO_ENABLED: Whether to Enable CGO (C Extensions) Support

The CGO_ENABLED environment variable specifies whether to enable CGO, which allows Go programs to call C code.

# CGO_ENABLED: Whether to enable CGO (C extensions) support.

Example:

```
export CGO_ENABLED=1

```

# GOMOD: Path to the go.mod file in use.

The GOMOD environment variable specifies the path to the go.mod file currently in use. This file defines the module's dependencies.

Example:

```
export GOMOD=/path/to/go.mod

```
# GOWORK: Path to the workspace file.

The GOWORK environment variable specifies the path to the workspace file, which can be used to manage multiple modules in a single workspace.

Example:

export GOWORK=/path/to/go.work

# CGO_CFLAGS: Flags for the C compiler.

The CGO_CFLAGS environment variable specifies flags to pass to the C compiler when building CGO code.

Example:
```
export CGO_CFLAGS="-I/usr/local/include"

```

# CGO_CPPFLAGS: Flags for the C preprocessor.

CGO_CPPFLAGS: Flags for the C Preprocessor
The CGO_CPPFLAGS environment variable specifies flags to pass to the C preprocessor when building CGO code.

Example:

```
export CGO_CPPFLAGS="-DDEBUG"

```

# CGO_CXXFLAGS: Flags for the C++ Compiler
The CGO_CXXFLAGS environment variable specifies flags to pass to the C++ compiler when building CGO code.

Example:

```
export CGO_CXXFLAGS="-I/usr/local/include/c++/v1"

```
# CGO_FFLAGS: Flags for the Fortran Compiler
The CGO_FFLAGS environment variable specifies flags to pass to the Fortran compiler when building CGO code.

Example:

```
export CGO_FFLAGS="-I/usr/local/include"

```
# CGO_LDFLAGS: Flags for the Linker
The CGO_LDFLAGS environment variable specifies flags to pass to the linker when building CGO code.

Example:

```

export CGO_LDFLAGS="-L/usr/local/lib"

```

# PKG_CONFIG: Path to pkg-config for Finding C Libraries

The PKG_CONFIG environment variable specifies the path to the pkg-config tool, which is used to find C libraries and their metadata.

Example:

```

export PKG_CONFIG=/usr/local/bin/pkg-config


```

# GOGCCFLAGS: Flags for the gccgo Compiler

The GOGCCFLAGS environment variable specifies flags to pass to the gccgo compiler.

Example:

```
export GOGCCFLAGS="-O2 -g"

```

These environment variables configure various aspects of Go's behavior, such as compiler settings, paths for files, module management, and cross-compilation targets.