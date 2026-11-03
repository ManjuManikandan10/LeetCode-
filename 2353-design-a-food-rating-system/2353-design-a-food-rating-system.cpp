class FoodRatings {
private:
    //food->rating
    unordered_map<string,int>mpp;
    // food -> cuisine
    unordered_map<string,string> foodToCuisine;
    //comparator for the heap
    struct cmp{
        bool operator()(const pair<int,string>&a,const pair<int,string>&b){
            if(a.first != b.first) 
                return a.first < b.first;  // higher rating first
            return a.second > b.second;    // lex smaller name first
        }
    };
    //cuisine->priority_queue on the basis of ratings
    unordered_map<string, priority_queue<pair<int,string>, vector<pair<int,string>>, cmp>> pq;

public:
    FoodRatings(vector<string>& foods, vector<string>& cuisines, vector<int>& ratings) {
        int n=foods.size();
        /*  construct map containing the foods and their corresponding ratings
            and
            the priority queue
        */
        for(int i=0;i<n;i++){
            mpp[foods[i]]=ratings[i];
            foodToCuisine[foods[i]]=cuisines[i];
            pq[cuisines[i]].push({ratings[i], foods[i]});
        }
    }
    
    void changeRating(string food, int newRating) {
        mpp[food]=newRating;
        string cuisine=foodToCuisine[food];
        pq[cuisine].push({newRating, food});
    }
    
    string highestRated(string cuisine) {
        int rating=pq[cuisine].top().first;
        string food=pq[cuisine].top().second;
        
        while(mpp[food]!=rating){
            pq[cuisine].pop();
            rating=pq[cuisine].top().first;
            food=pq[cuisine].top().second;
        }
        return food;
    }
};

/**
 * Your FoodRatings object will be instantiated and called as such:
 * FoodRatings* obj = new FoodRatings(foods, cuisines, ratings);
 * obj->changeRating(food,newRating);
 * string param_2 = obj->highestRated(cuisine);
 */