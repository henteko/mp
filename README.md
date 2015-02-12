mp [![Build Status](https://drone.io/github.com/henteko/mp/status.png)](https://drone.io/github.com/henteko/mp/latest)
====

## Description
mp is simple analyzer for .mobileprovision file

## Usage

### dump

```bash
$ mp /path/to/example.mobileprovision
<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
<dict>
    <key>AppIDName</key>
    ....
</dict>
</plist>
```

### read

```bash
$ mp --key UUID /path/to/example.mobileprovision
$ mp -k UUID /path/to/example.mobileprovision //short ver
XXXXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXXXX
```

### install

```bash
$ mp --install /path/to/example.mobileprovision
$ mp -i /path/to/example.mobileprovision //short ver
Installed to /Users/example/Library/MobileDevice/Provisioning Profiles/XXXXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXXXX.mobileprovision
```

## Install

To install, use `go get`:

```bash
$ go get -d github.com/henteko/mp
$ cd $GOPATH/src/github.com/henteko/mp
$ make install
```

## Contribution

1. Fork ([https://github.com/henteko/mp/fork](https://github.com/henteko/mp/fork))
1. Create a feature branch
1. Commit your changes
1. Rebase your local changes against the master branch
1. Run test suite with the `go test ./...` command and confirm that it passes
1. Run `gofmt -s`
1. Create new Pull Request

## Author

[henteko](https://github.com/henteko)
