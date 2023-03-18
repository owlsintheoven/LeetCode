class Solution(object):
    size = 0
    count = 0
    def totalNQueens(self, n):
        self.size = n
        count = 0
        m = [value for value in range(n)]
        self.helper(m, 0, [])
        return self.count

    def helper(self, m, start_c, placed):
        if len(placed) != 0:
            if start_c == self.size:
                self.count += 1
        for j in range(len(m)):
            newly_placed = start_c + m[j] * self.size
            valid = True
            for prev_place in placed:
                if (prev_place > newly_placed and (newly_placed - prev_place) % (self.size - 1) == 0) \
                    or (prev_place < newly_placed and (newly_placed - prev_place) % (self.size + 1) == 0):
                    valid = False
                    break
            if valid:
                new_m = m[:j]
                new_m.extend(m[j + 1:])
                placed.append(newly_placed)
                self.helper(new_m, start_c + 1, placed)
                placed.pop(-1)


if __name__ == '__main__':
    solution = Solution()
    n = 4
    ans = solution.totalNQueens(n)
    print(ans)