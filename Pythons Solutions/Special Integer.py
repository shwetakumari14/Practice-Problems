class Solution:
    def maxHeighrOfStairCase(self, arr, b):
        start, sum, ans = 0, 0, len(arr)

        for end in range(len(arr)):
            sum += arr[end]
            while sum > b:
                sum -= arr[start]
                start += 1

                ans = min(ans, end-start+1)

                if sum == 0:
                    break
            
            if sum == 0:
                ans = -1
                break
        
        return ans

obj = Solution()
arr, b = [1, 2, 3, 4, 5], 10
ans = obj.maxHeighrOfStairCase(arr, b)
print(ans)