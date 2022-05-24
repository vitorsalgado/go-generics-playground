# Go Arrays · [Generics Playground] ·  [![ci](https://github.com/vitorsalgado/goarrays/actions/workflows/ci.yml/badge.svg)](https://github.com/vitorsalgado/goarrays/actions/workflows/ci.yml)

Simple utility to handle arrays in similar way to JavaScript arrays API.  
Developed to learn Go Generics.

## Usage
```
s := Of([]string{"1", "2", "3"}).
    Map(func(i string) string { return fmt.Sprintf("%s-test", i) }).
    Filter(func(i string) bool { return strings.Contains(i, "1") }).
    Collect()
```
