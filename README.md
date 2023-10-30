# Readme - Explanation of the Quad Go Program

The Go program `QuadE` is designed to generate a pattern of characters that resembles a rectangular shape. The dimensions of the rectangle are determined by the values of `x` and `y`. The program uses the `github.com/01-edu/z01` package to interact with the console and print characters.

## Function Description

The `QuadE` function takes two integer parameters, `x` and `y`, representing the width and height of the rectangle. It then generates and prints a pattern based on these dimensions.

## How the Program Works

1. The program checks if both `x` and `y` are greater than 0. If either of them is not greater than 0, it will not generate the pattern.

2. The program prints the first line of the rectangle. It consists of the character 'A' repeated `x-2` times (if `x` is greater than 1), followed by 'C' if `x` is greater than 1. It ends with a newline character.

3. The program proceeds to print the middle lines. For each line in the middle (excluding the first and last lines), it starts with 'B', followed by `x-2` spaces (if `x` is greater than 1), and ends with 'B' (followed by a newline character if `x` is greater than 1).

4. After printing the middle lines `y-2` times, it moves on to the last line. If `y` is greater than 1, it prints 'C', followed by `x-2` 'B' characters (if `x` is greater than 1), and finally 'A' (followed by a newline character if `x` is greater than 1).

5. The pattern is now complete, and the function terminates.

## Example

Here's an example of how the `QuadE` function works:

```go
QuadE(5, 4)
```

Output:
```
ABC
B B
B B
ABA
```

In this example, `x` is 5, and `y` is 4, so the program generates a 5x4 rectangular pattern with 'A', 'B', and 'C' characters as shown.

Please note that the `QuadE` function generates a rectangle only if both `x` and `y` are greater than 0. If either of them is not, it won't print anything.

This readme provides an overview of the `QuadE` Go program and how it generates rectangular patterns based on input dimensions.
