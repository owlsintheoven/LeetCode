import math

class Solution(object):
    def mincostTickets(self, days, costs):
        max_day = days[-1]
        dp = [0 for _ in range(max_day + 1)]
        is_travel_day = [False for _ in range(max_day + 1)]
        for day in days:
            is_travel_day[day] = True
        for i in range(max_day + 1):
            if is_travel_day[i] == False:
                dp[i] = dp[i - 1]
            else:
                cost1 = dp[i - 1] + costs[0]
                cost7 = dp[max(0, i - 7)] + costs[1]
                cost30 = dp[max(0, i - 30)] + costs[2]
                dp[i] = min(cost1, cost7, cost30)
        return dp[max_day]
#    def mincostTickets(self, days, costs):
#        a = min(costs)
#        b = min(7 * a, min(costs[1:]))
#        c = min(30 * a, 5 * b, costs[2])
#        costs = [c, b, a]
##        group_if_exceed = [-1, min(math.floor(c / b) * 7, math.floor(c / a)), math.floor(b / a)]
##        print(group_if_exceed)
#        stacks = [[] for _ in range(3)]
#        print(days[0])
#        stacks[-1].append([days[0], 1, days[0]])
#        def check_all_stacks():
#            # check if we can add item into 7-day stack
#            # need to check through 1-day stack
#            if len(stacks[2]) != 0:
#                latest = stacks[2][-1][2]
#                idx = len(stacks[2]) - 1
#                while idx != -1 and latest - stacks[2][idx][0] < 7:
#                    idx -= 1
#                if (len(stacks[2]) - idx - 1) * costs[2] > costs[1]:
#                    stacks[1].append([stacks[2][idx + 1][0], len(stacks[2]) - idx - 1, latest])
#                    del stacks[2][idx + 1:]
#            # check if we can add item into 30-day stack
#            # find the lastest from both 1-day bucket and 7-day bucket
#            if len(stacks[1]) != 0:
#                latest = stacks[1][-1][2]
#                if len(stacks[2]) != 0:
#                    latest = max(latest, stacks[2][-1][2])
#                idx_1 = len(stacks[1]) - 1
#                while idx_1 != -1 and latest - stacks[1][idx_1][0] < 30:
#                    idx_1 -= 1
#                idx_2 = len(stacks[2]) - 1
#                while idx_2 != -1 and latest - stacks[2][idx_2][0] < 30:
#                    idx_2 -= 1
#                if (len(stacks[1]) - idx_1 - 1) * costs[1] + (len(stacks[2]) - idx_2 - 1) * costs[2] > costs[0]:
##                if sum([stacks[1][i][1] for i in range(len(stacks[1]) - 1, idx_1, -1)]) + sum(
##                        [stacks[2][i][1] for i in range(len(stacks[2]) - 1, idx_2, -1)]) > group_if_exceed[1]:
#                    stacks[0].append(
#                        [min(stacks[1][idx_1 + 1][0], stacks[2][idx_2 + 1][0]), sum([stacks[1][i][1] for i in range(len(stacks[1]) - 1, idx_1, -1)]) + sum([stacks[2][i][1] for i in range(len(stacks[2]) - 1, idx_2, -1)]), latest])
#                    del stacks[1][idx_1 + 1:]
#                    del stacks[2][idx_2 + 1:]
#
#        for day in days[1:]:
#            print(day)
#            # iterature through 30-day bucket, 7-day bucket and 1-day bucket
#            if len(stacks[0]) != 0 and day - stacks[0][-1][0] < 30:
#                stacks[0][-1][1] += 1
#                stacks[0][-1][2] = day
#                continue
#            if len(stacks[1]) != 0 and day - stacks[1][-1][0] < 7:
#                stacks[1][-1][1] += 1
#                stacks[1][-1][2] = day
#                check_all_stacks()
#                continue
#            stacks[2].append([day, 1, day])
#            check_all_stacks()
#            for stack in stacks:
#                print(stack)
#        for stack in stacks:
#            print(stack)
#        print()
#        total = 0
#        for i in range(3):
#            total += costs[i] * len(stacks[i])
#        return total



if __name__ == '__main__':
    solution = Solution()
    days = [1, 4, 6, 7, 8, 20]
    costs = [2, 7, 15]
    days = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 30, 31]
#    days = [1, 2, 3, 4, 5, 6, 7, 8]
#    costs = [2, 7, 15]
    days = [1,4,6,9,10,11,12,13,14,15,16,17,18,20,21,22,23,27,28]
    costs = [3,13,45]
    ans = solution.mincostTickets(days, costs)
    print(ans)

