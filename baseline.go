package main

import (
	"fmt"
	"strings"
)

type Fileinfo struct {
	Docid       int
	FileContent string
}

type Word struct {
	Name          string
	Ft            int
	PointedToHead []InvertedListItem
}

type InvertedListItem struct {
	Docid int
	Fd    int
}

func main() {
	index := make(map[string]Word)
	docs := []Fileinfo{
		Fileinfo{Docid: 1, FileContent: "The old night keeper keeps the keep in the town"},
		Fileinfo{Docid: 2, FileContent: "In the big old house in the big old gown"},
		Fileinfo{Docid: 3, FileContent: "The house in the town had the big old keep"},
		Fileinfo{Docid: 4, FileContent: "Where the old nignt keeper never did sleep"},
		Fileinfo{Docid: 5, FileContent: "The night keeper  keeps the keep in the nignt"},
		Fileinfo{Docid: 6, FileContent: "And keeps in the dark and sleeps in the night"},
	}

	for _,value := range docs{
		wordSlice := strings.Fields(value.FileContent);
		invertedListItemMap :=  make(map[string]InvertedListItem)
		for _,wordItem :=  range wordSlice{
			if invertedListItemValue,ok:=invertedListItemMap[wordItem];ok{
				invertedListItemValue.Fd += 1
			}else{
				invertedListItemMap[wordItem] = InvertedListItem{Docid:value.Docid , Fd:1}
			}
		}
		for listItemKeyString ,listItem := range invertedListItemMap{
			if indexItem,indexOk := index[listItemKeyString] ;indexOk{
				indexItem.Ft += 1
				indexItem.PointedToHead =append(indexItem.PointedToHead,listItem) 
				index[listItemKeyString] = indexItem;
			}else{

				word :=  Word{Name:listItemKeyString,Ft:1,PointedToHead:nil};
				word.PointedToHead = append(word.PointedToHead , listItem);
				index[listItemKeyString] = word
			}
		}
	}
	for _,item := range index{

		fmt.Printf("%7s%3d|%v\n",item.Name,item.Ft,item.PointedToHead)
	}

}
