class Solution:
    def findSingleElement(self, arr):
        if len(arr) == 1:
            return arr[0]
        low, high = 0, len(arr)-1

        while low < high-2:
            mid = low + (high-low)//2
            if arr[mid] == arr[mid+1]:
                mid += 1
            
            if(high -mid) %2 == 1:
                low = mid+1
            else:
                high = mid

        
        if arr[low] == arr[low+1]:
            return arr[high]
        return arr[low]

obj = Solution()
arr = [1, 1, 2, 3, 3, 4, 4]
ans = obj.findSingleElement(arr)
print(ans)