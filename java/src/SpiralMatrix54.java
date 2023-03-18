import java.util.List;
import java.util.ArrayList;

class Solution {
    enum Direction {
        Right,
        Down,
        Left,
        Up
    }

    public List<Integer> spiralOrder(int[][] matrix) {
        int r = matrix.length;
        int c = matrix[0].length;

        List<Integer> spiral = new ArrayList();

        int targetSize = c * r;
        Direction direction = Direction.Right;

        int currentRowIndex = 0;
        int currentColIndex = 0;
        int sizeLeftInRow = r - 1;
        int sizeLeftInCol = c;

        while (spiral.size() != targetSize) {
            switch (direction) {
                case Right:
                    for (int i = 0; i < sizeLeftInCol; i++) {
                        spiral.add(matrix[currentRowIndex][currentColIndex + i]);
                    }
                    currentColIndex += sizeLeftInCol - 1;
                    sizeLeftInCol--;
                    currentRowIndex++;
                    direction = Direction.Down;
                    break;
                case Down:
                    for (int i = 0; i < sizeLeftInRow; i++) {
                        spiral.add(matrix[currentRowIndex + i][currentColIndex]);
                    }
                    currentRowIndex += sizeLeftInRow - 1;
                    sizeLeftInRow--;
                    currentColIndex--;
                    direction = Direction.Left;
                    break;
                case Left:
                    for (int i = 0; i < sizeLeftInCol; i++) {
                        spiral.add(matrix[currentRowIndex][currentColIndex - i]);
                    }
                    currentColIndex -= sizeLeftInCol - 1;
                    sizeLeftInCol--;
                    currentRowIndex--;
                    direction = Direction.Up;
                    break;
                case Up:
                    for (int i = 0; i < sizeLeftInRow; i++) {
                        spiral.add(matrix[currentRowIndex - i][currentColIndex]);
                    }
                    currentRowIndex -= sizeLeftInRow - 1;
                    sizeLeftInRow--;
                    currentColIndex++;
                    direction = Direction.Right;
                    break;
            }
        }

        return spiral;
    }
}
//      c,r,c,r,c
// 3,3: 3,2,2,1,1
// 3,4: 4,2,3,1,2

// 5,1: 1,4
// 4,3: 3,3,2,2,1,1

// 1,1 -> 0,0
// 1,2 -> 0,1
// 2,1 -> 1,0 *
// 2,2 -> 1,0 *
// 2,3 -> 1,0 *
// 3,3 -> 1,1 /
// 3,4 -> 1,2 /
// 3,5 -> 1,3 /
