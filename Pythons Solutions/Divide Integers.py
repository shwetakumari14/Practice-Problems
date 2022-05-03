class Solution:
    def divideIntegers(self, a, b):
        sin, ptr = 1, 0

        if a < 1:
            sin = -sin
            a = a*(-1)
        if b < 1:
            sin = -sin
            b = b*(-1)

        for i in range(31, -1, -1):
            if (b << i) <= a:
                a -= (b << i)
                ptr += 1 << i
        
        if sin < 1:
            ptr = -ptr

        return ptr

obj = Solution()
a, b = 5, 2
ans = obj.divideIntegers(a, b)
print(ans)