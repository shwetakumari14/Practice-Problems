class Solution:
    def nCr(self, a, b, c):
        temp = []
        for r in range(0, a+1):
            temp.append([0 for c in range(0, b+1)])

        for i in range(a+1):
            j = 0
            while j <= min(i, b):
                if j == i or j == 0:
                    temp[i][j] = 1
                else:
                    temp[i][j] = (temp[i-1][j-1] % c + temp[i-1][j] % c) % c
                j += 1
                        
        return temp[a][b]%c

obj = Solution()
a, b, c = 5, 2, 13
ans = obj.nCr(a, b, c)
print(ans)