# Routes

![Go](https://github.com/WheresAlice/routes/workflows/Go/badge.svg?branch=trunk)

Inspired by the many http routers for Go, this router operates on strings.  It is perhaps best explained with the [example/example.go](example).

A router takes multiple routes.  Each route comprises of a regexp pattern and a function to be used if the input string matches.

The router has an `Exec` function, which takes a string and returns a matching function or nil.

If a function was returned, then it can be called with the original input string.
