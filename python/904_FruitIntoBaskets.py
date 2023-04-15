class Solution(object):
    def totalFruit(self, fruits):
        dp = [0 for _ in range(len(fruits))]
        dp[0], picked = 1, [fruits[0]]
        index_in_basket = {fruits[0]: 0}
        for i, fruit in enumerate(fruits[1:]):
            if fruit in picked or len(picked) != 2:
                dp[i + 1] = dp[i] + 1
            else:
                if fruits[i] == picked[0]:
                    fruit_fall_out = picked.pop(1)
                else:
                    fruit_fall_out = picked.pop(0)
                dp[i + 1] = i + 1 - index_in_basket[fruit_fall_out]
            if fruit not in picked:
                picked.append(fruit)
            index_in_basket[fruit] = i + 1
        return max(dp)


if __name__ == '__main__':
    solution = Solution()
#    fruits = [1, 2, 1]
#    fruits = [0, 1, 2, 2]
#    fruits = [1, 2, 3, 2, 2]
    fruits = [1, 0, 1, 4, 1, 4, 1, 2, 3]
#    fruits = [1,0,2,3,4]
    ans = solution.totalFruit(fruits)
    print(ans)

