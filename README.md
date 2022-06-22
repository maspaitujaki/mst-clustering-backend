# MST-Based Clustering Backend
Repositori berisi kode program bagian backend untuk website penyelesaian masalah MST-Based Clustering. Kode program memiliki implementasi pembuatan minimum spanning tree menggunakan algoritma Kruskal. Selain itu, terdapat algoritma clustering dari hasil minimum spanning tree yang dihasilkan. Program dijalankan sebagai Restful API yang harus terhubung ke basis data relasional DBMS PostgresSQL.

## Teknologi dan Framework
- Golang
- GORM
- Gorilla Mux
- PostgresSQL

## Penjelasan Algoritma Kruskal
1. Program menerima dan sebuah file CSV dengan jumlah entri/baris sebesar. Data-data di translate menjadi Node-Node dengan mencatat nama, koordinat x, dan koordinat Y. O(N)
2. Akan dihitung jarak antara Node dan disimpan sebagai Edge. Penghitungan hanya dilakukan satu arah. Contoh nya bila sudah dihitung edge (1->2), maka edge (2->1) tidak dihitung. O(N(N-1)/2) = O(N^2)
3. Edges diurutkan menaik menggunakan algoritma quicksort. O(log N^2)
4. Menyiapkan tipe bentukan bernama Tree yang berisi array of Node dan array of Array. Membuat variabel array of Tree
5. Memulai traversal kepada edge-edge yang sudah terurut
6. Untuk setiap edge akan diperiksa node awal dan node tujuan
7. Apabila Node awal dan Node tujuan pada edge belum terdapat pada array Tree, maka dibuat Tree baru dengan isi Node awal dan Node tujuan serta edge yang sedang diperiksa, lalu dimasukkan ke dalam array Tree.
8. Apabila Node awal atau Node tujuan tedapat pada array Tree namun node lainnya belum ada, maka Node yang belum tercatat serta edge akan dimasukkan ke dalam Tree lawannya.
9. Apabila Node awal dan Node akhir terdapat pada array Tree namun pada pohon-pohon yang berbeda, maka kedua pohon akan digabung dengan menambahkan edge ke pohon baru sehingga array tree akan berkurang satu elemen.
10. Pencarian akan dihentikan apabila semua Node sudah terdapat pada array tree dan hanya ada satu tree di array tree.
11. Tree tersebut merupakan minimum spanning tree dengan Node sebanyak N dan edge sebanyak N-1.

## Analisis Kompleksitas
Pada traversal edges akan diperiksa node-node pada array tree sehingga worst case nya adalah O(N^3).

## Referensi Belajar
[Mod-01 Lec-36 Clustering using minimal spanning tree](https://www.youtube.com/watch?v=crBp-AYtBWA)