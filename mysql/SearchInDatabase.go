package mysql

func SearchInDatabase(keyword string) []uint {
	db := ConnectMyDatabase()

	var Qids []uint
	query := "select Qid from Question where title like ? or content like ?"
	rows, err := db.Query(query, "%"+keyword+"%", "%"+keyword+"%")
	if err != nil {
		return nil
	}

	for rows.Next() {
		var qid uint
		if err := rows.Scan(&qid); err != nil {
			return nil
		}
		Qids = append(Qids, qid)
	}
	if err := rows.Err(); err != nil {
		return nil
	}
	defer db.Close()
	return Qids
}
