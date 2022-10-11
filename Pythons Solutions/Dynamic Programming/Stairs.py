class Solution:
    def climbStairsTopDown(self, n):
        mod = 1000000007
        dp = [1]*(n+1)

        for i in range(2, n+1):
            dp[i] = dp[i-1] + dp[i-2]
        
        return dp[n]
    
    def climbStairsBottomUp(self, n):
        mod = 1000000007
        ones, twos = 1, 1

        for i in range(2, n+1):
            temp = twos
            twos = (ones + twos)%mod
            ones = temp
        
        return twos%mod

obj = Solution()
ans = obj.climbStairsBottomUp(5)
print(ans)