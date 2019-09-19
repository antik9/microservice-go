package db

const schema = `
CREATE TABLE IF NOT EXISTS events (
	day DATE,
	beginning TIMESTAMP,
	endofevent TIMESTAMP,
	name VARCHAR(255),
	eventtype INTEGER,
	is_sent BOOLEAN
)
`

func InitialMigration() {
	calendar, _ := NewCalendar()
	calendar.conn.MustExec(schema)
}
