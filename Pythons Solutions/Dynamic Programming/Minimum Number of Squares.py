class Solution:
    def noOfSquaresTopDown(self, n):
        if n == 0:
            return 0
        dp = [0]*(n+1)
        dp[1] = 1

        sqrs = [i*i for i in range(330)]

        for i in range(2, n+1):
            for x in sqrs:
                if x > i:
                    break
                dp[i] = min(i, 1 + dp[i - x])
        
        return dp[n]

obj = Solution()
ans = obj.noOfSquaresTopDown(4)
print(ans)