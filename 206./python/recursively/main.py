from typing import Optional
class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next
class Solution: 
  def reverseList(self, head: Optional[ListNode]) -> Optional[ListNode]:                                                                                       
    if not head or not head.next:                                                                                                                            
      return head                                                                                                                                          
                                                                                                                                                              
    new_head = self.reverseList(head.next)  # 先遞迴到最後                                                                                                   
    head.next.next = head  # 讓下一個節點指回自己                                                                                                            
    head.next = None       # 斷開原本的指向                                                                                                                  
                                                                                                                                                              
    return new_head 