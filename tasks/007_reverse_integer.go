package tasks

/*
Reverse Integer
Difficulty: Medium

Given a signed 32-bit integer x, return x with its digits reversed.
If reversing x causes the value to go outside the signed 32-bit integer range [-2^31, 2^31 - 1],
then return 0.

Assume the environment does not allow you to store 64-bit integers (signed or unsigned).

Example 1:
  Input:  x = 123
  Output: 321

Example 2:
  Input:  x = -123
  Output: -321

Example 3:
  Input:  x = 120
  Output: 21

Constraints:
  - -2^31 <= x <= 2^31 - 1
*/

func Reverse(x int) int {
	const INT_MAX = 1<<31 - 1
	const INT_MIN = -1 << 31

	result := 0

	for x != 0 {
		digit := x % 10
		x /= 10

		// Check for overflow before multiplying by 10
		if result > INT_MAX/10 || (result == INT_MAX/10 && digit > 7) {
			return 0
		}
		if result < INT_MIN/10 || (result == INT_MIN/10 && digit < -8) {
			return 0
		}

		result = result*10 + digit
	}

	return result
}
