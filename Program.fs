// printfn "Hello from F#"

type Next<'T> =
    | Some of ref<'T>
    | None

type LinkedList<'T> =
    struct
        val Value: 'T
        val Next: Next<LinkedList<'T>>
        new(v: 'T, n: Next<LinkedList<'T>>) = { Value = v; Next = n }
    end

let b = new LinkedList<string>("b value", None)
let a = new LinkedList<string>("a value", Some(ref b))

System.Console.WriteLine(a.Value)
System.Console.WriteLine(b.Value)
