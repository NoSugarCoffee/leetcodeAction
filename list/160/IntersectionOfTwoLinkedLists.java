class ListNode {
  int val;
  ListNode next;
  ListNode(int x) {
    val = x;
    next = null;
  }
}

public class IntersectionOfTwoLinkedLists {

  public ListNode getIntersectionNode(ListNode headA, ListNode headB) {
    ListNode[] arrayA = this.convertToArrayList(headA);
    ListNode[] arrayB = this.convertToArrayList(headB);

    int a = arrayA.length - 1;
    int b = arrayB.length - 1;

    ListNode val = null;
    while(a >= 0 && b >= 0) {
      if (arrayA[a] == arrayB[b]) {
        val = arrayA[a];
      }
      a--;
      b--;
    }
    return val;
  }
  public ListNode[] convertToArrayList(ListNode start) {
    List<ListNode> list = new ArrayList<>();
    while(start !=null) {
      list.add(start);
      start = start.next;
    }
    if (list.isEmpty()) {
      return new ListNode[]{};
    }
    return list.toArray(new ListNode[list.size()]);
  }
}