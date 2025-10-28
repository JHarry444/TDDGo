# Bank Management TDD

A simple bank account management system in Go, developed using Test-Driven Development (TDD) principles. Includes support for logging via an interface and uses GoMock for mocking in tests.

## Project Structure

```
go.mod
go.sum
bank/
    bank.go
    bank_test.go
    bank_unit_test.go
    mock_bank.go
```

### File Explanations

- **[go.mod](go.mod)**  
  Go module definition file. Specifies the module name and dependencies.

- **[go.sum](go.sum)**  
  Go checksum file for dependency verification.

- **[bank/bank.go](bank/bank.go)**  
  Main implementation file.  
  - Defines the `Account` struct and its methods: `Deposit`, `Withdraw`, and `Balance`.
  - Defines the `Logger` interface for logging actions.
  - Exposes the `ErrInsufficientFunds` error.

- **[bank/bank_test.go](bank/bank_test.go)**  
  Table-driven and mock-based tests for the `Account` type using GoMock.  
  - Tests logging behavior and correct balance updates.
  - Uses the generated mock for the `Logger` interface.

- **[bank/bank_unit_test.go](bank/bank_unit_test.go)**  
  Unit tests for the `Account` type without mocks.  
  - Tests basic deposit, withdrawal, and insufficient funds scenarios.

- **[bank/mock_bank.go](bank/mock_bank.go)**  
  Auto-generated GoMock file for the `Logger` interface.  
  - Do not edit manually.  
  - Used in tests to mock logging behavior.

## Setup Instructions

### 1. Install Dependencies

Ensure you have Go installed (version 1.22.2 or higher).

Install GoMock and mockgen:

```sh
go install github.com/golang/mock/mockgen@v1.6.0
go get github.com/golang/mock/gomock
```

### 2. Generate Mocks

If you need to regenerate the `mock_bank.go` file (for example, after changing the `Logger` interface), run:

```sh
mockgen -source=bank/bank.go -destination=bank/mock_bank.go -package=bank
```

### 3. Run Tests

To run all tests:

```sh
go test ./bank/... -v
```

Note: -v means in verbose mode, it provides detailed output, listing the names of all executed test functions and their execution times

## Summary

This project demonstrates a simple bank account system with TDD, including:
- Core business logic in [`bank/bank.go`](bank/bank.go)
- Unit and integration tests in [`bank/bank_unit_test.go`](bank/bank_unit_test.go) and [`bank/bank_test.go`](bank/bank_test.go)
- Mocking of the `Logger` interface using GoMock ([`bank/mock_bank.go`](bank/mock_bank.go))

Feel free to extend the functionality or add more tests!