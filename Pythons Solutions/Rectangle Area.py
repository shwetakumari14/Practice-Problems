import re


class Solution:
    def area(self, a, b, c, d, e, f, g, h):
        x = min(c,g) - max(a,e)
        y = min(d,h) - max(b,f)

        ans = x*y
        if x <= 0 or y <= 0:
            ans = 0

        return ans

obj = Solution()
a, b, c, d, e, f, g, h = 0, 0, 4, 4, 2, 2, 6, 6
ans = obj.area(a, b, c, d, e, f, g, h)
print(ans)