class Solution:
    def totalIntersectionPoints(self, a, b):
        x = a*(a-1)//2
        y = b*(b-1) + 2*a*b
                        
        return x + y

obj = Solution()
a, b = 2, 2
ans = obj.totalIntersectionPoints(a, b)
print(ans)