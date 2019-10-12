class TreeNode:
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None

class Solution:
    def isLessThanRight(self,node :TreeNode) -> bool:
        todo :TreeNode = node.right
        print("xxx",todo)
        if todo == None:
            return True

        while todo.left != None:
            todo = todo.left

        if todo.val > node.val:
            return True
        return False

    def isMoreThanLeft(self,node :TreeNode) -> bool:
        todo :TreeNode = node.left
        if todo == None:
            return True
        while todo.right != None :
            todo = todo.right

        if todo.val < node.val:
            return True
        return False

    def traverse(self,node: TreeNode) -> bool:
        if node == None:
            return True

        if  self.traverse(node.left) == False:
            return False

        if  self.traverse(node.right) == False:
            return False

        # 每个结点都比左边大，比右边的小
        return self.isMoreThanLeft(node) and self.isLessThanRight(node)

    def isValidBST(self, root: TreeNode) -> bool:
        return self.traverse(root)

if __name__ == '__main__':
    t0 = TreeNode(5)
    t1 = TreeNode(1)
    t2 = TreeNode(4)
    t3 = TreeNode(3)
    t4 = TreeNode(6)
    t0.left = t1
    t0.right = t2
    t2.left= t3
    t2.right= t4

    print(Solution().isValidBST(t0))




