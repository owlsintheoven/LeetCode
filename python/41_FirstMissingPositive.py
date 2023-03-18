class Solution(object):
    def firstMissingPositive(self, nums):
        nums.sort()
        first_positive = 1
        for num in nums:
            if num == first_positive:
                first_positive += 1
            elif num > first_positive:
                break
        return first_positive


if __name__ == '__main__':
    nums = [7,8,9,11,12]
    solution = Solution()
    ans = solution.firstMissingPositive(nums)
    print(ans)