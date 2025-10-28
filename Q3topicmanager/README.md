# Topic Score Tracker in Go (TDD)

This project follows a Test-Driven Development (TDD) approach to build a system that:

1. Finds the highest number in an array.
2. Computes the top score for each topic.
3. Writes the result to a JSON file.

## Structure

- `highest.go`: Contains logic to find the highest number from a slice.
- `topic.go`: Computes highest score per topic using `highest.go`.
- `writer.go`: Persists results to a file.
- `main.go`: Demonstrates usage.

## How to Run

```bash
go mod init example.com/topic
go run main.go
```

## Output

Creates `top_scores.json` with content like:

```json
[
  {
    "Name": "Physics",
    "Score": 89
  },
  {
    "Name": "Art",
    "Score": 87
  },
  {
    "Name": "Comp Sci",
    "Score": 97
  }
]
```

## TDD Steps

1. Write a failing test (e.g. find highest number)
2. Implement minimal logic to make it pass.
3. Refactor to improve code quality.
4. Repeat.

## Code Smells Addressed

- **Long methods**: Split logic across modules.
- **Duplication**: Used reusable `FindHighest` method.
- **Tight coupling**: Logic is decoupled across responsibilities.
- **No SRP (Single Responsibility Principle)**: Followed clean separation.

