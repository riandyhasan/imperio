package operation

import "github.com/riandyhasan/imperio/model"

// generateRandomData simulates a new row insert.
func generateRandomData(schema *model.Schema) model.Schema {
	return *schema // Replace with actual data population logic
}

// generateUpdateData simulates a row update.
func generateUpdateData(schema *model.Schema) model.Schema {
	return *schema // Replace with actual update logic
}

// generateDeleteData simulates a row delete using primary key.
func generateDeleteData(schema *model.Schema) model.Schema {
	return *schema // Replace with key-only struct
}
