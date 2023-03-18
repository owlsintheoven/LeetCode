import time
class Solution(object):
    def combinationSum2(self, candidates, target):
        candidates.sort()
        candidates_transformed = []
        candidate_count_dict = {}
        for candidate in candidates:
            if candidate not in candidates_transformed:
                candidates_transformed.append(candidate)
                candidate_count_dict[candidate] = 1
            else:
                candidate_count_dict[candidate] += 1

        answer_transformed = self.helper(candidates_transformed, target)
        answer = []
        for combination in answer_transformed:
            valid = True
            for candidate in candidates_transformed:
                if candidate_count_dict[candidate] < combination.count(candidate):
                    valid = False
                    break
            if valid:
                answer.append(combination)

        return answer

    def helper(self, candidates, target):
        if target == 0:
            return [[]]
        elif target < 0:
            return None
        else:
            answer = []
            for i in range(len(candidates)):
                combinations = self.helper(candidates[i:], target - candidates[i])
                if combinations != None:
                    for combination in combinations:
                        new_combination = [candidates[i]]
                        new_combination.extend(combination)
                        answer.append(new_combination)
            return answer

#    def helper(self, candidates, target):
#        if target == 0:
#            return [[]]
#        elif target < 0:
#            return None
#        else:
#            answer = []
#            for i in range(len(candidates)):
#                combinations = self.helper(candidates[i+1:], target - candidates[i])
#                if combinations != None:
#                    for combination in combinations:
#                        new_combination = [candidates[i]]
#                        new_combination.extend(combination)
#                        if new_combination not in answer:
#                            answer.append(new_combination)
#                    # if found answer, skip next num with same value
#                    #while True:
#                    #    if i < len(candidates) - 1 and candidates[i] == candidates[i+1]:
#                    #        i += 1
#                    #    else:
#                    #        break
#                    #if answer == [] and i == len(candidates) - 1:
#                    #    break
#                    #elif answer != []:
#                    #    while i < len(candidates):
#                    #        if candidates[0] == candidates[i]:
#                    #            i += 1
#                    #        else:
#                    #            break
#            return answer


if __name__ == '__main__':
    solution = Solution()
    start = time.time()
    #answer = solution.combinationSum2([2,5,2,1,2], 5) # 1, 2, 2, 2, 5
    #answer = solution.combinationSum2([10,1,2,7,6,1,5], 8)
    #answer = solution.combinationSum2([1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1], 27)
    #answer = solution.combinationSum2([1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1], 27)
    #answer = solution.combinationSum2([1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1], 27)
    answer = solution.combinationSum2([1,2,5], 5)
    end = time.time()
    #answer = solution.combinationSum2([1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1], 30)
    print()
    print(answer)
    print(end-start)


'''javascript
/**
 * @param {number[]} candidates
 * @param {number} target
 * @return {number[][]}
 */
var combinationSum2 = function(c, target) {
    c.sort((a,b)=>a-b)
    let res = []

    let iterate = (index,sum,temp) =>{
        if(sum>target) return;
        if(sum == target){
            res.push(temp)
            return;
        }
        // 1 1 2 5 6 7 10 
        for(let i =index; i<c.length;i++){
            if(i != index && c[i] == c[i-1]) continue;
            iterate(i+1,sum+c[i],[...temp,c[i]])
        }
    }
    iterate(0,0,[])
    return res;
};
'''