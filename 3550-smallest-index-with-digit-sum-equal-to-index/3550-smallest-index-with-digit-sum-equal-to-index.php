class Solution
{
	/**
	 * @param int[] $nums
	 * @return int
	 */
	function smallestIndex($nums)
	{
		foreach ($nums as $index => $num) {
			$ints = str_split($num);
			if ($index === array_sum($ints)) return $index;
		}
		return -1;
	}
}