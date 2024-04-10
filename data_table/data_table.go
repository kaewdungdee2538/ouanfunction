package data_table

// DataRow represents a row in the DataTable.
type DataRowMessage struct {
	Code          int
	MassageNameTh string
	MassageNameEn string
}

// DataTable represents the table with ID as the key.
type DataTableMessage struct {
	data map[int]DataRowMessage
}

// NewDataTable initializes a new DataTable.
func NewDataTable() *DataTableMessage {
	return &DataTableMessage{
		data: make(map[int]DataRowMessage),
	}
}

// AddRow adds a row to the DataTable with the specified ID.
func (dt *DataTableMessage) AddRowMessage(id int, nameTh string, nameEn string) {
	dt.data[id] = DataRowMessage{Code: id, MassageNameTh: nameTh, MassageNameEn: nameEn}
}

// FindRowByID retrieves a row from the DataTable by its ID.
func (dt *DataTableMessage) FindRowMessageByID(id int) (DataRowMessage, bool) {
	row, ok := dt.data[id]
	return row, ok
}
