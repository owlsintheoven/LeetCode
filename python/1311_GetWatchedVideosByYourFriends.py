class Solution(object):
    def watchedVideosByFriends(self, watchedVideos, friends, id, level):
        level_friends = set(f for f in friends[id])
        visisted = [True if i in level_friends else False for i in range(len(friends))]
        for _ in range(1, level):
            temp = set()
            while len(level_friends) != 0:
                for f in friends[level_friends.pop()]:
                    if not visisted[f]:
                        temp.add(f)
                        visisted[f] = True
            level_friends = temp
        if id in level_friends:
            level_friends.remove(id)
        watched = {}
        while len(level_friends) != 0:
            for v in watchedVideos[level_friends.pop()]:
                watched[v] = watched[v] + 1 if v in watched else 1
        return [v[0] for v in sorted(watched.items(), key=lambda a: (a[1], a[0]))]


if __name__ == '__main__':
    solution = Solution()
    watchedVideos = [["A", "B"], ["C"], ["B", "C"], ["D"]]
    friends = [[1, 2], [0, 3], [0, 3], [1, 2]]
    id = 0
    level = 1
    ans = solution.watchedVideosByFriends(watchedVideos, friends, id, level)
    print(ans)

