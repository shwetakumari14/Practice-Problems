class Solution:
    def singleNumber(self, arr):
        ans = arr[0]

        for i in range(1, len(arr)):
            ans = ans ^ arr[i]

        return ans

obj = Solution()
arr = [1, 2, 2, 3, 1]
ans = obj.singleNumber(arr)
print(ans)