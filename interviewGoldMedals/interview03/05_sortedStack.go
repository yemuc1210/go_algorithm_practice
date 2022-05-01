package interview03

type stack struct{
	nums []int
}

//初始化栈
func newStack() *stack {
   return &stack{
	   nums: []int{},
   }
}

//栈的操作
func (s *stack) push(n int) {
   s.nums = append(s.nums, n)
}
func (s *stack) pop() int {
   if s.isEmpty() {
	   return -1
   }
   var res = s.nums[len(s.nums)-1] //取出第一个
   s.nums = s.nums[:len(s.nums)-1]
   return res
}
func (s *stack) top() int {
   if s.isEmpty() {
	   return -1
   }
   var res = s.nums[len(s.nums)-1] //取出第一个
   return res
}

func (s *stack) isEmpty() bool {
   if len(s.nums) == 0 {
	   return true
   }
   return false
}

type SortedStack struct {
	 s1 *stack
	 s2 *stack
}
func Constructor5() SortedStack {
	  return SortedStack{
		  s1:newStack(),
		  s2:newStack(),
	  }
}

func (this *SortedStack) Push(val int)  {
   
	  for {
		   if this.s1.isEmpty() || val<=this.s1.top(){
			  break 
		   }
			 this.s2.push(this.s1.top())
			 this.s1.pop()
	  } 
	  this.s1.push(val)
	   
	  for {
		  if this.s2.isEmpty() {
			 break
		  }
		  this.s1.push(this.s2.top())
		  this.s2.pop()
	  }
}


func (this *SortedStack) Pop(){
	  
	   if  this.s1.isEmpty(){
		  return
	   }
		this.s1.pop()
}


func (this *SortedStack) Peek() int { 
	 if this.s1.isEmpty(){
		 return -1
	 }
	 return this.s1.top()
}


func (this *SortedStack) IsEmpty() bool {  
		return this.s1.isEmpty()
}


/**
* Your SortedStack object will be instantiated and called as such:
* obj := Constructor();
* obj.Push(val);
* obj.Pop();
* param_3 := obj.Peek();
* param_4 := obj.IsEmpty();
**/