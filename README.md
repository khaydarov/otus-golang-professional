Course Plan
---

1. Introduction [Homework #1](./hw01_hello_otus)

    * GOPATH, GOROOT
    * Go modules
    * Compilation, formatting and linting

2. Syntax

    * Basic types
    * Arrays, Slices, Strings
    * Structs, Interfaces

3. Features and common mistakes [Homework #2](./hw02_unpack_string)

    * Scope, Closures
    * Defers
    * Slices, Maps, Runes

4. Best practices working with errors [Homework #3](./hw03_frequency_analysis)

    * Errors
    * Panic, Defer, Recover

5. Testing in Go

    * Testing apps in Go
    * Table testing
    * "testing" pkg and testify library

6. Advanced testing in Go

    * Mocking, Faker
    * Mutation test
    * Golden files

7. Interfaces internals [Homework #4](./hw04_lru_cache)

    * Empty interface (interface{})
    * Type Cast
    * Switch with interfaces

8. Goroutines and channels

    * Channel internals
    * Buffered ana unbuffered channels
    * Select operator, Timer

9. Synchronization primitives in details [Homework #5](./hw05_parallel_execution)

    * Working with sync package: WaitGroup, Once, Mutex
    * Race detector
    * Atomics

10. Additional synchronization primitives

    * Sync.Pool, read-write Mutex
    * Concurrent-safe maps
    * Memory model in GO

11. Concurrency patterns [Homework #6](./hw06_pipeline_execution)

    * Data synchronization patterns
    * Generators and pipeline
    * Working with multiple channels: or, fanin, fanout

12. Go internals. Scheduler

    * Main schedule structures: P, M, G
    * Goroutine switching mechanism
    * Processing syscals and netcals

13. Working with I/O [Homework #7](./hw07_file_copying)

    * Standard interfaces: Reader, Scanner, Writer, Closer
    * Working with data intensive apps
    * Regexp

14. Go internals. Memory and Garbage collection

    * Memory model in Linux
    * Memory features of Go
    * Garbage collection

15. Configuring and logging

    * Working with config files
    * Working with environment variables
    * Logging - log/slog

16. CLI [Homework #8](./hw08_envdir_tool)

    * Working with flags: flags, pflag, cobra
    * Working with arguments: os.Args, os.Args[1:], os.Args[1]
    * Working with subcommands

17. Reflection

    * Reflection internals
    * Working with reflect pkg: reflect.Type and reflect.Value
    * Working with unsafe pkg: unsafe.Pointer

18. Code generation [Homework #9](./hw09_struct_validator)

    * Working with go:generate
    * Useful libraries: impl, stringer, mockgen, easyjson
    * AST

19. Generics

    * Generics in Go 2.0
    * Working with generics in Go 1.0: code generation, type casting, empty interface

20. Profiling and debugging

    * Profiling and Debugging Go apps
    * Commands `go tool pprof` & `go tool trace`
    * Flame graphs

