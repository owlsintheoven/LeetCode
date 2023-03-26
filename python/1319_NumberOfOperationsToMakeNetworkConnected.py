class Solution(object):
    def makeConnected(self, n, connections):
        if len(connections) < n - 1:
            return -1
        adj = [[] for _ in range(n)]
        for connection in connections:
            adj[connection[0]].append(connection[1])
            adj[connection[1]].append(connection[0])
        print(adj)
        visited = [False for _ in range(n)]
        not_visited = 0
        for i in range(n):
            if not visited[i]:
                self.dfs(adj, visited, i)
                not_visited += 1
        return not_visited - 1

    def dfs(self, adj, visited, src):
        visited[src] = True
        for i in adj[src]:
            if not visited[i]:
                self.dfs(adj, visited, i)


#   def makeConnected(self, n, connections):
#       if len(connections) < n - 1:
#           return -1
#       connected_groups = []
#       for connection in connections:
#           exist_in_group = []
#           for i in range(2):
#               node = connection[i]
#               for j in range(len(connected_groups)):
#                   if node in connected_groups[j]:
#                       exist_in_group.append(j)
#                       break
#               if len(exist_in_group) != i + 1:
#                   exist_in_group.append(-1)
#           if exist_in_group[0] == -1 and exist_in_group[1] == -1:
#               connected_groups.append(connection)
#           elif exist_in_group[0] == -1:
#               connected_groups[exist_in_group[1]].append(connection[0])
#           elif exist_in_group[1] == -1:
#               connected_groups[exist_in_group[0]].append(connection[1])
#           elif exist_in_group[0] != exist_in_group[1]:
#               exist_in_group.sort()
#               group = connected_groups.pop(exist_in_group[1])
#               connected_groups[exist_in_group[0]].extend(group)
#           print(connected_groups)
#       connected_total = sum([len(group) for group in connected_groups])
#       rearrangement = len(connected_groups) - 1 + (n - connected_total)
#
#       return rearrangement


if __name__ == '__main__':
    solution = Solution()
#    n = 4
#    connections = [[0, 1], [0, 2], [1, 2]]
    n = 6
    connections = [[0, 1], [0, 2], [0, 3], [1, 2], [1, 3]]
#    n = 100
#    connections = [[17,51],[33,83],[53,62],[25,34],[35,90],[29,41],[14,53],[40,84],[41,64],[13,68],[44,85],[57,58],[50,74],[20,69],[15,62],[25,88],[4,56],[37,39],[30,62],[69,79],[33,85],[24,83],[35,77],[2,73],[6,28],[46,98],[11,82],[29,72],[67,71],[12,49],[42,56],[56,65],[40,70],[24,64],[29,51],[20,27],[45,88],[58,92],[60,99],[33,46],[19,69],[33,89],[54,82],[16,50],[35,73],[19,45],[19,72],[1,79],[27,80],[22,41],[52,61],[50,85],[27,45],[4,84],[11,96],[0,99],[29,94],[9,19],[66,99],[20,39],[16,85],[12,27],[16,67],[61,80],[67,83],[16,17],[24,27],[16,25],[41,79],[51,95],[46,47],[27,51],[31,44],[0,69],[61,63],[33,95],[17,88],[70,87],[40,42],[21,42],[67,77],[33,65],[3,25],[39,83],[34,40],[15,79],[30,90],[58,95],[45,56],[37,48],[24,91],[31,93],[83,90],[17,86],[61,65],[15,48],[34,56],[12,26],[39,98],[1,48],[21,76],[72,96],[30,69],[46,80],[6,29],[29,81],[22,77],[85,90],[79,83],[6,26],[33,57],[3,65],[63,84],[77,94],[26,90],[64,77],[0,3],[27,97],[66,89],[18,77],[27,43]]
    ans = solution.makeConnected(n, connections)
    print(ans)