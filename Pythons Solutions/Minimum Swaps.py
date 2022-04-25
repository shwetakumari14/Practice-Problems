from turtle import right


class Solution:
    def minSwaps(self, arr, b):
        minCount = 0

        for i in range(len(arr)):
            if arr[i] < b:
                minCount +=1
        
        if minCount <= 1:
            return 0
        else:
            rightCount, leftCount, count = 0, 0, 0
            while rightCount < minCount:
                if arr[rightCount] > b:
                    count += 1
                rightCount += 1
            ans = count
            while rightCount < len(arr):
                if arr[rightCount] > b:
                    count += 1
                if arr[leftCount] > b:
                    count -= 1
                
                ans = min(ans, count)

                rightCount += 1
                leftCount += 1

        return ans

obj = Solution()
arr, b = [1, 12, 10, 3, 14, 10, 5], 8
ans = obj.minSwaps(arr, b)
print(ans)