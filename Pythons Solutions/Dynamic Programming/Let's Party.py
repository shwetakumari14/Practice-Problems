class Solution:
    def countPeopleTopDown(self, n):
        if n <= 2:
            return n
        dp = [0]*(n+1)
        dp[1], dp[2], mod = 1, 2, 10003

        for i in range(3, n+1):
            dp[i] = (dp[i-1] + dp[i-2]%mod * (i-1)%mod)%mod
        
        return dp[n]
    
    def countPeopleBottomUp(self, n):
        if n <= 2:
            return n

        ones, twos, mod = 1, 2, 10003

        for i in range(3, n+1):
            temp = twos
            twos =(twos + ones%mod*(i-1))%mod 
            ones = temp
        
        return twos

obj = Solution()
ans = obj.countPeopleTopDown(5)
print(ans)