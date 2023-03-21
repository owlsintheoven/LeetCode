class Solution(object):
    def arrayPairSum(self, nums):
        nums.sort(reverse=True)
        total = 0
        for i in range(1, len(nums), 2):
            total += nums[i]
        return total


if __name__ == '__main__':
    solution = Solution()
#    nums = [5, 1, 3, 4]
#    nums = [1, 4, 3, 2]
#    nums = [6, 2, 6, 5, 1, 2]
    nums = [1, 2, 3, 4, 5, 6]
#    nums = [1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1]
    ans = solution.arrayPairSum(nums)
    print(ans)