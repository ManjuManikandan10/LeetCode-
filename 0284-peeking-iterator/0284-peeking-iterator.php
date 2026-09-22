class PeekingIterator {
    private int $iterator = 0;
    function __construct($arr) {
        $this->arr = $arr;
    }
    function next() {
        $this->iterator += 1;
        return $this->arr[$this->iterator - 1];
    }
    function peek() {
        return $this->arr[$this->iterator];
    }
    function hasNext() {
        $bool =  $this->iterator === count($this->arr);
        return !$bool;
    }
}