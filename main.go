package main

import "fmt"
import "log"
import "github.com/geeksbaek/goinside"

type ArticleDraft struct {
	GallID  string
	Subject string
	Content string
	Images  []string
}

// 고정닉 세션 생성


func main() {
  fmt.Println("Hello World")
  s, err := goinside.Login("bangbaechick", "gusahr1")
  if err != nil {
          log.Fatal(err)
  }
  draft := NewArticleDraft("iu_new", "ㅌㅅㅌ", "ㅌㅅㅌ", "이미지 경로")
  if err := s.Write(draft); err != nil {
          log.Fatal(err)
  }



}

func NewArticleDraft(gallID, subject, content string, images ...string) *ArticleDraft {
	return &ArticleDraft{gallID, subject, content, images}
}

