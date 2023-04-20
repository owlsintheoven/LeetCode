class Solution(object):
    def arrayRankTransform(self, arr):
        rank = {}
        for a in sorted(arr):
            rank.setdefault(a, len(rank) + 1)
        return list(map(rank.get, arr))

#        if len(arr) == 0:
#            return []
#        temp = arr.copy()
#        temp.sort()
#        rank_dict = {}
#        rank = 1
#        prev = temp[0] - 1
#        for a in temp:
#            if prev != a:
#                rank_dict[a] = rank
#                rank += 1
#            prev = a
#        return [rank_dict[a] for a in arr]
#        a_max, a_min = max(arr), min(arr)
#        buckets = [False for _ in range(a_max - a_min + 1)]
#        for a in arr:
#            buckets[a - a_min] = True
#        rank_dict = {}
#        rank = 1
#        for i in range(len(buckets)):
#            if buckets[i]:
#                rank_dict[i + a_min] = rank
#                rank += 1
#        return [rank_dict[a] for a in arr]


if __name__ == '__main__':
    solution = Solution()
#    ans = solution.arrayRankTransform([40, 10, 20, 30])
#    ans = solution.arrayRankTransform([37,12,28,9,100,56,80,5,12])
#    ans = solution.arrayRankTransform([100, 100, 100])
    ans = solution.arrayRankTransform([])
    print(ans)

