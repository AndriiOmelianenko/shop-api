# Shop API

API for some common online shop.

## Usage

```text
NAME:
   shop-api - Simple and lightweight Shop API

USAGE:
   app [global options] command [command options] [arguments...]

VERSION:
   0.1.0

COMMANDS:
     daemon, d  Start API
     seed, s    Seed the database with random values
     help, h    Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --mongo value   MongoDB address URL in format: mongodb://<user>:<password>@<address>:<port> (default: "mongodb://localhost:27017") [$MONGO]
   --dbname value  Database name for Shop (default: "shop") [$DBNAME]
   --help, -h      show help
   --version, -v   print the version
```

## Build

Build binary:

```bash
GOOS=linux go build -o ./app
```

Build docker image:

```bash
docker build -t shop-api .
```

## Deploy

Start Shop-API and MongoDB:

```bash
docker-compose up -d
```

Logs observation:

```bash
docker-compose logs -f api
```

## API endpoints

Client zone(/):

### General

* GET / - welcome page
    ```bash
    curl -X GET http://127.0.0.1:8080/
    ```
* GET /version - get version
    ```bash
    curl -X GET -H "Content-Type: application/json" http://127.0.0.1:8080/version
    ```

### Items

* GET /items - get list of items
    ```bash
    curl -X GET -H "Content-Type: application/json" http://127.0.0.1:8080/items
    ```
* GET /items/{itemID} - get one item
    ```bash
    curl -X GET -H "Content-Type: application/json" http://127.0.0.1:8080/items/<itemID>
    ```

### Categories

* GET /categories - get list of categories
    ```bash
    curl -X GET -H "Content-Type: application/json" http://127.0.0.1:8080/categories
    ```
* GET /categories/{categoryID} - get one category
    ```bash
    curl -X GET -H "Content-Type: application/json" http://127.0.0.1:8080/categories/<categoryID>
    ```

### Orders

* GET /orders - get list of orders
    ```bash
    curl -X GET -H "Content-Type: application/json" http://127.0.0.1:8080/orders
    ```
* GET /orders/{orderID} - get one order
    ```bash
    curl -X GET -H "Content-Type: application/json" http://127.0.0.1:8080/orders/<orderID>
    ```
* POST /orders - create an order
    ```bash
    curl -X POST -H "Content-Type: application/json" http://127.0.0.1:8080/orders -d '{"status": "new"}'
    ```
* POST /orders/{orderID}/items - add item to the order (by id and count)
    ```bash
    curl -X POST -H "Content-Type: application/json" http://127.0.0.1:8080/orders/<orderID> -d '{"item_id": "<itemID>", "item_cnt": 5}'
    ```

## Database

MongoDB was used as database.

### Entities

* item: id, alias, title, desc, price, pictures, count, categoryID, created, updated
* category: id, parent_id, alias, title, desc, logo, created, updated
* orders: id, sum, status, created, updated
* ordered: id, item sum, item count, orderID, itemID, created, updated

Each item belongs to the one category.

### Seed

With [fake](https://github.com/icrowley/fake) generate categories (two levels: 5-15 first level with 5-20 on each), items for 50-150 items on each category

To seed the database with random values:

```bash
docker-compose exec api ./app seed
```

## #TODO

### General

* add `/v1` prefix to api
* Write business logic for selecting, filtering of categories, items, create and update orders

### API

* GET /assets - for static files like pictures, js, css

Admin zone(/admin):

* import category {filePath}, simple csv format
* import logos {filePath}, should unpach zip file and update categories with corrisponding pictures (match by alias as suffics), store referenece to the files, handler for static files should be added
* import items {filePath}, simple csv format
* import pictures {filePath}, should unpach zip file and update items with corrisponding pictures (match by alias as suffics), store JSON string  with list of referenece to the files, handler for static files should be added
* update order {orderID} {status}, use data from csv file
* export orders {from-date} {to-date} {status}(optional), generate csv file with orders and suppor for filter by date and status
* POST /items - create new item
* PUT /items/{itemID} - update item info
* DELETE /items/{itemID} - delete item
* POST /items/{itemID}/picture - add picture to the item
* POST /categories - add new category
* PUT /categories/{categoryID} - update category info
* DELETE /categories/{categoryID} - delete category info
* PUT /orders/{orderID} - update order info
* POST /import/categories - upload list of categories (json)
* POST /import/items - upload list of items (json)
* POST /import/pictures - upload archive of pictures for existing items (zip, alias suffix matching)
* GET /export/orders - download list of orders
