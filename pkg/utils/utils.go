package utils

// Add is a utility function that returns the sum of two integers.
func Add(a, b int) int {
    return a + b
}

// Subtract is a utility function that returns the difference of two integers.
func Subtract(a, b int) int {
    return a - b
}

// Multiply is a utility function that returns the product of two integers.
func Multiply(a, b int) int {
    return a * b
}

// Divide is a utility function that returns the quotient of two integers.
// It returns an error if the divisor is zero.
func Divide(a, b int) (int, error) {
    if b == 0 {
        return 0, fmt.Errorf("division by zero")
    }
    return a / b, nil
}