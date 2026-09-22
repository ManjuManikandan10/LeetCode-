/**
 * // This is the interface that allows for creating nested lists.
 * // You should not implement it, or speculate about its implementation
 * class NestedInteger {
 *
 *     // if value is not specified, initializes an empty list.
 *     // Otherwise initializes a single integer equal to value.
 *     function __construct($value = null)
 *
 *     // Return true if this NestedInteger holds a single integer, rather than a nested list.
 *     function isInteger() : bool
 *
 *     // Return the single integer that this NestedInteger holds, if it holds a single integer
 *     // The result is undefined if this NestedInteger holds a nested list
 *     function getInteger()
 *
 *     // Set this NestedInteger to hold a single integer.
 *     function setInteger($i) : void
 *
 *     // Set this NestedInteger to hold a nested list and adds a nested integer to it.
 *     function add($ni) : void
 *
 *     // Return the nested list that this NestedInteger holds, if it holds a nested list
 *     // The result is undefined if this NestedInteger holds a single integer
 *     function getList() : array
 * }
 */

class NestedIterator {
    /**
     * @var array
     */
    private $flattenedList;
    /**
     * @var int
     */
    private $position;

    /**
     * @param NestedInteger[] $nestedList
     */
    function __construct($nestedList) {
        $this->flattenedList = [];
        $this->flattenNestedList($nestedList);
        $this->position = 0;
    }

    /**
     * @param NestedInteger[] $nestedList
     */
    private function flattenNestedList($nestedList) {
        foreach ($nestedList as $nestedInteger) {
            if ($nestedInteger->isInteger()) {
                $this->flattenedList[] = $nestedInteger->getInteger();
            } else {
                $this->flattenNestedList($nestedInteger->getList());
            }
        }
    }

    /**
     * @return Integer
     */
    function next() {
        return $this->flattenedList[$this->position++];
    }

    /**
     * @return Boolean
     */
    function hasNext() {
        return $this->position < count($this->flattenedList);
    }
}

/**
 * Your NestedIterator object will be instantiated and called as such:
 * $obj = NestedIterator($nestedList);
 * $ret_1 = $obj->next();
 * $ret_2 = $obj->hasNext();
 */