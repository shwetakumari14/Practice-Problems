class Solution:
    def fibbonacciTopDown(self, n):
        dp = [0]*(n+1)
        dp[1] = 1

        for i in range(2, n+1):
            dp[i] = dp[i-1] + dp[i-2]
        
        return dp[n]
    
    def fibbonacciBottomUp(self, n):
        ones, twos = 0, 1

        for i in range(2, n+1):
            temp = twos
            twos = ones + twos
            ones = temp
        
        return twos

obj = Solution()
ans = obj.fibbonacciBottomUp(6)
print(ans)