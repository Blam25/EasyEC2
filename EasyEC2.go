// EasyEC2 project EasyEC2.go
package EasyEC2

//"log"

//"sync"

//"github.com/hajimehoshi/ebiten/v2"
//"github.com/hajimehoshi/ebiten/v2/ebitenutil"

var identifier int

func NewEntity() Entity {
	new := Entity{}
	new.id = identifier
	identifier++
	return new
}

type Entity struct {
	id int
}

func (s Entity) Getid() int {
	return s.id
}

type validComp interface {
	Getid() int
}

func NewComp[T validComp]() *Component[T] {
	new := &Component[T]{index: make(map[int]int)}
	return new
}

type Component[T validComp] struct {
	index    map[int]int
	theArray []T
}

func (s *Component[T]) Add(object T) {
	s.index[object.Getid()] = len(s.theArray)
	s.theArray = append(s.theArray, object)
}

func (s *Component[T]) Remove(id int) {
	//index of object to be removed
	index := s.index[id]

	//object to be removed
	object := s.theArray[index]

	//delete id and index of said object from map
	delete(s.index, object.Getid())
	//set value of deleted index to the last object in array, thereby deleting it
	s.theArray[index] = s.theArray[len(s.theArray)-1]
	//get id of moved index
	movedId := s.theArray[index].Getid()
	//set new index of moved object correctly in map
	s.index[movedId] = index
	//delete the last (now duplicated) object from the array
	s.theArray = s.theArray[:len(s.theArray)-1]
}

func (s *Component[T]) GetArr() []T {
	return s.theArray
}

func (s *Component[T]) Get(id int) T {
	if _, ok := s.index[id]; !ok {
		var zero T
		return zero
	}
	return s.theArray[s.index[id]]
}

func (s *Component[T]) Contains(id int) bool {
	if _, ok := s.index[id]; !ok {
		return false
	}
	return true
}
