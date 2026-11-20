
open Microsoft.Data.Sqlite

use con = new SqliteConnection("Data Source=student.db")
con.Open()

let cmd = con.CreateCommand()
cmd.CommandText <- "CREATE TABLE IF NOT EXISTS Student (Id INTEGER, Name TEXT)"
cmd.ExecuteNonQuery() |> ignore

// CREATE
cmd.CommandText <- "INSERT INTO Student VALUES (1, 'Aman')"
cmd.ExecuteNonQuery() |> ignore
printfn "Student Added"

// READ
cmd.CommandText <- "SELECT * FROM Student"
use reader = cmd.ExecuteReader()
while reader.Read() do
    printfn "%d %s" (reader.GetInt32(0)) (reader.GetString(1))
reader.Close()

// UPDATE
cmd.CommandText <- "UPDATE Student SET Name='Rahul' WHERE Id=1"
cmd.ExecuteNonQuery() |> ignore
printfn "Student Updated"

// DELETE
cmd.CommandText <- "DELETE FROM Student WHERE Id=1"
cmd.ExecuteNonQuery() |> ignore
printfn "Student Deleted"
