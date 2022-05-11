class Solution:
    def findElementInRotatedArray(self, arr, b):
        low, high = 0, len(arr)-1

        while low < high:
            mid = low + (high-low)//2
            if arr[mid] >= arr[low]:
                if arr[mid] < b or b < arr[low]:
                    low = mid + 1
                else:
                    high = mid
            else:
                if arr[mid] >= b or b >= arr[low]:
                    low = mid + 1
                else:
                    high = mid
        
        if arr[low] == b:
            return low
        
        return -1

obj = Solution()
arr, b = [4, 5, 6, 7, 0, 1, 2, 3], 2
ans = obj.findElementInRotatedArray(arr, b)
print(ans)