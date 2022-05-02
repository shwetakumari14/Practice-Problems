class Solution:
    def columnTitle(self, n):
        ans = ""

        while n > 0:
            ans = chr((n-1)%26 + 65) + ans
            n = (n-1)//26

        return ans

obj = Solution()
n = 27
ans = obj.columnTitle(n)
print(ans)