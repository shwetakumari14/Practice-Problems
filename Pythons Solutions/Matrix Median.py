class Solution:
    def matrixMedian(self, arr):
        low, high = 0, 1000000000
        medianPos = len(arr)*(len(arr[0]))//2
        ans = -1

        while low <= high:
            mid = low + (high-low)//2
            count = 0

            for i in range(len(arr)):
                count += self.countEleLessThanMedian(arr[i], mid)
            
            if count > medianPos:
                high = mid - 1
            else:
                ans = mid
                low = mid + 1
        
        return ans
    
    def countEleLessThanMedian(self, array, val):
        low, high, ans = 0, len(array)-1, -1

        while low <= high:
            mid = low + (high-low)//2
            if array[mid] < val:
                ans = mid
                low = mid + 1
            else:
                high = mid - 1
        
        return ans + 1

obj = Solution()
arr = [[1, 3, 5],[2, 6, 9],[3, 6, 9]]
ans = obj.matrixMedian(arr)
print(ans)