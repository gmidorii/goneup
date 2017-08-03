# DB

## DB name
goneup

## DB list
- t_oneup

### t_oneup
| column       | type    | key     | Null | Default | overview |
|--------------|---------|---------|------|---------|----------|
| id           | INTEGER | primary | NO   |         |          |
| title        | TEXT    |         | NO   |         |          |
| value        | TEXT    |         | NO   |         |          |
| created_date | TEXT    |         | NO   |         |          |
| updated_date | TEXT    |         | NO   |         |          |

```sql
CREATE TABLE t_oneup {
	id INTEGER PRIMARY KEY,
	title TEXT,
	value TEXT,
	created_date TEXT,
	updated_date TEXT
}
```
