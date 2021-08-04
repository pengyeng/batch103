package batch103

import "database/sql"

type WriterType interface {
	Write(data []BatchData)
	SetParameters(values map[string]string)
}

type BaseWriter struct {
	parameters map[string]string
}

type DBWriter struct {
	BaseWriter
	ConnectionString string
	DbType           string
}

func (b *BaseWriter) SetParameters(values map[string]string) {
	b.parameters = values
}

/* Parameters Getter */
func (b BaseWriter) Parameters() map[string]string {
	return b.parameters
}

func (d *DBWriter) DBConnect() *sql.DB {

	//Establish DB Connection
	db, err := sql.Open(d.DbType, d.ConnectionString)
	// if there is an error opening the connection, handle it
	if err != nil {
		panic(err.Error())
	}

	return db
}
