# GopherLibs => Big League Stats [![Go Reference](https://pkg.go.dev/badge/github.com/gopherlibs/big-league-stats.svg)](https://pkg.go.dev/github.com/gopherlibs/big-league-stats) [![Software License](https://img.shields.io/badge/license-MIT-blue.svg)](https://raw.githubusercontent.com/gopherlibs/big-league/trunk/LICENSE)

*This project is really early and figuring out its identity. Don't expect a stable (Go) API at the moment.*

`Big League Stats` is a Go (Golang) wrapper and SDK for the official MLB Stats API.
A basic CLI is also included.

The CLI is and will remain fairly basic.
If you want more functionality and a full blown TUI, you can give [mlbt](https://github.com/mlb-rs/mlbt) a try.


## Table of Contents

- [Goals](#goals)
- [Requirements](#requirements)
- [Usage](#usage)
- [Development](#development)
- [Credits](#credits)


## Goals

- The first right now is to get this module usable to return the current MLB standings. This will then be used in the [wtfutil/wtf](https://github.com/wtfutil/wtf) project as a widget/module.
- Keep improving MLB data
- add the first NFL data, again, standings


## Requirements

- The minimum Go version supported is v1.25.x.
- Internet, access is needed to pull data via APIs.


## Usage

`Big League Stats` is a Go module so the best way to use it is to add an import in the file you want to use it in and then run `go mod tidy` to get it downloaded.

```go
import(
	"github.com/gopherlibs/big-league-stats/sdk"
)
```

Alternatively, you can run `go get github.com/gopherlibs/big-league-stats` in your project directory.


## Usage

```go
package main

import (
	"fmt"

	"github.com/gopherlibs/big-league-stats/sdk"
)

func main() {

	// choose a top league
	conference := sdk.MLB.NationalLeague

	// pull standings data
	c.Standings()

	// look through a divison's teams to see data
	for _, t := conference.East.Teams(){
		fmt.Println( t.Name )
	}
}
```

## Development

This library is written and tested with Go v1.25+ in mind.
`go fmt` is your friend.
Please feel free to open Issues and PRs are you see fit.
Any PR that requires a good amount of work or is a significant change, it would be best to open an Issue to discuss the change first.


## Credits

This module was written by Ricardo N Feliciano (FelicianoTech).
This repository is licensed under the MIT license.
This repo's license can be found [here](./LICENSE).
