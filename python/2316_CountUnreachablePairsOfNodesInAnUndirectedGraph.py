class Solution(object):
    def countPairs(self, n, edges):
        adj = [[] for _ in range(n)]
        for edge in edges:
            adj[edge[0]].append(edge[1])
            adj[edge[1]].append(edge[0])
        count_in_group = []
        visited = [False for _ in range(n)]
        for i in range(n):
            if not visited[i]:
                count_in_group.append(0)
                self.dfs(adj, visited, i, count_in_group)
        for _ in range(n - sum(count_in_group)):
            count_in_group.append(1)
        to_be_paired = sum(count_in_group[1:])
        total = 0
        for i in range(len(count_in_group) - 1):
            total += count_in_group[i] * to_be_paired
            to_be_paired -= count_in_group[i + 1]
        return total

    def dfs(self, adj, vistied, src, counts):
        vistied[src] = True
        counts[-1] += 1
        for i in adj[src]:
            if not vistied[i]:
                self.dfs(adj, vistied, i, counts)


if __name__ == '__main__':
    solution = Solution()
    n = 3
    edges = [[0, 1], [0, 2], [1, 2]]
#    n = 7
#    edges = [[0, 2], [0, 5], [2, 4], [1, 6], [5, 4]]
    ans = solution.countPairs(n, edges)
    print(ans)