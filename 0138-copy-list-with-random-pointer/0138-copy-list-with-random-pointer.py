class Solution:
    def copyRandomList(self, head):
        if head is None:
            return None

        # Map original nodes to copied nodes
        old_to_new = {}

        # Create a new node for every original node
        current = head
        while current:
            old_to_new[current] = Node(current.val)
            current = current.next

        # Connect next and random pointers
        current = head
        while current:
            old_to_new[current].next = old_to_new.get(current.next)
            old_to_new[current].random = old_to_new.get(current.random)
            current = current.next

        return old_to_new[head]