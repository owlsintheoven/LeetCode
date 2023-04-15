class Solution(object):
    def generateMatrix(self, n):
        res = [[0 for _ in range(n)] for _ in range(n)]
        num = 1
        r, c = 0, 0
        direction = 'right'
        while num <= pow(n, 2):
            if direction == 'right':
                while c < n and res[r][c] == 0:
                    res[r][c] = num
                    c += 1
                    num += 1
                c -= 1
                r += 1
                direction = 'down'
            elif direction == 'down':
                while r < n and res[r][c] == 0:
                    res[r][c] = num
                    r += 1
                    num += 1
                r -= 1
                c -= 1
                direction = 'left'
            elif direction == 'left':
                while c >= 0 and res[r][c] == 0:
                    res[r][c] = num
                    c -= 1
                    num += 1
                c += 1
                r -= 1
                direction = 'up'
            else:
                while r >= 0 and res[r][c] == 0:
                    res[r][c] = num
                    r -= 1
                    num += 1
                r += 1
                c += 1
                direction = 'right'
        return res


if __name__ == '__main__':
    solution = Solution()
    ans = solution.generateMatrix(3)
    print(ans)

