package main
import( //импортирование 1 пакета: errors для создания и возврата ошибок
	"errors" 
)

//Массив
type Array struct{
	data []string
	length int
}

//NewArray создаёт новый пустой массив
func NewArray() *Array {
	return &Array{}
}

//Append добавляет элемент в конец массива
func (arr *Array) ArAdd(value string) {
	arr.data = append(arr.data, value)
	arr.length++
}

//Get возвращает элемент по индексу
func (arr *Array) AGet(index int) (string, error){
	if index<0 || index>= arr.length{
		return "0", errors.New("Index out of range")
}
return arr.data[index], nil
}

//Set устанавливает значение элемента по индексу 
func (arr *Array) ASet(index int, value string) error{
	if index < 0|| index >= arr.length{
		return errors.New("Index out of range")
}
arr.data[index] = value
return nil
}

//Length возвращает текущую длину массива
func (arr *Array) ALength() int{
	return arr.length
}

//Удаление элементов массива и смещение 
func (arr *Array) ADel(index int) string{
	if index > 0 && index < arr.length{
		deletedValue:=arr.data[index] //Сохраняем значение элемента
		copy(arr.data[index:], arr.data[index+1:]) //Смещаем элементы влево на одну позицию
		arr.data=arr.data[:len(arr.data)-1]
		arr.length-- //Уменьшаем длину массива 
		return deletedValue //Возвращаем удалённое значение
	}
	return "index not found"
}

