import sys

class Solution:
    def maxDiff(sefl, arr):
        max1, min1, max2, min2 = -1000000, 1000000, -1000000, 1000000

        for i in range(len(arr)):
            if max1 < arr[i]+i:
                max1 = arr[i]+i

            if min1 > arr[i]+i:
                min1 = arr[i]+i

            if max2 < arr[i]-i:
                max2 = arr[i]-i
            
            if min2 > arr[i]-i:
                min2 = arr[i]-i

        return max((max1-min1), (max2-min2))

obj = Solution()
arr = [1, 3, -1]
ans = obj.maxDiff(arr)
print(ans)