# Walkthrough of Tests and Code Smells

## Tests Overview

### 1. highest_test.go
- **Purpose**: Verifies logic in `FindHighest` to return the largest number in an integer array.
- **Covers**:
  - Normal cases
  - Edge cases (empty slice, single item)
  - Multiple highest values
  - Different orderings

### 2. topic_test.go
- **Purpose**: Tests `FindHighestPerTopic` to ensure correct topic-score pairs are generated.
- **Validates**:
  - Correct mapping of topics to top scores
  - Accurate handling of multiple topics

### 3. writer_test.go
- **Purpose**: Tests the `WriteResultsToFile` function.
- **Checks**:
  - File is created correctly
  - Results are written in JSON format
  - File cleanup handled post-test

---

## Code Smells Identified & Avoided

1. **Tight Coupling**
   - Separated core logic (highest, topic) from I/O (writer).
   - Each function/module does one thing (SRP).

2. **Magic Values**
   - All inputs are parameterized in test cases. No hardcoded behavior.

3. **Lack of Error Handling**
   - Tests validate and handle expected errors (e.g., empty input array).

4. **Poor Naming**
   - Used clear function and struct names like `FindHighest`, `Result`, and `Topic`.

5. **Test Duplication**
   - Used table-driven tests to avoid repetition in `highest_test.go`.

6. **Unclear Output**
   - JSON output is tested in writer tests, ensuring usability.

---

## How to Run Tests

From project root, run:

```
go test ./topic/...
```

Ensure you’re in a folder with `go.mod` initialized.