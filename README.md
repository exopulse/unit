# exopulse unit package
Golang package unit contains type wrappers for common units like size. Conversions from and to human readable formats are supported.

[![CircleCI](https://circleci.com/gh/exopulse/unit.svg?style=svg)](https://circleci.com/gh/exopulse/unit)
[![GitHub license](https://img.shields.io/github/license/exopulse/unit.svg)](https://github.com/exopulse/unit/blob/master/LICENSE)

# Overview

This package contains type wrappers for common units like size. Conversions from and to human readable formats are supported.

## Features

### Size 

Size is simple type wrapper around uint64 providing parsing/formatting methods for size expressions.

Supported formats:
 - 15
 - 20 KB
 - 30 M
 - 40M
 - 1.5 GB


# Using unit package

## Installing package

Use go get to install the latest version of the library.

    $ go get github.com/exopulse/unit
 
Include unit in your application.
```go
import "github.com/exopulse/unit"
```

## Use Size to parse size expression from string
```go
s, err := ParseSize("120 MB")
```

## Use type cast to create a Size object from integer
```go
s := Size(1024)
```

## Use String() method to render size in "unit" format (no space between value and unit)
```go
fmt.Println(s.String())
```
 1024MB

## Use Humanize() method to render size in "human" format (space inserted between value and unit)
```go
fmt.Println(s.Humanize())
```
 1024 MB

# About the project

## Contributors

* [exopulse](https://github.com/exopulse)

## License

Unit package is released under the MIT license. See
[LICENSE](https://github.com/exopulse/unit/blob/master/LICENSE)
