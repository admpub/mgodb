package mgodb

type ResultSet map[string]interface{}

func (r ResultSet) MakeId() interface{} {
	return r["_id"]
}
