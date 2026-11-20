
open System
open Microsoft.Data.Sqlite

// 1. Student Record
type Student = {
    Id: int
    Name: string
    Marks: int
}

// 2. Student Data
let students = [
    {Id=1; Name="Aman"; Marks=80}
    {Id=2; Name="Rahul"; Marks=90}
    {Id=3; Name="Karan"; Marks=60}
]

// 3. Higher Order Functions
let names = List.map (fun s -> s.Name) students
let passed = List.filter (fun s -> s.Marks >= 40) students

printfn "Names: %A" names
printfn "Passed Students: %A" passed

// 4. Pattern Matching
let result marks =
    match marks with
    | m when m >= 40 -> "Pass"
    | _ -> "Fail"

students |> List.iter (fun s ->
    printfn "%s: %s" s.Name (result s.Marks))

// 5. SQLite Database
use con = new SqliteConnection("Data Source=student.db")
con.Open()

let cmd = con.CreateCommand()
cmd.CommandText <- "CREATE TABLE IF NOT EXISTS Student (Id INT, Name TEXT, Marks INT)"
cmd.ExecuteNonQuery() |> ignore

// CREATE
cmd.CommandText <- "INSERT INTO Student VALUES (1,'Aman',80)"
cmd.ExecuteNonQuery() |> ignore

// READ
cmd.CommandText <- "SELECT * FROM Student"
use reader = cmd.ExecuteReader()
while reader.Read() do
    printfn "%d %s %d"
        (reader.GetInt32(0))
        (reader.GetString(1))
        (reader.GetInt32(2))
reader.Close()

// UPDATE
cmd.CommandText <- "UPDATE Student SET Marks=95 WHERE Id=1"
cmd.ExecuteNonQuery() |> ignore

// DELETE
cmd.CommandText <- "DELETE FROM Student WHERE Id=1"
cmd.ExecuteNonQuery() |> ignore

printfn "CRUD Operations Completed"

// 6. Async Workflow
let work = async {
    do! Async.Sleep(1000)
    printfn "Async Task Completed"
}

Async.RunSynchronously work
