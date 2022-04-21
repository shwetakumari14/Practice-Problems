import re


class Solution:
    def trailingZeros(self, num):
        ans, i, div = 0, 5, num/5

        while div > 0:
            ans += div
            i *= 5
            div = num/i
        
        return int(ans)

obj = Solution()
print(obj.trailingZeros(25))