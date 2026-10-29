
open System
open System.Text.RegularExpressions

// 1. Reverse string and check palindrome
let reverseString (s: string) =
    new string(Array.rev (s.ToCharArray()))

let isPalindrome (s: string) =
    s = reverseString s

// 2. Transform string using pattern matching
let transformString (s: string) =
    match s with
    | x when x.Length > 0 && x |> Seq.forall Char.IsLetter ->
        s.ToUpper()
    | x when x.Length > 0 && x |> Seq.forall Char.IsDigit ->
        s + s
    | _ -> s

// 3. Validate email
let isValidEmail (email: string) =
    Regex.IsMatch(email, @"^[^@\s]+@[^@\s]+\.[^@\s]+$")

// Test 1: Reverse and palindrome
printfn "----- Reverse and Palindrome -----"
["madam"; "hello"; "level"]
|> List.iter (fun s ->
    printfn "%s -> Reverse: %s, Palindrome: %b"
        s (reverseString s) (isPalindrome s))

// Test 2: String transformation
printfn "\n----- String Transformation -----"
["hello"; "123"; "abc123"]
|> List.iter (fun s ->
    printfn "%s -> %s" s (transformString s))

// Test 3: Email validation
printfn "\n----- Email Validation -----"
["abc@gmail.com"; "wrongemail"; "user@test.org"]
|> List.iter (fun email ->
    printfn "%s -> %b" email (isValidEmail email))


