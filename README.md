# KMS

## Getting Started
1. Make sure you have [Go](https://go.dev) installed.
2. Clone the repo
```bash
git clone https://github.com/maulanar/kms.git
```
3. Go to the directory and run go mod tidy to add missing requirements and to drop unused requirements
```bash
cd kms && go mod tidy
```
3. Setup your .env file
```bash
cp .env-example .env && vi .env
```
4. Start
```bash
go run main.go
```

## Build for production
1. Pastikan docker & golang (1.21.3) terinstal
2. Buat folder "BE_KMS", buat folder "storages" didalam BE_KMS untuk attachments dan pastikan posisi user sedang berada didalam folder "BE_KMS" tsb.
3. Buat file .env COPY dari .env-example dan sesuaikan nilai :
    - APP_URL (default ke 127.0.0.1)
    - DB_DATABASE
    - DB_USERNAME
    - DB_PASSWORD
4. pastikan tidak ada image docker yang memiliki nama yg sama dengan yg akan di install, dengan run command ini : 
    4.1 ```bash docker stop kms ```
    4.2 ```bash docker rm kms ```
5. Download image kms terbaru, dengan run command : 
    ```bash docker pull maulanar/kms:latest ```
6. Jalankan image yang telah di download, lalu jalankan command ini (baca note dibawah terlebih dahulu): 
    ```bash docker run -d --name kms --restart always  --add-host=host.docker.internal:host-gateway  -p 4001:4001 -v /home/user-kms/be_kms/storages:/app/storages --env-file .env maulanar/kms:latest ```
    NOTE : ganti /home/user-kms/be_kms/storages dengan realpath folder storage pada point 2, bisa dicek dengan cmd : pwd
7. Dapat dicek dengan mengakses {{APP_URL}}:4001/api/v1/ping