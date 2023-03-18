import java.util.Arrays;
import java.util.List;
import java.util.ArrayList;

// 1. sort candidates
// 2. iterate through candidates
// 3. use stack to store candidates from smallest to largest
// 4. keep deduct target with the top of the stack, and push the top to a queue
// 5. if deduct to 0, store this combinations
// 6. if not, delete the tail of the queue and pop up the stack, back to 4 and loop

// change target
// change candidates
public class CombinationSum39 {
    public List<List<Integer>> combinationSum(int[] candidates, int target) {
        ArrayList<Integer> answer = new ArrayList();

        return null;
    }

    // change target

    // change candidatesae-stg-va-common-psql.postgres.database.usgovcloudapi.nets

    public List<List<Integer>> iterateCandidates(List<Integer> candidates, int target) {
        ArrayList<Integer> answer = new ArrayList<>();
        Integer pos = 0;
        Integer length = candidates.size();
        while (pos != length) {
            List<Integer> combination = helper(candidates.subList(pos, length), target);
            if (combination != null) {

            }
        }
    }

    public List<Integer> helper(List<Integer> candidates, int target) {
        if (target == 0) {
            return new ArrayList<Integer>(Arrays.asList(candidates.get(0)));
        }
        else if (target < 0) {
            return null;
        }
        else {
            Integer pos = 0;
            while (pos != candidates.size()) {
                List<Integer> combination = helper(candidates, target - candidates.get(pos));
                if combination != null {

                }
            }
        }
        return null;
    }
}
