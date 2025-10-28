# Highest Score Finder in Golang (TDD Approach)

This module implements logic to find the highest score from an array of integers using Test-Driven Development.

## Features

- Returns the highest value in a non-empty array
- Returns an error if the array is empty
- Supports edge cases like duplicate highest scores or single-element arrays

## Example Inputs and Outputs

| Input                     | Output |
|--------------------------|--------|
| `[4, 5, -8, 3, 11, -21]`  | `11`   |
| `[]`                     | `error`|
| `[7, 13]`                | `13`   |
| `[13, 4]`                | `13`   |
| `[3, 5, 5, 2]`           | `5`    |

## How to Run

```bash
go test
```

## Files

- `highest.go`: contains the function to find the highest score.
- `highest_test.go`: unit tests using `testing` and table-driven tests.
