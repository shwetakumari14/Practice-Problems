class Solution:
    def firstMissingPositive(self, arr):

        for i in range(len(arr)):
            if arr[i] >=1 and arr[i] < len(arr):
                if arr[i] != i+1 and arr[arr[i]-1] != arr[i]:
                    arr[arr[i]-1], arr[i] = arr[i], arr[arr[i]-1]
                    
                    if arr[i] >=1 and arr[i] < len(arr):
                        if arr[i] != i+1 and arr[arr[i]-1] != arr[i]:
                            i -= 1
        
        for i in range(len(arr)):
            if i >= 0 and i < len(arr):
                if arr[i] != i+1:
                    return i+1


        return len(arr)+1

obj = Solution()
arr = [3, 4, -1, 1]
ans = obj.firstMissingPositive(arr)
print(ans)