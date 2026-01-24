const { add, subtract, multiply, divide } = require("./calculator");

describe("Calculator", () => {
  describe("add", () => {
    test("adds two positive numbers", () => {
      expect(add(1, 2)).toBe(3);
    });

    test("adds negative and positive", () => {
      expect(add(-1, 1)).toBe(0);
    });
  });

  describe("subtract", () => {
    test("subtracts two numbers", () => {
      expect(subtract(5, 3)).toBe(2);
    });
  });

  describe("multiply", () => {
    test("multiplies two numbers", () => {
      expect(multiply(2, 3)).toBe(6);
    });
  });

  describe("divide", () => {
    test("divides two numbers", () => {
      expect(divide(6, 2)).toBe(3);
    });

    test("throws on division by zero", () => {
      expect(() => divide(1, 0)).toThrow("Cannot divide by zero");
    });
  });

  // Note: power and factorial are NOT tested - should show as uncovered
});
