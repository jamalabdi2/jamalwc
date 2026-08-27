# jamalwc

A small, experimental implementation of the Unix `wc` command, written in Go.

This project is primarily a hands-on exercise for learning Go and exploring Unix-style systems programming. Rather than using higher-level abstractions, the implementation is built while studying how Go interacts with files, standard input, byte streams, and the operating system.

## About

`wc` is a Unix utility that reports information about files or standard input, such as:

* Number of lines
* Number of words
* Number of bytes

`jamalwc` aims to recreate these core behaviors while exploring how they can be implemented in Go.

## Usage

Build the program:

```bash
go build -o jamalwc
```

Run it with a file:

```bash
./jamalwc data/hello.txt
```

Or pipe data through standard input:

```bash
echo "hello world" | ./jamalwc
```

You can also use shell pipelines:

```bash
cat data/jamal.txt | ./jamalwc
```

## Why I Built This
To learn Go by implementing a small real world project and improving my understanding of Go programming and Unix systems.

The goal is to use a small Unix utility as a way to learn by implementation:

* Go's `os` package
* `io.Reader` and byte streams
* Standard input and output
* File descriptors
* File metadata and `os.File`
* Unix file modes
* Pipes and standard input
* Command-line arguments
* Byte and character processing
* Error handling in Go

The project will evolve as I explore more of Go's standard library and Unix internals.

## Learning Approach

Instead of starting with a framework or following a large tutorial, I'm building small pieces of Unix functionality and reading the relevant Go standard library source code along the way.

`jamalwc` is one of those experiments.

## Status

🚧 Experimental / Work in Progress

The implementation is intentionally small and may not yet match all of the behavior or edge cases of the system `wc` implementation.

## License

This project is for educational and experimental purposes.
