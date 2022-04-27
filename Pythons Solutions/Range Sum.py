mod = 1000000007

def matrixMul(m1, m2):
    res = []
    for r in range(0, 2):
        res.append([0 for c in range(0, 2)])
    
    for i in range(2):
        for j in range(2):
            for k in range(2):
                res[i][j] = (res[i][j] + m1[i][k] * m2[k][j]) % mod
                res[i][j] = res[i][j] % mod
    
    return res

def power(x, y):
    res, half = [], []
    for r in range(0, 2):
        res.append([0 for c in range(0, 2)])
        half.append([0 for c in range(0, 2)])

    if y == 0:
        res[0][0] = 1
        res[0][1] = 0
        res[1][0] = 0
        res[1][1] = 1

        return res
    if y == 1:
        return x

    half = power(x, y//2)
    half = matrixMul(half, half)
    if y % 2 == 1:
        half = matrixMul(x, half)
    
    return half


def fibonacci(n):
    if n == 1 or n == 2:
        return 1
    
    matrix, mat = [], []
    for r in range(0, 2):
        matrix.append([0 for c in range(0, 2)])
        mat.append([0 for c in range(0, 2)])

    matrix[0][0] = 1
    matrix[0][1] = 1
    matrix[1][0] = 1

    mat = power(matrix, n-1)
    return mat[0][0]

class Solution:
    def sumOfRange(self, a, b):
        fibA = fibonacci(a+1)
        fibB = fibonacci(b+2)

        res = fibB - fibA

        if res < 0:
            res += mod

        return res % mod

obj = Solution()
a, b = 0, 3
ans = obj.sumOfRange(a, b)
print(ans)