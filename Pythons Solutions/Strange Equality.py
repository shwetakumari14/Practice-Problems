import math

class Solution:
    def findSpecialXOR(self, a):
        numOfBits = int(math.log2(a) + 1)
        x, y = 0, 0

        for i in range(numOfBits):
            if a & (1 << i) != 0:
                continue
            x += 1 << i
        
        y = 1 << numOfBits

        return x ^ y

obj = Solution()
a = 5
ans = obj.findSpecialXOR(5)
print(ans)