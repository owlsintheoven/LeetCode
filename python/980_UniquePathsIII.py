class Solution(object):
    def uniquePathsIII(self, grid):
        def search(visited, curr_path, worked_paths):
            r, c = curr_path[-1]
            if len(curr_path) == goal_step:
                if (c == goal_c and abs(r-goal_r)) == 1 or (r == goal_r and abs(c-goal_c) == 1):
                    worked_paths.append(curr_path.copy())
            visited[r][c] = True
            if r > 0 and not visited[r-1][c] and grid[r-1][c] == 0:
                curr_path.append([r-1, c])
                search(visited, curr_path, worked_paths)
                curr_path.pop()
            if r < rows - 1 and not visited[r+1][c] and grid[r+1][c] == 0:
                curr_path.append([r+1, c])
                search(visited, curr_path, worked_paths)
                curr_path.pop()
            if c > 0 and not visited[r][c-1] and grid[r][c-1] == 0:
                curr_path.append([r, c-1])
                search(visited, curr_path, worked_paths)
                curr_path.pop()
            if c < cols - 1 and not visited[r][c+1] and grid[r][c+1] == 0:
                curr_path.append([r, c+1])
                search(visited, curr_path, worked_paths)
                curr_path.pop()
            visited[r][c] = False

        rows, cols = len(grid), len(grid[0])
        worked_paths = []
        visited = [[False for _ in range(cols)] for _ in range(rows)]
        init_path = []
        goal_step = 1
        for i in range(rows):
            for j in range(cols):
                if grid[i][j] == 1:
                    visited[i][j] = True
                    init_path.append([i, j])
                elif grid[i][j] == 2:
                    goal_r, goal_c = i, j
                elif grid[i][j] == 0:
                    goal_step += 1
        search(visited, init_path, worked_paths)
        return len(worked_paths)


if __name__ == '__main__':
    solution = Solution()
    grid = [[1, 0, 0, 0], [0, 0, 0, 0], [0, 0, 2, -1]]
    grid = [[1, 0, 0, 0], [0, 0, 0, 0], [0, 0, 0, 2]]
    grid = [[0, 1], [2, 0]]
    ans = solution.uniquePathsIII(grid)
    print(ans)