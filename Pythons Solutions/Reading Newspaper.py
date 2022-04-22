class Solution:
    def readNews(self, b, a):
        sum,i = 0,0

        while sum < a:
            sum+= b[i]
            i+=1
            i = i%7
        
        if i == 0:
            return 7
        return i
       

obj = Solution()
b, a = [15, 20, 20, 15, 10, 30, 45], 100
ans = obj.readNews(b, a)
print(ans)