class Solution(object):
    def maximumGap(self, nums):
        if len(nums) < 2:
            return 0
        # pigeonhole sort
        min = nums[0]
        max = nums[0]
        for num in nums:
            if num < min:
                min = num
            elif num > max:
                max = num
        pigeonhole = [[] for i in range(max - min + 1)]
        for num in nums:
            pigeonhole[num - min].append(num)
        pointer = 0
        for i in range(max - min + 1):
            while len(pigeonhole[i]) != 0:
                num = pigeonhole[i].pop(0)
                nums[pointer] = num
                pointer += 1
        max_gap = 0
        for i in range(1, len(nums)):
            gap = nums[i] - nums[i-1]
            max_gap = gap if gap > max_gap else max_gap
        return max_gap


if __name__ == '__main__':
    solution = Solution()
    nums = [3, 6, 9, 1]
    max_gap = solution.maximumGap(nums)
    print(max_gap)