import re


class Solution:
    def minJumps(self, arr):
        if len(arr) <= 1:
            return 0
        
        if arr[0] == 0:
            return -1
        
        jumps, maxReached, steps = 1, arr[0], arr[0]

        for i in range(1, len(arr)):
            if i == len(arr)-1:
                return jumps
            
            maxReached = max(maxReached, arr[i]+i)
            steps -= 1
            if steps == 0:
                jumps +=1
                if i >= maxReached:
                    return -1
                steps = maxReached - i

        return jumps

obj = Solution()
arr = [1, 2, 3, 4, 5]
ans = obj.minJumps(arr)
print(ans)