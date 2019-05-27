# sliding-window

A grep-like utility for removing blocks of text. This is useful in
post processing files that have unexpected blocks of text.

## Usage

Don't include the trigger tags in the output

```
./sw data/test.txt start end
first
last
```

Include the trigger tags in the output

```
./sw -i data/test.txt start end
first
start
end
last
```

The test file `data/test.txt`. This file is also used for `go test`.

```
first
start
remove
end
last
```