# grug iter

Golang iterators for mediocre programmers like me.

We implement a pull iterator, drawing heavy inspiration from Rust.
The scaffolding in `grug.go` is less than 70 lines of code.

### Features

- Custom iterators on structs
- Ability to map functions onto iterators
- Ability to filter elements from an iterator

### Getting Started

It's simple, implement ~a trait~ an interface and use `grug.NewIterator`
around your struct! Check out the examples.

### Examples

- `fibonacci`
- `reference_item`

Run the examples using

```sh
go run ./examples/EXAMPLE_NAME
```

like

```sh
go run ./examples/fibonacci
```
