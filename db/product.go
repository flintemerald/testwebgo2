package db

func GetProductsList() []string {
  res := []string{}
  rows, err := con.Query("SELECT id, name_rus, name_eng FROM product LIMIT 9000")
  if err != nil {
    return res
  }
  for rows.Next() {
    var id int64
    var name_rus string
    var name_eng string
    rows.Scan(&id, &name_rus, &name_eng)
    res = append(res, name_rus)
  }
  return res
}
