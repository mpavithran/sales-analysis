# **Sales Analysis System (Golang + MySQL + GORM)**

This project provides a **Sales Analysis System** that allows users to fetch **revenue data, top products, and order details** using a **RESTful API** built with **Golang (Gin), MySQL, and GORM**.

---

## **📌 Prerequisites**

Before running the project, ensure you have the following installed:

1. **Golang (>= 1.18)** – [Download Golang](https://golang.org/dl/)
2. **MySQL (>= 8.0)** – [Download MySQL](https://dev.mysql.com/downloads/)
3. **Git** – [Download Git](https://git-scm.com/)
4. **Postman** _(optional, for testing APIs)_ – [Download Postman](https://www.postman.com/)

---

## **🚀 Step-by-Step Execution**

### **1️⃣ Clone the Repository**

```sh
git clone https://github.com/your-repo/sales-analysis.git
cd sales-analysis
```

---

### **2️⃣ Set Up Environment Variables**

Create a `.env` file in the project root and add the following:

```
DB_USER=root
DB_PASSWORD=yourpassword
DB_HOST=localhost
DB_PORT=3306
DB_NAME=sales_db
SERVER_PORT=8080
```

_(Update credentials based on your MySQL setup.)_

---

### **3️⃣ Install Dependencies**

```sh
go mod tidy
```

---

### **4️⃣ Set Up the Database**

Log in to MySQL and create the database:

```sql
CREATE DATABASE sales_db;
```

---

### **5️⃣ Run the Server**

```sh
go run main.go
```

**✅ The server will start at** `http://localhost:8080`

---

## **📌 API Endpoints**

| **Endpoint**    | **Method** | **Request Body (JSON)**                                                                                    | **Description**                    |
| --------------- | ---------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------- |
| `/upload-csv`   | `POST`     | **Multipart Form-Data:** Upload a CSV file containing order details                                        | Bulk upload orders from a CSV file |
| `/top-products` | `GET`      | Query Params: `date_from`, `date_to`, `category` _(optional)_, `region` _(optional)_, `top` _(default: 5)_ | Fetch top-selling products         |
| `/revenue`      | `GET`      | Query Params: `date_from`, `date_to`, `product`, `category`, `region` _(optional)_                         | Get total revenue                  |

---

Here’s the updated **README.md** to reflect the **CSV upload feature**:

---

## **🛠 Upload Orders via CSV**

### **1️⃣ CSV Format**

Your CSV file should have the following columns:

```
Order ID,Product ID,Customer ID,Product Name,Category,Region,Date of Sale,Quantity Sold,Unit Price,Discount,Shipping Cost,Payment Method,Customer Name,Customer Email,Customer Address
```

Example:

```
1001,P123,C456,UltraBoost Running Shoes,Shoes,North America,2023-12-15,2.00,180,0.1,10,Credit Card,John Smith,johnsmith@email.com,"123 Main St, Anytown, CA 12345"
```

---

### **2️⃣ Upload via Postman**

1. **Open Postman**
2. Select **POST** request
3. Enter `http://localhost:8080/upload-csv`
4. Go to **Body** → Select **form-data**
5. **Key**: `file`, **Type**: `File`, and **Choose your CSV file**
6. Click **Send**

---

### **3️⃣ Upload via cURL**

```sh
curl -X POST -F "file=@orders.csv" http://localhost:8080/upload-csv
```
