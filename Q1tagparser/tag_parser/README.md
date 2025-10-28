# Tag Parser in Golang

This Go module implements a tag parser that accepts a comma-delimited list of tags and returns a clean slice of individual tags.

## Features

- Ignores preceding/trailing commas
- Trims whitespace around tags
- Removes empty tags
- Preserves internal tag phrases like `"C# byte code"`

## Example Inputs and Outputs

| Input | Output |
|-------|--------|
| `"golang"` | `["golang"]` |
| `"golang, python"` | `["golang", "python"]` |
| `"java, C#, python"` | `["java", "C#", "python"]` |
| `" , csharp , , python "` | `["csharp", "python"]` |
| `"C# byte code, python"` | `["C# byte code", "python"]` |

## How to Run

```bash
go test
```

## File Structure

- `parser.go`: Contains the `ParseTags` implementation.
- `parser_test.go`: Unit tests for different scenarios.
