class Solution:
    def allocateBooks(self, arr, b):
        n, sum = len(arr), 0
        if n < b:
            return -1
        
        for i in range(len(arr)):
            sum += arr[i]
        
        start, end, ans = 0, sum, 1000000000
        
        for i in range(len(arr)):
            mid = start + (end - start)//2
            if self.isAllocationPossible(arr, b, mid):
                ans = min(ans, mid)
                end = mid - 1
            else:
                start = mid + 1

        return ans
    
    def isAllocationPossible(self,arr, b, curMin):
        studentReq, currSum = 1, 0

        for i in range(len(arr)):
            if arr[i] > curMin:
                return False
            if currSum + arr[i] > curMin:
                studentReq += 1
                currSum = arr[i]
                if studentReq > b:
                    return False
            else:
                currSum += arr[i]

        
        return True

obj = Solution()
arr, b = [12, 34, 67, 90], 2
ans = obj.allocateBooks(arr, b)
print(ans)