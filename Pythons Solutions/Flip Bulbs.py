class Solution:
    def flipBulbs(self, str):
        temp = [0]*len(str)
        ans = [-1]*2

        for i in range(len(str)):
            if str[i] == "0":
                temp[i] = 1
            elif str[i] == "1":
                temp[i] = -1

        k, max_so_far, max_till_now = 0, 0, 0

        for i in range(len(temp)):
            if max_till_now +  temp[i] < 0:
                k = i+1
                max_till_now = 0
            else:
                max_till_now += temp[i]
            
            if max_so_far < max_till_now:
                max_so_far = max_till_now
                ans[0] = k
                ans[1] = i
        
        if ans[0] == -1:
            return []
        ans[0] += 1
        ans[1] += 1
        return ans

obj = Solution()
str = "010"
ans = obj.flipBulbs(str)
print(ans)