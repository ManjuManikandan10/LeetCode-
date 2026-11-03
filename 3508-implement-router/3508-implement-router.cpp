class Node
{
    public:
    int source,dest,time;
    Node* next;
    Node* prev;
    Node(int s,int d,int t)
    {
        source=s;
        dest=d;
        time=t;
        next=NULL;
        prev=NULL;
    }
};
struct TupleHash
{
    size_t operator()(const tuple<int,int,int>& t) const
    {
        auto [a,b,c]=t;
        return hash<int>()(a)^(hash<int>()(b)<<1)^(hash<int>()(c)<<2);
    }
};
class Router {
public:
    //LRU to evict oldest packet
    //map to see if a node with same attributes exist already
    //forwarding is deleting at tail
    //need a mapping for dest->time too 
    int sz,maxi;
    unordered_map<tuple<int,int,int>,Node*,TupleHash> mapping;
    unordered_map<int,deque<int>> destTime;
    Node* dummyhead,*dummytail;
    void addAfterHead(tuple<int,int,int> data)
    {
        Node* newNode=new Node(get<0>(data),get<1>(data),get<2>(data));
        Node* afterHead=dummyhead->next;
        dummyhead->next=newNode;
        newNode->prev=dummyhead;
        newNode->next=afterHead;
        afterHead->prev=newNode;
        mapping[data]=newNode;
        destTime[newNode->dest].push_back(newNode->time);
        sz++;
    }
    void deleteNode(Node* node)
    {
        Node* beforeNode=node->prev,*afterNode=node->next;
        beforeNode->next=afterNode;
        afterNode->prev=beforeNode;
        destTime[node->dest].pop_front();
        sz--;
        delete node;
    }
    Router(int memoryLimit) {
        dummyhead=new Node(0,0,0);
        dummytail=new Node(0,0,0);
        sz=0;
        dummyhead->next=dummytail;
        dummytail->prev=dummyhead;
        maxi=memoryLimit;
    }
    
    bool addPacket(int src, int dest, int time) {
        tuple<int,int,int> t=make_tuple(src,dest,time);
        if(mapping.find(t)!=mapping.end()) return false;
        if(sz==maxi)
        {
            Node* toDelete=dummytail->prev;
            tuple<int,int,int> toDel=make_tuple(toDelete->source,toDelete->dest,toDelete->time);
            mapping.erase(toDel);
            deleteNode(toDelete);
        }
        addAfterHead(t);
        return true;
    }
    
    vector<int> forwardPacket() {
        vector<int> ans;
        if(sz==0) return ans;
        Node* node=dummytail->prev;
        ans.push_back(node->source);
        ans.push_back(node->dest);
        ans.push_back(node->time);
        mapping.erase({node->source,node->dest,node->time});
        deleteNode(node);
        return ans;
    }
    
    int getCount(int destination, int startTime, int endTime) {
        auto &dq=destTime[destination];
        auto l=lower_bound(dq.begin(),dq.end(),startTime);
        auto r=upper_bound(dq.begin(),dq.end(),endTime);
        return r-l;
    }
};

/**
 * Your Router object will be instantiated and called as such:
 * Router* obj = new Router(memoryLimit);
 * bool param_1 = obj->addPacket(source,destination,timestamp);
 * vector<int> param_2 = obj->forwardPacket();
 * int param_3 = obj->getCount(destination,startTime,endTime);
 */