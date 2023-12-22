# Welcome to Practical-Go project

## Day1

* Go introduction (was devloped by Google)
* Basic Go commands

    * `go init application_name`
    * `go build `
    * `go run  `

## Day 2
* Strings  (Unicode/UTF-8)
* Rune -> In Go characters are represented in Uni-code point (int32). Runes can be 1 byte or 4 bytes.
* fomatting in Go (fmt package)

## Day 3
* https API calls (example: `curl -i https:\\` )
* Structure definition

       type struct_name struct{
            Name string # Capital letter indicates access modifier **public**
            public_repos int # small letter indicates access modifier **private**
        }
* **JSON <--> GO**
    * Data type mapping
        JSON | GO
        -----|-----
        true/false | true/false
        null | nil
        String | string
        Numbers | int8, int16, int32, int64, uint8...float32, float64..
        Array | []any or []interfaces
        Object | map[string]any

    * Encoding and Decoding
        JSON | GO
        -----|-----
        io.Reader | json.Decoder
        []byte | json.UnMarshal

        GO | JSON
        ---- | ----
        io.Writer | json.Encoder
        []byte | json.Marshal
 



