type TweetCounts struct {
    Tweet map[string][]int // key: name val: []int times
}

func Constructor() TweetCounts {
    return TweetCounts{
        Tweet: map[string][]int{},
    }
}


func (this *TweetCounts) RecordTweet(tweetName string, time int)  {
    if _, ok := this.Tweet[tweetName]; !ok {
        this.Tweet[tweetName] = []int{time}
    } else {
        this.Tweet[tweetName] = append(this.Tweet[tweetName], time)
    }
}


func (this *TweetCounts) GetTweetCountsPerFrequency(freq string, tweetName string, startTime int, endTime int) []int {
    interval := GetInterval(freq)
	
    if _, ok := this.Tweet[tweetName]; !ok {
        return []int{}
    }
    
    result := make([]int, (endTime - startTime + interval) / interval)
    for _, time := range this.Tweet[tweetName] {
        if time > endTime || time < startTime {
            continue
        }
        
        thisInterval := (time-startTime)/interval
        result[thisInterval]++
    }
    
    return result
}


func GetInterval(freq string) int {
    if freq == "minute" {
        return 60
    }
    
    if freq == "hour" {
        return 3600
    }
    
    if freq == "day" {
        return 86400
    }
    
    return 0 // never reach here
}

/**
 * Your TweetCounts object will be instantiated and called as such:
 * obj := Constructor();
 * obj.RecordTweet(tweetName,time);
 * param_2 := obj.GetTweetCountsPerFrequency(freq,tweetName,startTime,endTime);
 */