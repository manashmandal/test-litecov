"""Utility functions - this file has NO tests at all."""


def max_value(a: int, b: int) -> int:
    """Return the maximum of two values."""
    if a > b:
        return a
    return b


def min_value(a: int, b: int) -> int:
    """Return the minimum of two values."""
    if a < b:
        return a
    return b


def abs_value(n: int) -> int:
    """Return the absolute value."""
    if n < 0:
        return -n
    return n


def clamp(value: int, min_val: int, max_val: int) -> int:
    """Clamp value between min and max."""
    if value < min_val:
        return min_val
    if value > max_val:
        return max_val
    return value


def is_even(n: int) -> bool:
    """Check if a number is even."""
    return n % 2 == 0


def is_odd(n: int) -> bool:
    """Check if a number is odd."""
    return n % 2 != 0
