class Solution {

    /**
     * @param Integer[] $nums
     * @param Integer $k
     * @return Integer[]
     */
    function maxSlidingWindow($nums, $k) {
        if($k==1){
            return $nums;
        }

        $result=array();
        $count=count($nums);
        if($k>=$count){
            $result[]=max($nums);
            return $result;
        }

        $queue=new SplQueue();

        for($i=0;$i<$k;$i++){
            while(!$queue->isEmpty() && $nums[$i] > $nums[$queue->top()]){
                $queue->pop();
            }
            $queue->push($i);
        }

        for(;$i<$count;$i++){
            $result[]=$nums[$queue->bottom()];

            if($queue->bottom() <= $i-$k){
                $queue->shift();
            }

            while(!$queue->isEmpty() && $nums[$i] > $nums[$queue->top()]){
                $queue->pop();
            }
            $queue->push($i);

        }
        $result[]=$nums[$queue->bottom()];
        return $result;

    }
}