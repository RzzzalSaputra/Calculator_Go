# **Simple Go Calculator (CLI)**

Program ini adalah kalkulator sederhana berbasis **Command Line Interface (CLI)** yang ditulis menggunakan bahasa **Go**.  
Aplikasi ini menyediakan empat operasi aritmetika dasar serta menerapkan konsep **function as parameter** untuk memanggil fungsi operasi.

---

## **Fitur Utama**
- **Pertambahan**
- **Pengurangan**
- **Perkalian**
- **Pembagian**
- **Menu berulang** sampai pengguna memilih keluar
- Penggunaan **fungsi sebagai parameter** untuk modularitas

---

## **Struktur Program**

### **1. main()**
- Menampilkan menu kalkulator  
- Menerima pilihan dari pengguna  
- Menjalankan fungsi operasi sesuai menu  
- Mengatur perulangan program

### **2. inputAngka(f func(int, int))**
- Meminta dua angka dari pengguna  
- Memanggil fungsi operasi (add, sub, multi, div) berdasarkan parameter yang diterima

### **3. Fungsi Operasi**
- **add()** – Pertambahan  
- **sub()** – Pengurangan  
- **multi()** – Perkalian  
- **div()** – Pembagian  

Semua fungsi mencetak hasil langsung ke terminal.

---

## **Cara Menjalankan Program**

1. Pastikan Go telah terpasang.  
   Periksa dengan perintah:
   ```bash
   go version
   ```
2. Simpan kode sebagai:
   ```text
   main.go
   ```
3. Jalankan program:
   ```bash
   go run main.go
   ```

---

## **Contoh Output**
```
Pilihan Menu Kalkulator
1. Pertambahan
2. Pengurangan
3. Perkalian
4. Pembagian

Masukkan Pilihan Menu: 1
Masukkan Angka 1: 10
Masukkan Angka 2: 5
Hasil Pertambahan 10 + 5 = 15

Apakah Ingin Lanjut atau Tidak (Ketik (y) jika lanjut (n) jika tidak): y
```

---

## **Catatan**
- Operasi pembagian menggunakan **integer division**.  
- Hindari memasukkan angka **0** sebagai pembagi untuk mencegah error pada saat runtime.

