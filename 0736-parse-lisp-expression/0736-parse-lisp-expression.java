class Solution {
    public int evaluate(String expression) {
        Map<String, Stack<String>> m = new HashMap<>(); // it is store variable as key and a stack with its latest value on the top
        Stack<String> operation = new Stack<>(), operand = new Stack<>(); 
        StringBuilder sb = new StringBuilder();
        for (char ch : expression.toCharArray()) {
            if (ch == ' ') {    
                if(sb.length() > 0){    // we got some string formed
                    insert(sb.toString(), m, operand, operation); // insert the string into either operator/oprand stack
                    sb.setLength(0);    //reset set the string
                }
            } else if (ch == '(') {     // inorder mark as starting point of expression push into stack 
                operand.add("(");
                sb.setLength(0);
            } else if (ch == ')') {    // indicates we can resolve some part  
                if(sb.length() > 0){   // insert currently formed str before resolving
                    insert(sb.toString(), m, operand, operation);  
                    sb.setLength(0);
                }   
                resolve(m, operand, operation); // resolve the top operation
            } 
            else    sb.append(ch);    // just try to form a opertation/operand;
        }
        return Integer.parseInt(operand.pop());  // result is leftout element in operand;
    }

    private void insert(String str, Map<String, Stack<String>> m, Stack<String> operand, Stack<String> operation) {
        if (str.equals("let") || str.equals("mult") || str.equals("add"))    operation.add(str); // if the string is operation just add in operation stack
        else {
            if(operand.peek().charAt(0) > 96 && m.containsKey(str)) str = m.get(str).peek()+""; // check if previous operand was variable and if current one is also a variable, current one should be present in map and hence get the latest value 
            if (operation.peek().equals("let")) {   // if last operation is let then 
                if (str.charAt(0) < 97 && operand.peek().charAt(0) > 96) { // check current is a value and previous is variable
                    if (!m.containsKey(operand.peek()))    m.put(operand.peek(), new Stack<>()); // take the previous operand/variable and assign its latest value as current
                    m.get(operand.peek()).add(str);
                } 
                operand.add(str);   // add the current to operand
            } else {    // if the last opertaion is mult/add then 
                if (m.containsKey(str))    operand.add(m.get(str).peek() + "");  // check if current is variable if yes then add its latest value to operand
                else    operand.add(str);   // if current itself is the value
            }
        }
    }

    private void resolve(Map<String, Stack<String>> m, Stack<String> operand, Stack<String> operation) {
        String currOperation = operation.pop(); // get the lastest operation to perform
        String res = ""; // stores the result of current operation
        if (currOperation.equals("let")) {  
            res = operand.pop();     // first eleme will be result of expression since we will always have (let v1, e1, v2, e2...exp) as per question
            if(m.containsKey(res))  res = m.get(res).peek()+""; // check if the result of expression is variable if yes replace it will variables latest value
            while (!(operand.peek()).equals("(")) { // do untile we reach starting point of the operation
                String key = operand.pop();     // remove the operand
                if(m.containsKey(key))    m.get(key).pop(); // if the operand was variable remove the latest value of that variable
            }
        } else {
            int num1 = Integer.parseInt(operand.pop()), num2 = Integer.parseInt(operand.pop());  // if any operation other than let then take top 2 operand which will be numbers for any valid input
            if (currOperation.equals("add"))    num1 += num2;   // perform the required operation and store it in num1
            else    num1 *= num2;
            res = num1 + ""; // num1 is the result
        }
        operand.pop(); // remove start point i.e., '('
        if(!operation.isEmpty() && operation.peek().equals("let") && operand.peek().charAt(0) > 96){ // if the previous operation is let and previous operand is variable
            if(!m.containsKey(operand.peek()))    m.put(operand.peek(), new Stack<>());
            m.get(operand.peek()).add(res); // assign the variable to current result 
        }
        operand.add(res); // add the result to operand
    }
}