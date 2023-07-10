package main

import (
	"fmt"
	"os"
	
)

type Card struct{
	English string
	Russian string
}

type Inserter interface{
	insert(c Card)error
	
}
type Displayer interface{
	display()
}

type Searcher interface {
	search(w string) error
}

type Worker interface {
	Inserter
	Displayer
	Searcher
}


func NewCard() *Card{
	var eng, rus string
	fmt.Println("Заполните карточку")
	fmt.Println("English: ")
	fmt.Scan(&eng)
	fmt.Println("Русское: ")
	fmt.Scan(&rus)
	return &Card{
		English: eng,
		Russian: rus,
	}
}


type StorageMap struct {
	data map[string]string
}

func NewStorageMap() *StorageMap{
	return &StorageMap{
		data: make(map[string]string, 0),
	}
}

func (s *StorageMap) insert(c Card) error {
	s.data[c.English] = c.Russian
	return nil
}

func (s *StorageMap) display(){
	for k, v := range s.data{
		fmt.Println(k,":", v)
	}
}

func (s *StorageMap) search(word string)  error {
	_, ok := s.data[word]
	if ok {
		fmt.Printf("Перевод:%s - %s\n",word, s.data[word])
	}else{
		fmt.Println("Нет такого слова в базe")
	}
	
	return nil
}


type StorageList struct{
	data []Card
}

func NewStorageList() *StorageList{
	return &StorageList{
		data: make([]Card, 0),
	}
}

func (s *StorageList) insert(c Card) error{
	s.data = append(s.data, c)
	return nil
}


func (s *StorageList) display(){
	for k, v := range s.data{
		fmt.Println(k,":", v)
	}
}

func (s *StorageList) search(word string) error {
	for _, v := range s.data{
		if word == v.English{
			fmt.Printf("Перевод:%s - %s\n",word, v.Russian)
		}else{
			fmt.Println("Нет такого слова в базe")
		}

	}
    
	

	return nil
}


func allTypeInsert(s Worker, c Card) error{
	s.insert(c)
	return nil
}

func DisplayData(s Worker){
	s.display()
}

func AllSearch(s Worker, word string){
	fmt.Println("Ищем слово: ")
	fmt.Scan(&word)
	s.search(word)
}



func main()  {
	var ch int
	var word string 
	data_dict := NewStorageMap()
	data_list := NewStorageList()

	var m, l Worker = &StorageMap{data_dict.data}, &StorageList{data_list.data} 

	

	for {
		fmt.Println("Введите цифру 1-ввод, 2-распечатать", 
					"3-распечатать 4-search key 5-exit")
		fmt.Scan(&ch)
		switch ch {
		case 1:
			card := NewCard()
			err := allTypeInsert(m, *card)
			if err != nil{
				fmt.Println(err.Error())
			}
			allTypeInsert(l, *card)
			if err != nil {
				fmt.Println(err.Error())
			}
		case 2:
			DisplayData(m)
		case 3:
			DisplayData(l)
		case 4:
			AllSearch(m, word)
		case 5:
			AllSearch(l, word)	
		case 6:
			os.Exit(0)
		}
	}




	

	
	


}