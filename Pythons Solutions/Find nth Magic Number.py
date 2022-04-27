class Solution:
    def nthMagicNumber(self, a):
        ans, i = 0, 1

        while a > 0:
            i *= 5
            if a % 2 == 1:
                ans += i
            a = a // 2
        
        return ans

obj = Solution()
a = 3
ans = obj.nthMagicNumber(a)
print(ans)