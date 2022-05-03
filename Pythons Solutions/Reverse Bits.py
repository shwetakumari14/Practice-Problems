class Solution:
    def reverseBits(self, num):
        ans = 0

        for i in range(32):
            ans = (ans << 1) + (num & 1)
            num >>= 1
        return ans

obj = Solution()
num = 10000000000000000000000000000011
ans = obj.reverseBits(num)
print(ans)