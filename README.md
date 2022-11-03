# tinysoundfont.go

Go bindings for [tinysoundfont].

## Building

This package requires CGo to be enabled, which is the default for native builds.
However, you have to configure the specific compiler for your target if you are cross-compiling:

```
$ export CGO_ENABLED=1
$ export CC=(cross compiler)
$ GOOS=(target OS) GOARCH=(target arch) go build ...
```

## License

tinysoundfont.go is licensed under the MIT License.

[tinysoundfont]: https://github.com/schellingb/TinySoundFont
