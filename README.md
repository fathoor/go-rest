# <p align="center">Go REST API (Gin, GORM, and PostgreSQL)</p>

##  <p align="center">DTS Kominfo x Hacktiv8 - Golang 2022</p>
 <p align="center">Repository ini berisi source code untuk pembelajaran Golang di <b>FGA Kominfo x Hacktiv8 2022 - Golang</b>.</p>

***

### Instalasi
1. Clone repository ini
```bash
git clone https://github.com/fathoor/go-rest.git
```
2. Masuk ke direktori repository
```bash
cd go-rest
```
3. Jalankan perintah berikut untuk menginstall dependensi
```bash
go mod tidy
```
4. Konfigurasi nilai variabel pada file `.env`
```bash
cp .env.example .env
```
5. Jalankan perintah berikut untuk menjalankan aplikasi
```bash
go run main.go
```
6. Tersedia beberapa endpoint yang dapat diakses, antara lain:

    | Endpoint | Method | Deskripsi |
    | --- | --- | --- |
    | /orders | POST | Menambahkan daftar order |
    | /orders | GET | Menampilkan seluruh daftar order |
    | /orders/:orderId | PUT | Mengubah daftar order berdasarkan `orderId` |
    | /orders/:orderId | DELETE | Menghapus daftar order berdasarkan `orderId` |

    Bentuk request body untuk endpoint di atas sebagai berikut:
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