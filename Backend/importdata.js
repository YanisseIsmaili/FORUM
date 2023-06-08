const sqlite3 = require('sqlite3').verbose();
 
// open the database
let db = new sqlite3.Database('database.go');
 
let sql = `SELECT id, nom FROM societes WHERE id  = ?`;
let Id = 1;
 
// first row only
db.get(sql, [Id], (err, row) => {
  if (err) {
    return console.error(err.message);
  }
  return row
    ? console.log(row.id, row.nom)
    : console.log(`No data found with the id ${Id}`);
 
});
 
// close the database connection
db.close();