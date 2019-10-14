import time
from typing import List

# 102. 二叉树的层次遍历

class TreeNode:
    def __init__(self, x:int, left=None, right=None):
        self.val = x
        self.left = left
        self.right = right

class Solution:
    def levelOrder(self, root: TreeNode) -> List[List[int]]:
        if root is None:
            return None
        result=[]
        queue = [[root,0]]
        while len(queue) != 0:
            item = queue.pop(0)
            if len(result) <= item[1] :
                result.append([item[0].val])
            else :
                result[item[1]].append(item[0].val)
            if item[0].left is not None :
                queue.append([item[0].left,item[1]+1])
            if item[0].right is not None :
                queue.append([item[0].right,item[1]+1])
        return result


if __name__ == '__main__':
    t = TreeNode(3,TreeNode(9),TreeNode(20,TreeNode(15),TreeNode(7)))
    print(Solution().levelOrder(t))










