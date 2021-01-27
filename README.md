# colorwriter: a tabwriter that is escape code tolerant

This is a fork of text/tabwriter that is escape code friendly, in that:

- Escape codes (particularly, color codes; more are planned) are not counted in the width calculation for each table
- Escape codes are still printed

When coupled with a library like [fatih/color](https://github.com/fatih/color) you can get some great results. Check out our [example](example). For a quick demo, run these commands (will pull the repository):

```bash
go get -d github.com/erikh/colorwriter
go run github.com/erikh/colorwriter/example
```

Which will print some diagnostics and show you how the two libraries differ in
effect.

## Author

- Almost all of this is taken from golang 1.15's frozen text/tabwriter package, and belongs to the Go authors and Google, Inc.
- Erik Hollensbe <github@hollensbe.org> is responsible for the colorwriter additions and this repository.

## License

- BSD 3-Clause (Go License)
