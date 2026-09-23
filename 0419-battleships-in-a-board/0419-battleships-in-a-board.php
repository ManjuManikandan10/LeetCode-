class Solution {

    /**
     * @param String[][] $board
     * @return Integer
     */
    function countBattleships($board) {
        $m = count($board);
        $n = count($board[0]);
        $count = 0;
        
        for ($i = 0; $i < $m; $i++) {
            for ($j = 0; $j < $n; $j++) {
                if ($board[$i][$j] == 'X' && ($i == 0 || $board[$i - 1][$j] == '.') && ($j == 0 || $board[$i][$j - 1] == '.')) {
                    $count++;
                }
            }
        }
        
        return $count;
    }
}

$solution = new Solution();
$board1 = [["X",".",".","X"],[".",".",".","X"],[".",".",".","X"]];
$board2 = [["."]];
echo $solution->countBattleships($board1) . "\n"; // Output: 2
echo $solution->countBattleships($board2) . "\n"; // Output: 0