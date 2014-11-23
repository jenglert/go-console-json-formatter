go-console-json-formatter
=========================

A silly little console json formatter, written in go.

Installation
============

```Shell
go get github.com/jenglert/go-console-json-formatter
```

You may want to symlink the executable to somewhere on your path:

```Shell
ln -s $GOPATH/bin/go-console-json-formatter /usr/local/bin/jf
```

Formatting Json
===============
```Shell
cat some_json_file.json | jf
```
