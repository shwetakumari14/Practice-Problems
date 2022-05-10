class Solution:
    def peakElement(self, arr):
        if len(arr)==1:
            return arr[0]

        if arr[0] >= arr[1]:
            return arr[0]
        
        if arr[len(arr)-1] >= arr[len(arr)-2]:
            return arr[len(arr)-1]

        low, high = 0, len(arr)-2

        while low <= high:
            mid = low + (high-low)//2

            if arr[mid]>=arr[mid-1] and arr[mid] >= arr[mid+1]:
                return arr[mid]
            elif arr[mid] > arr[mid-1]:
                low = mid + 1
            else:
                high = mid - 1
        
        return -1

obj = Solution()
arr = [1, 2, 3, 4, 5]
ans = obj.peakElement(arr)
print(ans)