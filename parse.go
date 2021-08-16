package configer

func PostgresMapToDsn(m map[string]interface{}) string {

	var dsn string
	for k, v := range m {
		//TODO	Maybe we should use a better performance string splicing scheme?
		dsn += k + "=" + v.(string) + " "
	}
	return dsn
}
