# gossdeep

Go version of ssdeep [http://ssdeep.sourceforge.net/].

## Status

[![GoDoc](https://godoc.org/github.com/dutchcoders/gossdeep?status.svg)](https://godoc.org/github.com/dutchcoders/gossdeep)

[1]: https://travis-ci.org/dutchcoders/gossdeep.png
[2]: https://travis-ci.org/dutchcoders/gossdeep

## Installation
```
go get github.com/dutchcoders/gossdeep

// use in your .go code
import (
    "github.com/dutchcoders/gossdeep"
)
```

## Usage

```
hash, err := HashString("test")
if err != nil {
        t.Fatal(err)
}
```

## Contact me

If I can help you, you have an idea or you are using gossdeep in your projects, don't hesitate to drop me a line (or a pull request): [@remco_verhoef](https://twitter.com/remco_verhoef)

## About

Written by [remco_verhoef](http://dutchcoders.io).

## License

[BSD 3-Clause](http://opensource.org/licenses/BSD-3-Clause) license, as [Go language](http://golang.org/LICENSE).
