
open MongoDB.Driver
open MongoDB.Bson

let client = MongoClient("mongodb://localhost:27017")
let db = client.GetDatabase("College")
let col = db.GetCollection<BsonDocument>("Student")

let id = ObjectId.GenerateNewId()

// CREATE
let student = BsonDocument()
student.Add("_id", BsonObjectId(id))
student.Add("Name", BsonString("Aman"))
student.Add("Age", BsonInt32(20))
col.InsertOne(student)
printfn "Student Added"

// READ
let data = col.Find(Builders<BsonDocument>.Filter.Empty).ToList()
data |> Seq.iter (printfn "%O")

// UPDATE
let filter = Builders<BsonDocument>.Filter.Eq("_id", id)
let update = Builders<BsonDocument>.Update.Set("Name", "Rahul")
col.UpdateOne(filter, update)
printfn "Student Updated"

// DELETE
col.DeleteOne(filter)
printfn "Student Deleted"
