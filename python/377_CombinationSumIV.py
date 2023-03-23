class Solution(object):
    perm_dict = {}
    def combinationSum4(self, nums, target):
        nums.sort()
        all = []
        self.helper(nums, target, [], all)
        for a in all:
            print(a)
        count = 0
#        for arr in all:
#            print(arr)
#            takes = [1]
#            prev = arr[0]
#            for a in arr[1:]:
#                if a != prev:
#                    takes.append(1)
#                    prev = a
#                else:
#                    takes[-1] += 1
#            count += self.possiblities(len(arr), takes)

        return count

    def helper(self, nums, target, current, all):
        if target == 0:
            all.append(current.copy())
            return True
        elif target < 0:
            return False
        else:
            for i in range(len(nums)):
                current.append(nums[i])
                found = self.helper(nums[i:], target - nums[i], current, all)
                current.pop()
                if not found:
                    continue

    def permutation(self, n):
        if n in self.perm_dict:
            return self.perm_dict[n]
        if n == 0 or n == 1:
            return 1
        else:
            perm = n * self.permutation(n - 1)
            self.perm_dict[n] = perm
            return perm

    def combination(self, n, r):
        if r == 1:
            return n
        elif n == r:
            return 1
        return int(self.permutation(n) / (self.permutation(r) * self.permutation(n - r)))

    def possiblities(self, space, takes):
        # space: 4, takes: [3, 1]
        total = 1
        for take in takes:
            total *= self.combination(space, take)
            space -= take
        return total



if __name__ == '__main__':
    solution = Solution()
    nums = [1, 2, 3]
    target = 6
#    nums = [9]
#    target = 3
#    nums = [4, 2, 1]
#    target = 32
#    target = 4
#    target = 8 #[2, 1, 4, 1]
    nums = [10, 20, 30, 40, 50, 60, 70, 80, 90, 100, 110, 120, 130, 140, 150, 160, 170, 180, 190, 200, 210, 220, 230, 240, 250,
     260, 270, 280, 290, 300, 310, 320, 330, 340, 350, 360, 370, 380, 390, 400, 410, 420, 430, 440, 450, 460, 470, 480,
     490, 500, 510, 520, 530, 540, 550, 560, 570, 580, 590, 600, 610, 620, 630, 640, 650, 660, 670, 680, 690, 700, 710,
     720, 730, 740, 750, 760, 770, 780, 790, 800, 810, 820, 830, 840, 850, 860, 870, 880, 890, 900, 910, 920, 930, 940,
     950, 960, 970, 980, 990, 111]
    target = 999
    ans = solution.combinationSum4(nums, target)
    print(ans)