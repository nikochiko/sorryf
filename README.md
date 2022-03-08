# sorryf
A Go pkg born out of a typo, which a Go dev may interpret as a function that takes format string and arguments for substitution.

The inspiration for this package are `fmt`, `errors`, and `log` packages that have the commonly used `Printf`, `Errorf`, `Fatalf`
methods. After working very long with this convention, a Go developer, after seeing some noun followed by an `f` would remember
this convention and think that this is a method that accepts a format string and substitution arguments.

One day, this happened to me and I decided to take it seriously. https://twitter.com/n1kochiko/status/1501171417882054657

## Usage

Install the pkg with `go get github.com/nikochiko/sorryf`.

It has exactly one method `sorryf.Sorryf`, with the same argument and return types as `fmt.Sprintf` (it returns a string), only
that a frowning face (☹️) will be prepended to the string. `sorryf.Sorryf("hello %s", world)` would return `"☹️ hello, world"`.

Use it in code like:
```go
package main

import (
        "fmt"

        "github.com/nikochiko/sorryf"
)

func main() {
        fmt.Println(sorryf.Sorryf("Sorry brother"))
}
```

This will output:
```
☹️ Sorry brother
```
