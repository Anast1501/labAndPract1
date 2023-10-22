package main

import (
	"errors"
)

//Структура HashMap представляет хэш-таблицу с фиксированным размером массива table
type HashMap struct {
	table [512]*Pair
}
//Структура Pair представляет пару ключ-значение, которую храним в таблице
type Pair struct {
	key   string
	value string
}

//Функция calcHash вычисляет хеш ключа и возвращает индекс в массиве table
//Если ключ пуст, то она возращает ошибку
func calcHash(key string, size int) (int, error) {
	if len(key) == 0 {
		return 0, errors.New("no value")
	}
	hash := 0
	for i := 0; i < len(key); i++ {
		hash += int(key[i])
	}
	return hash % size, nil
}

//Метод Insert вставляет пару ключ-значение в хеш-таблицу 
//Если ключ уже существует, то он возвращает ошибку
func (hmap *HashMap) Insert(key string, value string) error {
	p := &Pair{key, value}
	hash, err := calcHash(key, len(hmap.table))
	if err != nil {
		return errors.New("unacceptable key")
	}
	if hmap.table[hash] == nil {
		hmap.table[hash] = p
		return nil
	}
	if hmap.table[hash].key == key {
		return errors.New("this key already exists")
	}
	for i := (hash + 1) % len(hmap.table); i != hash; i = (i + 1) % len(hmap.table) {
		if hmap.table[i] == nil {
			hmap.table[i] = p
			return nil
		}
		if hmap.table[i].key == key {
			return errors.New("this key already exists")
		}
	}
	return errors.New("table is full")
}

//Метод Get получает значение по ключу из хеш-таблицы
//Если ключ не найден, то он возвращает ошибку 
func (hmap *HashMap) Get(key string) (string, error) {
	hash, err := calcHash(key, len(hmap.table))
	if err != nil {
		return "", errors.New("unacceptable key")
	}
	if hmap.table[hash] != nil && hmap.table[hash].key == key {
		return hmap.table[hash].value, nil
	}
	for i := (hash + 1) % len(hmap.table); i != hash; i = (i + 1) % len(hmap.table) {
		if hmap.table[i] != nil && hmap.table[i].key == key {
			return hmap.table[i].value, nil
		}
	}
	return "", errors.New("no such key")
}

//Метод Del удаляет элемент из хеш-таблицы по ключу
//Если ключ не найден, то он возвращает ошибку
func (hmap *HashMap) Del(key string) error {
	hash, err := calcHash(key, len(hmap.table))
	if err != nil {
		return errors.New("unacceptable key")
	}
	if hmap.table[hash] == nil {
		return errors.New("nothing to delete")
	}
	if hmap.table[hash].key == key {
		hmap.table[hash] = nil
		return nil
	}
	for i := (hash + 1) % len(hmap.table); i != hash; i = (i + 1) % len(hmap.table) {
		if hmap.table[i] != nil && hmap.table[i].key == key {
			hmap.table[i] = nil
			return nil
		}
	}
	return errors.New("no such key")
}

/*
func main() {
	// Создаем экземпляр HashMap
	hmap := HashMap{}

	// Вставляем пару ключ-значение
	err := hmap.Insert("n1", "a")
	if err != nil {
		fmt.Println("Ошибка при вставке:", err)
	}

	err = hmap.Insert("n2", "b")
	if err != nil {
		fmt.Println("Ошибка при вставке:", err)
	}

	err = hmap.Insert("n3", "c")
	if err != nil {
		fmt.Println("Ошибка при вставке:", err)
	}

	// Выводим все элементы через цикл
	for i := 0; i < len(hmap.table); i++ {
		if hmap.table[i] != nil {
			fmt.Printf("Ключ: %s, Значение: %s\n", hmap.table[i].key, hmap.table[i].value)
		}
	}
}
*/