package store

var (
	propertyRequestSQL = `SELECT property_name, property_value, default_value, settings.description, data_type, enabled, active_from_date,
		group_settings.group_code, group_settings.group_name, group_settings.description
		from public."SYSTEM_SETTINGS" settings
		full outer join public."SETTINGS_GROUP" group_settings
		on settings.group_code = group_settings.group_code;`

	propertyInsertSQL = `INSERT INTO public."SYSTEM_SETTINGS" (property_name, property_value, default_value,
						description, data_type, enabled, active_from_date,
						group_code) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`
)
