class Solution:
    def minXOR(self, arr):
        arr.sort()
        ans = 1000000

        for i in range(len(arr)-1):
            min = arr[i]^arr[i+1]
            if ans > min:
                ans = min
        return ans

obj = Solution()
arr = [0, 4, 7, 9]
ans = obj.minXOR(arr)
print(ans)