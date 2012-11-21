# utfconv

utfconv is the fast way to convert []byte to string by chooing charset

## Installation

Make sure you have the a working Go environment. See the [install instructions](http://golang.org/doc/install.html).

To install utfcon.go, simply run:

    go get github.com/xsoameix/utfconv

To compile it from source:

    git clone git@github.com:xsoameix/utfconv.git
    cd utfconv && go build

## Example

    package main

    import (
        "github.com/xsoameix/utfconv"
    )

    func main() {
        utf8 := []byte("Hello, world")
        utf16be := []byte("\x00\x48\x00\x65\x00\x6C" +
            "\x00\x6C\x00\x6F\x00\x2C\x00\x20"+
            "\x00\x77\x00\x6F\x00\x72\x00\x6C"+
            "\x00\x64")
        fmt.Printf("%s\n%s\n",
            utfconv.Read("utf8", utf8),
            utfconv.Read("utf16be", utf16be))
    }

To run the application, put the code in a file called hello.go and run:

    go build hello.go
