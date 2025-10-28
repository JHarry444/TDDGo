
# Shipping Cost Calculator (TDD & Mutation Testing)

## Description
This example demonstrates a real-world shipping cost calculator.

### Rules:
- Orders >= 1000 → Free shipping (₹0)
- Orders < 1000 → ₹50 flat shipping
- Orders < 0 → Invalid input (-1)

## Files
- `shipping.go`: Implementation
- `shipping_test.go`: Basic test cases
- Add more tests after running mutation testing

## How to Run
```bash
go test ./...
gremlins run
```
