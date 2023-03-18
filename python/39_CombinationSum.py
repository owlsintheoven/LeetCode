class Solution:
    def combinationSum(self, candidates, target):
        if target == 0:
            return [[]]
        elif target < 0:
            return None
        else:
            answer = []
            for i in range(len(candidates)):
                combinations = self.combinationSum(candidates[i:], target - candidates[i])
                if combinations != None:
                    for combination in combinations:
                        new_combination = [candidates[i]]
                        new_combination.extend(combination)
                        answer.append(new_combination)
            return answer


if __name__ == '__main__':
    solution = Solution();
    answer = solution.combinationSum([1,2,5], 5)
    print(answer)

