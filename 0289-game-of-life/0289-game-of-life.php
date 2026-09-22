class Solution {

    /**
     * @param Integer[][] $board
     * @return NULL
     */
    function gameOfLife(&$board) {
        $rows = count($board);
        $cols = count($board[0]);

        // Process each cell
        for ($r = 0; $r < $rows; $r++) {
            for ($c = 0; $c < $cols; $c++) {
                $liveNeighbors = $this->countLiveNeighbors($board, $r, $c, $rows, $cols);

                // Apply the rules
                if ($board[$r][$c] == 1 && ($liveNeighbors < 2 || $liveNeighbors > 3)) {
                    $board[$r][$c] = -1; // Live -> Dead
                } elseif ($board[$r][$c] == 0 && $liveNeighbors == 3) {
                    $board[$r][$c] = 2; // Dead -> Live
                }
            }
        }

        // Update the board to the new state
        for ($r = 0; $r < $rows; $r++) {
            for ($c = 0; $c < $cols; $c++) {
                if ($board[$r][$c] == -1) {
                    $board[$r][$c] = 0; // Dead
                } elseif ($board[$r][$c] == 2) {
                    $board[$r][$c] = 1; // Live
                }
            }
        }
    }

    // to count live neighbors for a cell
    private function countLiveNeighbors($board, $r, $c, $rows, $cols) {
        $directions = [
            [-1, -1], [-1, 0], [-1, 1],
            [0, -1],         [0, 1],
            [1, -1], [1, 0], [1, 1]
        ];

        $liveNeighbors = 0;

        foreach ($directions as $direction) {
            $newRow = $r + $direction[0];
            $newCol = $c + $direction[1];

            if ($newRow >= 0 && $newRow < $rows && $newCol >= 0 && $newCol < $cols && abs($board[$newRow][$newCol]) == 1) {
                $liveNeighbors++;
            }
        }

        return $liveNeighbors;
    }
}