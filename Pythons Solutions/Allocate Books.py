class Solution:
    def allocateBooks(self, arr, b):
        pageSum, n = 0, len(arr)

        for num in arr:
            pageSum += num
        
        low, high, ans = max(arr), pageSum, pageSum

        while low <= high:
            mid = low + (high - low)//2
            if self.isAllocationPossible(arr, b, mid):
                ans = min(ans, mid)
                high = mid - 1
            else:
                low = mid + 1
        
        return ans
        
    def isAllocationPossible(self,arr, b, curMin):
        currSum, n, cnt = 0, len(arr), 1

        for i in range(n):
            if currSum + arr[i] > curMin:
                cnt += 1
                currSum = arr[i]

                if cnt > b:
                    return 0
            else:
                currSum += arr[i]
        
        return 1
        

obj = Solution()
arr, b = [12, 34, 67, 90], 2
ans = obj.allocateBooks(arr, b)
print(ans)