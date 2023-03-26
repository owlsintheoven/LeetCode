class Solution(object):
    def numIslands(self, grid):
        if not grid:
            return 0
        islands = 0
        for i in range(len(grid)):
            for j in range(len(grid[0])):
                if grid[i][j] == '1':
                    self.dfs(grid, i, j)
                    islands += 1
        return islands

    def dfs(self, grid, r, c):
        if r < 0 or r == len(grid) or c < 0 or c == len(grid[0]):
            return
        if grid[r][c] != '1':
            return
        grid[r][c] = '#'
        self.dfs(grid, r - 1, c)
        self.dfs(grid, r + 1, c)
        self.dfs(grid, r, c - 1)
        self.dfs(grid, r, c + 1)

#    def numIslands(self, grid):
#        r = len(grid)
#        c = len(grid[0])
#        adj = [[] for _ in range(r * c)]
#        idx = -1
#        for i in range(r):
#            for j in range(c):
#                idx += 1
#                if grid[i][j] == '0':
#                    continue
#                # up
#                if i != 0 and grid[i - 1][j] == '1':
#                    adj[idx].append(idx - c)
#                # down
#                if i != r - 1 and grid[i + 1][j] == '1':
#                    adj[idx].append(idx + c)
#                # left
#                if j != 0 and grid[i][j - 1] == '1':
#                    adj[idx].append(idx - 1)
#                # right
#                if j != c - 1 and grid[i][j + 1] == '1':
#                    adj[idx].append(idx + 1)
#        visited = [False for _ in range(r * c)]
#        islands = 0
#        idx = -1
#        for i in range(r):
#            for j in range(c):
#                idx += 1
#                if grid[i][j] == '1' and not visited[idx]:
#                    self.dfs(adj, visited, idx)
#                    islands += 1
#        return islands
#
#
#    def dfs(self, adj, vistied, src):
#        vistied[src] = True
#        for i in adj[src]:
#            if not vistied[i]:
#                self.dfs(adj, vistied, i)


if __name__ == '__main__':
    solution = Solution()
#    grid = [
#        ["1", "1", "1", "1", "0"],
#        ["1", "1", "0", "1", "0"],
#        ["1", "1", "0", "0", "0"],
#        ["0", "0", "0", "0", "0"]
#    ]
    grid = [
        ["1", "1", "0", "0", "0"],
        ["1", "1", "0", "0", "0"],
        ["0", "0", "1", "0", "0"],
        ["0", "0", "0", "1", "1"]
    ]
    ans = solution.numIslands(grid)
    print(ans)

