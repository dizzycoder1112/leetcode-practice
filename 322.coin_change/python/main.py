from typing import List
class Solution:
    def coinChange(self, coins: List[int], amount: int) -> int:
        dp = [float('inf')] * (amount+1)
        dp[0] = 0

        for coin in coins:
            for i in range(coin, amount+1):
                dp[i] = min(dp[i], dp[i-coin]+1)

        if dp[amount] == float('inf'):
            return -1
        else:
            return dp[amount]


if __name__ == "__main__":
    sol = Solution()

    # Test 1: coins=[1,2,5], amount=11 -> 5+5+1 = 3枚
    print(f"Test 1: {sol.coinChange([1, 2, 5], 11)}")  # Expected: 3

    # Test 2: coins=[2], amount=3 -> 无法凑成
    print(f"Test 2: {sol.coinChange([2], 3)}")  # Expected: -1

    # Test 3: coins=[1], amount=0 -> 0枚
    print(f"Test 3: {sol.coinChange([1], 0)}")  # Expected: 0

    # Test 4: coins=[1], amount=1 -> 1枚
    print(f"Test 4: {sol.coinChange([1], 1)}")  # Expected: 1

    # Test 5: coins=[1,2,5], amount=100 -> 20枚 (5*20)
    print(f"Test 5: {sol.coinChange([1, 2, 5], 100)}")  # Expected: 20