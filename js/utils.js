/**
 * Utility functions - this file has NO tests at all.
 */

function max(a, b) {
  if (a > b) {
    return a;
  }
  return b;
}

function min(a, b) {
  if (a < b) {
    return a;
  }
  return b;
}

function abs(n) {
  if (n < 0) {
    return -n;
  }
  return n;
}

function clamp(value, minVal, maxVal) {
  if (value < minVal) {
    return minVal;
  }
  if (value > maxVal) {
    return maxVal;
  }
  return value;
}

function isEven(n) {
  return n % 2 === 0;
}

function isOdd(n) {
  return n % 2 !== 0;
}

module.exports = {
  max,
  min,
  abs,
  clamp,
  isEven,
  isOdd,
};
