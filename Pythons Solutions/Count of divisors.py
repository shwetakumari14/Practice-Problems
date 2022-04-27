class Solution:
    def countDivisors(self, arr):
        n = 1000005
        divisor = [0]*n

        for i in range(n):
           divisor[i] = i
        
        for i in range(2, n):
            if divisor[i] == i:
                j = i*i
                while j <n:
                    if divisor[j] == j:
                        divisor[j] = i
                    j += i
        
        solution = []

        for i in range(len(arr)):
            temp = arr[i]
            ans = 1
            while temp != 1:
                count = 1
                div = divisor[temp]
                while temp != 1 and temp%div == 0:
                    count += 1
                    temp = temp // div
                
                ans *= count
            solution.append(ans)
        
        return solution

       

obj = Solution()
arr = [2, 2, 3, 2]
ans = obj.countDivisors(arr)
print(ans)