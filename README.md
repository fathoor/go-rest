# <p align="center">Go REST API (Gin, GORM, and PostgreSQL)</p>

##  <p align="center">DTS Kominfo x Hacktiv8 - Golang 2022</p>
 <p align="center"><i>This repository contains source code for Golang learning in <b>DTS Kominfo x Hacktiv8 2022 - Golang</b>.</i></p>

***

### Installation
1. Clone this repository
```bash
git clone https://github.com/fathoor/go-rest.git
```
2. Access the directory
```bash
cd go-rest
```
3. Install dependencies
```bash
go mod tidy
```
4. Configure environment variable in `.env` file
```bash
cp .env.example .env
```
5. Run the API
```bash
go run main.go
```
6. There are several endpoints that are accessible:

    | Endpoint | Method | Description |
    | --- | --- | --- |
    | `/orders` | *POST* | Add new order |
    | `/orders` | *GET* | Show all orders |
    | `/orders/:orderId` | *PUT* | Update order based on `orderId` |
    | `/orders/:orderId` | *DELETE* | Delete order based on `orderId` |

    The request body for these endpoints is in `JSON` format, for example:
    ```json
    {
        "customerName": "Customer Name",
        "items":[
            {
                "itemCode": "ItemCode",
                "description": "ItemDescription",
                "quantity": ItemQuantity
            },
            {
                "itemCode": "ItemCode",
                "description": "ItemDescription",
                "quantity": ItemQuantity
            },
            <!-- Add more if needed  -->
        ]
    }
    ```

***

<p align="center"><i>This repository is only for educational purpose, as I will use this only to upload my assignments for this particular class.</i></p>