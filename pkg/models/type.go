package models

import (
  "errors"
  "github.com/asyoume/postgres"
  "strings"
)

type Test1 struct { 
    Test1   string `json:"test1"`
    Test2   string `json:"test2"`
}

func Test1Reflect(s interface{}, column []string) (*[]interface{},*string, error) {
  rel := make([]interface{}, 0, 10)
  rel_str := ""
  sr := s.(*Test1)
  for _, v := range column {
    rel_str = rel_str + v +","
    switch v { 
    case "test1":
      rel = append(rel, &(sr.Test1))
    case "test2":
      rel = append(rel, &(sr.Test2))
    default:
      rel_str = ""
      return &rel,&rel_str, errors.New(v + ",字段不存在")
    }
  }
  rel_str= strings.TrimRight(rel_str, ",")
  return &rel,&rel_str, nil
}

func Test1New() interface{}{
  return &Test1{}
}

func Test1Check(s interface{}) bool{
  _, ok := s.(*Test1)
  return ok
}

func Test1Check2(s interface{}) bool{
  _, ok := s.(*[]Test1)
  return ok
}

func Test1Add(all interface{},s interface{}){
   all_data := all.(*[]Test1)
   sr := *s.(*Test1)

   new_sr := Test1{}
   new_sr = sr
  *all_data = append(*all_data, new_sr)
}

func init() {
  postgres.SqlFuncMap["test1"] = Test1Reflect
  postgres.SqlNewMap["test1"] = Test1New
  postgres.SqlAddMap["test1"] = Test1Add
  postgres.SqlCheckMap["test1"] = Test1Check
  postgres.SqlCheck2Map["test1"] = Test1Check2
}