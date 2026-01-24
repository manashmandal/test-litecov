"""Tests for calculator module."""

import pytest
from calculator import add, subtract, multiply, divide


def test_add():
    assert add(1, 2) == 3
    assert add(-1, 1) == 0


def test_subtract():
    assert subtract(5, 3) == 2
    assert subtract(3, 5) == -2


def test_multiply():
    assert multiply(2, 3) == 6
    assert multiply(0, 5) == 0


def test_divide():
    assert divide(6, 2) == 3.0
    with pytest.raises(ValueError):
        divide(1, 0)


# Note: power and factorial are NOT tested - should show as uncovered
