# BooksApi
Rest Api for listing books


###  Usage 
go run main.go 

Docker image  had been created for this server sake: sidikhub/bookapi:latest

```bash
docker run -p 8000:8000 sidikhub/bookapi:latest
```

## Client 
localhost:8000/api/books //See all books

localhost:8000/api/book //Add book 
{
	{
	    "title":"Au pays des merveilles",
	    "author":{
	        "firstname":"Joel",
	        "lastname":"sidikou"
	    }
	}
}

localhost:8000/api/books/1 // Modify book

localhost:8000/api/books/2 // Delete book
