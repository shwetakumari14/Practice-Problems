class Solution:
    def nextPermutation(self, arr):
        breakPoint = 0
        for i in range(len(arr)-1, -1, -1):
            if arr[i] > arr[i-1]:
                breakPoint = i-1
                break

        for i in range(len(arr)-1, -1, -1):
            if arr[i] > arr[breakPoint]:
                arr[i],arr[breakPoint] = arr[breakPoint], arr[i]
                break

        breakPoint +=1
        n = len(arr)-1
        while breakPoint < n:
            arr[breakPoint], arr[n] = arr[n], arr[breakPoint]
            breakPoint+=1
            n-=1

        return arr
        
        

obj = Solution()
arr = [1, 5, 8, 4, 7, 3, 5, 6, 1]
ans = obj.nextPermutation(arr)
print(ans)