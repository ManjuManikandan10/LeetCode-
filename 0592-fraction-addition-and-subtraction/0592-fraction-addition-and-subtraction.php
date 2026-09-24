class Solution {

    /**
     * @param String $expression
     * @return String
     */
    function fractionAddition($expression) {
        $num = 0; 
        $denom = 1;

        $i = 0;

        while($i < strlen($expression))
        {
            $currentNum = 0;
            $currentDenom = 0;

            $isNegative = false;

            if($expression[$i] == '-' || $expression[$i] == '+')
            {
                if($expression[$i] == '-')
                {
                    $isNegative = true;
                }
                $i++;  
            }
            
            
            while(is_numeric($expression[$i]))
            {
                $val = $expression[$i] - '0';
                $currentNum = $currentNum * 10 + $val;
                $i++;
            }

            if($isNegative)
            {
                $currentNum *= -1;
            }

            $i++;
            
            while($i < strlen($expression) && is_numeric($expression[$i]))
            {
                $val = $expression[$i] - '0';
                $currentDenom = $currentDenom * 10 + $val;
                $i++;
            }

            $num = $num * $currentDenom + $currentNum * $denom;
            $denom = $denom * $currentDenom;
        }

        $greatestCommon = abs($this->gcd($num, $denom));

        $num /= $greatestCommon;
        $denom /= $greatestCommon;

        return $num . '/' . $denom;
    }

    function gcd($num1, $num2)
    {
        if($num1 == 0)
        {
            return $num2;
        }

        return $this->gcd($num2 % $num1, $num1);
    }
}