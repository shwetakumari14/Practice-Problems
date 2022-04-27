class Solution:
    def stringRank(self, str):
        ans, mod, n = 0, 1000003, len(str)

        for i in range(n):
            count, j = 0, i+1
            while j<n:
                if ord(str[j]) < ord(str[i]):
                    count += 1
                j += 1
            
            ans += (count * self.fib(n-1-i))%mod
            
        
        return (ans + 1)%mod
    
    def fib(self, n):
        if n <=1:
            return 1
        return self.fib(n-1) + self.fib(n-2)

obj = Solution()
str = "acb"
ans = obj.stringRank(str)
print(ans)