class Solution(object):
    def spiralMatrixIII(self, rows, cols, rStart, cStart):
        if rows == 0 or cols == 0:
            return []
        def in_boundaries(r, c):
            return 0 <= r < rows and 0 <= c < cols
        count, total = 1, rows * cols
        res = [[rStart, cStart]]
        left, right = cStart - 1, cStart + 1
        top, bottom = rStart - 1, rStart + 1
        r, c = rStart, cStart
        direction = 'right'
        while count < total:
            if direction == 'right':
                while c < right:
                    c += 1
                    if in_boundaries(r, c):
                        res.append([r, c])
                        count += 1
                right += 1
                direction = 'down'
            elif direction == 'down':
                while r < bottom:
                    r += 1
                    if in_boundaries(r, c):
                        res.append([r, c])
                        count += 1
                bottom += 1
                direction = 'left'
            elif direction == 'left':
                while c > left:
                    c -= 1
                    if in_boundaries(r, c):
                        res.append([r, c])
                        count += 1
                left -= 1
                direction = 'up'
            else:
                while r > top:
                    r -= 1
                    if in_boundaries(r, c):
                        res.append([r, c])
                        count += 1
                top -= 1
                direction = 'right'
        return res


if __name__ == '__main__':
    solution = Solution()
    ans = solution.spiralMatrixIII(rows=5, cols=6, rStart=1, cStart=4)
    print(ans)

