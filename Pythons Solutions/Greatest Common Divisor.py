class Solution:
    def gcd(self, a, b):

        if a < b:
            a, b = b, a 

        if b == 0:
            return a
        
        return self.gcd(b, a%b)

obj = Solution()
ans = obj.gcd(12,6)

print(ans)