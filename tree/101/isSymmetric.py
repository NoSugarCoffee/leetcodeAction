class TreeNode:
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None

# 101. 对称二叉树
# 给定一个二叉树，检查它是否是镜像对称的。

# 采用中序遍历，并输出到数组的方式
# 会遇到
#   1
#  2 2
# 2 2
# 此种形式的左中右和右中左输出的是一致的
# 添加层号，一起判断就行
class Solution:
    def __init__(self):
        self.list = []

    def isSymmetric(self, root: TreeNode) -> bool:
        print(len(self.list))
        self.inOrder(root,1)
        for i in range(len(self.list)//2):
            if self.list[i] != self.list[len(self.list) - i -1]:
                return False
        return True

    def inOrder(self,root: TreeNode,level :int):
        if root is None:
            return
        self.inOrder(root.left,level + 1)
        self.list.append({root.val:level})
        self.inOrder(root.right,level + 1)





