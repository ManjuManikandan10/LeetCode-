var isReachable = function(targetX, targetY) {
    while(true){
        // first make both odd
        while(targetX % 2 == 0){
            targetX /= 2;
        }
        while(targetY % 2 == 0){
            targetY /= 2;
        }
        
        if(targetX == 1 || targetY == 1){
            return true;
        }

        // if they same and not anyone is 1 then obvius that it would never become any of them 1
        if(targetX == targetY){
            return false;
        }
        
        // two odd addition is always even
        // so make even(using addition of both)who is greater 
        if(targetX > targetY){
            targetX += targetY;
        }
        else{
            targetY += targetX;
        }
    }
    
    return true;
};