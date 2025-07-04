# Go DSA Patterns for Code Interviews

This repository is a collection of common Data Structures and Algorithms (DSA) patterns implemented in Go. It's designed to be a resource for practicing for technical interviews.

## Project Structure

The project is organized by pattern, with each pattern residing in its own package. This separation makes it easy to focus on one pattern at a time.

Each package contains two main files:
- `pattern_name.go`: The implementation of the pattern.
- `pattern_name_test.go`: The unit tests for the implementation.

### Available Patterns

-   **Two Pointers**: Problems that can be solved by iterating with two pointers, often from opposite ends of a data structure.
    -   Location: `/twopointers`

## Usage

To run all the tests for all patterns, you can execute the following command from the root of the project:

```sh
go test -v ./...
```

To run tests for a specific package, navigate to that directory or specify the path:

```sh
go test -v ./twopointers
```
