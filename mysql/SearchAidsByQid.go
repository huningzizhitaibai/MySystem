package mysql

func SearchAidsByQid(qid int) []int {
	db := ConnectMyDatabase()

	var Aids []int
	query := "select Aid from Answer where ConnectQuestion = ?"
	rows, err := db.Query(query, qid)
	if err != nil {
		return nil
	}

	for rows.Next() {
		var aid int
		if err := rows.Scan(&aid); err != nil {
			return nil
		}
		Aids = append(Aids, aid)
	}
	if err := rows.Err(); err != nil {
		return nil
	}
	defer db.Close()
	return Aids
}
