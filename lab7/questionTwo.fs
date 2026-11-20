
let numbers1 = [1; 2; 3; 4; 5]
let numbers2 = [2; 4; 6; 8]

// 1. List.map - Square each number
let square = List.map (fun x -> x * x)

printfn "Map: %A" (square numbers1)
printfn "Map: %A" (square numbers2)

// 2. List.filter - Select even numbers
let even = List.filter (fun x -> x % 2 = 0)

printfn "Filter: %A" (even numbers1)
printfn "Filter: %A" (even numbers2)

// 3. List.fold - Calculate sum
let sum = List.fold (fun acc x -> acc + x) 0

printfn "Fold: %d" (sum numbers1)
printfn "Fold: %d" (sum numbers2)
