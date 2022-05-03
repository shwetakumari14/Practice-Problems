class Solution:
    def singleNumber2(self, arr):
        ones, twos, threes = 0, 0, 0
        
        for i in range(len(arr)):
            twos |= (ones & arr[i])
            ones ^= arr[i]
            threes = ~(ones & twos)
            ones &= threes
            twos &= threes

        return ones

obj = Solution()
arr = [1, 2, 4, 3, 3, 2, 2, 3, 1, 1]
ans = obj.singleNumber2(arr)
print(ans)