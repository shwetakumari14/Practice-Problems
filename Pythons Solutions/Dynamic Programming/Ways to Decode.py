class Solution:
    def waysToDecode(self, str):
        n = len(str)
        if n == 0:
            return 0
        
        if str[0] == '0':
            return 0

        
        dp = [0]*(n+1)
        dp[0], dp[1] = 1, 1

        for i in range(2, n+1):
            if str[i-1] >= '1' and str[i-1] <= '9':
                dp[i] = dp[i-1]
            if str[i-2] == '1':
                dp[i] += dp[i-2]
            elif str[i-2] == '2' and str[i-1] >= '0' and str[i-1] <= '6':
                dp[i] += dp[i-2]
        
        return dp[n]

obj = Solution()
str = "4126"
ans = obj.waysToDecode(str)
print(ans)