# CSV 2 JSON

## This is just a simple API written in GO that reads a standard CSV file with the first row as a the header

### once read, an http json response is sent back to the client with the CSV data in json format. One of the beauties of Go is the use map[string]interface{} or just the ability to use an empty interface for arbitrary data

you can just hold everything in a simple struct

```go
type CSVData struct {
 Header []string                 `json:"header"`
 Rows   []map[string]interface{} `json:"data"`
}
```

Single Endpoint

```bash
POST   /api/v1/upload
```

multi-part ( file )

### example

### CSV File

```console
"Code","Name","Category","Quantity"
"f230fh0g3","Bamboo Watch","Accessories","24"
"nvklal433","Black Watch","Accessories","61"
"zz21cz3c1","Blue Band","Fitness","2"
"244wgerg2","Blue T-Shirt","Clothing","25"
"h456wer53","Bracelet","Accessories","73"
"av2231fwg","Brown Purse","Accessories","0"
"bib36pfvm","Chakra Bracelet","Accessories","5"
"mbvjkgip5","Galaxy Earrings","Accessories","23"
"vbb124btr","Game Controller","Electronics","2"
"cm230f032","Gaming Set","Electronics","63"
````

```json
{
    "header": [
        "Code",
        "Name",
        "Category",
        "Quantity"
    ],
    "data": [
        {
            "Category": "Accessories",
            "Code": "f230fh0g3",
            "Name": "Bamboo Watch",
            "Quantity": "24"
        },
        {
            "Category": "Accessories",
            "Code": "nvklal433",
            "Name": "Black Watch",
            "Quantity": "61"
        },
        {
            "Category": "Fitness",
            "Code": "zz21cz3c1",
            "Name": "Blue Band",
            "Quantity": "2"
        },
        {
            "Category": "Clothing",
            "Code": "244wgerg2",
            "Name": "Blue T-Shirt",
            "Quantity": "25"
        },
        {
            "Category": "Accessories",
            "Code": "h456wer53",
            "Name": "Bracelet",
            "Quantity": "73"
        },
        {
            "Category": "Accessories",
            "Code": "av2231fwg",
            "Name": "Brown Purse",
            "Quantity": "0"
        },
        {
            "Category": "Accessories",
            "Code": "bib36pfvm",
            "Name": "Chakra Bracelet",
            "Quantity": "5"
        },
        {
            "Category": "Accessories",
            "Code": "mbvjkgip5",
            "Name": "Galaxy Earrings",
            "Quantity": "23"
        },
        {
            "Category": "Electronics",
            "Code": "vbb124btr",
            "Name": "Game Controller",
            "Quantity": "2"
        },
        {
            "Category": "Electronics",
            "Code": "cm230f032",
            "Name": "Gaming Set",
            "Quantity": "63"
        }
    ]
}
```
