# goutil

A minimal, reusable Go starter boilerplate for utility functions and structs.
Perfect for jumpstarting your own shared library.

## Getting Started

To use this as a boilerplate for your own utilities, you can clone or fork this repository.

**Important:** Before publishing your own package, replace `github.com/yourusername/goutil` with your actual module path in `go.mod` and all import statements:

```bash
# Example using `find` and `sed` (macOS/Linux)
find . -type f -name '*.go' -exec sed -i '' 's|github.com/yourusername/goutil|github.com/myname/myutils|g' {} +
```

## Installation (For users of your package)

Once published to GitHub, other projects can install your package using:

```bash
go get github.com/yourusername/goutil
```

## Usage Example

```go
package main

import (
	"fmt"
	"github.com/yourusername/goutil/strutil"
	"github.com/yourusername/goutil/sliceutil"
)

func main() {
	// String utility
	reversed := strutil.Reverse("hello world")
	fmt.Println(reversed) // dlrow olleh

	// Slice utility (Generics)
	numbers := []int{1, 2, 3, 4, 5}
	hasThree := sliceutil.Contains(numbers, 3)
	fmt.Println("Contains 3:", hasThree) // Contains 3: true
}
```

## Structure

* `strutil/`: String manipulation utilities.
* `sliceutil/`: Generic slice utilities.
* `internal/`: Private helpers. Other modules cannot import these packages.
* `.github/workflows/test.yml`: GitHub Action for automated testing on PRs and commits to main.

## Internal packages

Go treats any directory named `internal` as private to the module (and to packages under the same parent). Create one by adding a folder and a package inside it:

```
goutil/
  internal/
    validate/
      validate.go
  strutil/
  sliceutil/
```

Import it from code in this repo only:

```go
import "github.com/rijalghodi/goutil/internal/validate"
```

Rules:

* The last `internal` path segment is the boundary. `github.com/rijalghodi/goutil/sliceutil` can import `github.com/rijalghodi/goutil/internal/validate`.
* An external module cannot import `github.com/rijalghodi/goutil/internal/...`. The compiler rejects it.
* Nested internals work the same way: `internal/foo/internal/bar` is only visible under `internal/foo`.

Use `internal` for shared helpers you do not want as public API. Public functions stay in `strutil` and `sliceutil`.

## Running Tests

To run the unit tests for all packages:

```bash
go test -v ./...
```
