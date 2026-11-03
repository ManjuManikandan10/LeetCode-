class AuctionSystem {
public:
map<pair<int,int>,int>m;
map<int,set<pair<int,int>>>m1;

    AuctionSystem() {
        
    }
    
    void addBid(int userId, int itemId, int bidAmount) {
        if(m.find({userId,itemId})!=m.end())
        {
            int amount=m[{userId,itemId}];
            m[{userId,itemId}]=bidAmount;
            m1[itemId].erase({amount,userId});
            m1[itemId].insert({bidAmount,userId});
        }
        m[{userId,itemId}]=bidAmount;
        if(m1.find(itemId)==m1.end())
        {
            set<pair<int,int>>st;
            st.insert({bidAmount,userId});
            m1[itemId]=st;
        }
        else{
            m1[itemId].insert({bidAmount,userId});
        }
    }
    
    void updateBid(int userId, int itemId, int newAmount) {
        int amount=m[{userId,itemId}];
        m.erase({userId,itemId});
        m[{userId,itemId}]=newAmount;
        m1[itemId].erase({amount,userId});
        m1[itemId].insert({newAmount,userId});

    }
    
    void removeBid(int userId, int itemId) {
         int amount=m[{userId,itemId}];
        m.erase({userId,itemId});
         m1[itemId].erase({amount,userId});
    }
    
    int getHighestBidder(int itemId) {
        if(m1.find(itemId)==m1.end())
        {
            return -1;
        }
        if(m1[itemId].size()==0)
        {
            return -1;
        }
        auto it =prev(m1[itemId].end());
        return it->second;
    }
};

/**
 * Your AuctionSystem object will be instantiated and called as such:
 * AuctionSystem* obj = new AuctionSystem();
 * obj->addBid(userId,itemId,bidAmount);
 * obj->updateBid(userId,itemId,newAmount);
 * obj->removeBid(userId,itemId);
 * int param_4 = obj->getHighestBidder(itemId);
 */