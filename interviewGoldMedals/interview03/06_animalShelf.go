package interview03

import "container/list"

/*
面试题 03.06. 动物收容所

"先进先出"，收养人只能收养最老的

队列
*/

//使用list包
type AnimalShelf struct {
	listAnimal *list.List
}


func Constructor6() AnimalShelf {
	return AnimalShelf{listAnimal: list.New()}
}


func (this *AnimalShelf) Enqueue(animal []int)  {
	this.listAnimal.PushBack(animal)
}


func (this *AnimalShelf) DequeueAny() []int {
	if this.listAnimal.Len() == 0{
		return []int{-1,-1}
	}
	animal := this.listAnimal.Front()
	defer this.listAnimal.Remove(animal)
	return animal.Value.([]int)

}


func (this *AnimalShelf) DequeueDog() []int {
	for animal := this.listAnimal.Front(); animal != nil; animal = animal.Next(){
		if animal.Value.([]int)[1] == 1 {
			defer this.listAnimal.Remove(animal)
			return animal.Value.([]int)
		}
	}
	return []int{-1, -1}
}


func (this *AnimalShelf) DequeueCat() []int {
	for animal := this.listAnimal.Front(); animal != nil; animal = animal.Next(){
		if animal.Value.([]int)[1] == 0 {
			defer this.listAnimal.Remove(animal)
			return animal.Value.([]int)
		}
	}
	return []int{-1, -1}

}


/**
 * Your AnimalShelf object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Enqueue(animal);
 * param_2 := obj.DequeueAny();
 * param_3 := obj.DequeueDog();
 * param_4 := obj.DequeueCat();
 */