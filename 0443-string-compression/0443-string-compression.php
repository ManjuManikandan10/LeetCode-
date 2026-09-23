class Solution {

    /**
     * @param String[] $chars
     * @return Integer
     */
    function compress(&$chars) {
        $n = count($chars);
        if ($n == 1) return 1;
    
        $write = 0; 
        $read = 0;
    
        while ($read < $n) {
            $char = $chars[$read];
            $count = 0;
    
            while ($read < $n && $chars[$read] == $char) {
                $read++;
                $count++;
            }
    
            $chars[$write++] = $char;
    
            if ($count > 1) {
                foreach (str_split((string)$count) as $digit) {
                    $chars[$write++] = $digit;
                }
            }
        }
    
        return $write;          
    }
}