class Solution(object):
    def minPathSum(self, grid):
        min_sum_graph = [[0 for _ in range(len(grid[0]) + 1)] for _ in range(len(grid) + 1)]
        r, c = len(grid), len(grid[0])
        # init edge
        # bottom edge
#        for i in range(c - 1, -1, -1):
#            for j in range(i + 1):
#                min_sum_graph[r - 1][j] += grid[r - 1][i]
#        # right edge
#        for i in range(r - 1, -1, -1):
#            for j in range(i + 1):
#                min_sum_graph[j][c - 1] += grid[i][c - 1]
#        min_sum_graph[r - 1][c - 1] -= grid[r - 1][c - 1]
#        for row in min_sum_graph:
#            print(row)
        r_curr, c_curr = r, c
        while r_curr != 0 and c_curr != 0:
            r_curr = r_curr - 1 if r_curr != 0 else 0
            c_curr = c_curr - 1 if c_curr != 0 else 0
            # up
            tmp = r_curr
            while tmp >= 0:
                if c_curr != c - 1:
                    min_sum_graph[tmp][c_curr] = grid[tmp][c_curr] + min(min_sum_graph[tmp + 1][c_curr], min_sum_graph[tmp][c_curr + 1])
                else:
                    min_sum_graph[tmp][c_curr] = grid[tmp][c_curr] + min_sum_graph[tmp + 1][c_curr]
                tmp -= 1
            # left
            tmp = c_curr
            while tmp >= 0:
                if r_curr != r - 1:
                    min_sum_graph[r_curr][tmp] = grid[r_curr][tmp] + min(min_sum_graph[r_curr + 1][tmp], min_sum_graph[r_curr][tmp + 1])
                else:
                    min_sum_graph[r_curr][tmp] = grid[r_curr][tmp] + min_sum_graph[r_curr][tmp + 1]
                tmp -= 1
        return min_sum_graph[0][0]

#    def minPathSum(self, grid):
#        sums = []
#        self.trace(grid, 0, sums)
#        return min(sums)
#
#    def trace(self, grid, acc, sums):
#        if len(grid) == 0 or len(grid[0]) == 0:
#            return
#        acc += grid[0][0]
#        if len(grid) == 1 and len(grid[0]) == 1:
#            sums.append(acc)
#            return
#        # go down
#        self.trace(grid[1:], acc, sums)
#        # go right
#        self.trace([row[1:] for row in grid], acc, sums)


if __name__ == '__main__':
    solution = Solution()
#    grid = [[0]]
#    grid = [[1, 2, 3]]
#    grid = [[1], [2], [3]]
#    grid = [[1, 2, 3], [4, 5, 6]]
#    grid = [[1, 3, 1], [1, 5, 1], [4, 2, 1]]
    grid = [[7, 1, 3, 5, 8, 9, 9, 2, 1, 9, 0, 8, 3, 1, 6, 6, 9, 5], [9, 5, 9, 4, 0, 4, 8, 8, 9, 5, 7, 3, 6, 6, 6, 9, 1, 6],
     [8, 2, 9, 1, 3, 1, 9, 7, 2, 5, 3, 1, 2, 4, 8, 2, 8, 8], [6, 7, 9, 8, 4, 8, 3, 0, 4, 0, 9, 6, 6, 0, 0, 5, 1, 4],
     [7, 1, 3, 1, 8, 8, 3, 1, 2, 1, 5, 0, 2, 1, 9, 1, 1, 4], [9, 5, 4, 3, 5, 6, 1, 3, 6, 4, 9, 7, 0, 8, 0, 3, 9, 9],
     [1, 4, 2, 5, 8, 7, 7, 0, 0, 7, 1, 2, 1, 2, 7, 7, 7, 4], [3, 9, 7, 9, 5, 8, 9, 5, 6, 9, 8, 8, 0, 1, 4, 2, 8, 2],
     [1, 5, 2, 2, 2, 5, 6, 3, 9, 3, 1, 7, 9, 6, 8, 6, 8, 3], [5, 7, 8, 3, 8, 8, 3, 9, 9, 8, 1, 9, 2, 5, 4, 7, 7, 7],
     [2, 3, 2, 4, 8, 5, 1, 7, 2, 9, 5, 2, 4, 2, 9, 2, 8, 7], [0, 1, 6, 1, 1, 0, 0, 6, 5, 4, 3, 4, 3, 7, 9, 6, 1, 9]]
    ans = solution.minPathSum(grid)
    print(ans)

