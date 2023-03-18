class Solution(object):
    def threeSum(self, nums):
        neg_nums = []
        pos_nums = []
        for num in nums:
            if num < 0:
                neg_nums.append(num)
            elif num > 0:
                pos_nums.append(num)
        neg_nums.sort()
        pos_nums.sort()
        return nums

if __name__ == '__main__':
    solution = Solution()
    nums = [-1, 0, 1, 2, -1, -4]
    ans = solution.threeSum(nums)
    print(ans)