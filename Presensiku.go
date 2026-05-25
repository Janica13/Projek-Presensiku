package main

import (
	"fmt"
	"strings"
)

type Mahasiswa struct {
	NIM        string
	Nama       string
	JumlahAlpa int
}

type JadwalKelas struct {
	IDKelas    string
	MataKuliah string
	Ruangan    string
}

type LogKehadiran struct {
	IDKelas     string
	NIM         string
	PertemuanKe int
	Status      string
}

var daftarMahasiswa []Mahasiswa
var daftarJadwal []JadwalKelas
var daftarLog []LogKehadiran

func main() {
	var pilihan string 

	daftarJadwal = append(daftarJadwal, JadwalKelas{"IF-01", "Algoritma Pemrograman", "Lab 2"})

	for {
		fmt.Println("\n==================================================")
		fmt.Println("    SISTEM MONITORING PRESENSI (SiPresensi)       ")
		fmt.Println("==================================================")
		fmt.Println("A. Kelola Data Mahasiswa (Tambah, Ubah, Hapus)")
		fmt.Println("B. Catat Status Kehadiran Mahasiswa")
		fmt.Println("C. Cari Data Mahasiswa (Sequential & Binary)")
		fmt.Println("D. Urutkan Data Mahasiswa (Selection & Insertion)")
		fmt.Println("E. Tampilkan Statistik Kelas & Alpa Terbanyak")
		fmt.Println("X. Keluar Aplikasi")
		fmt.Println("==================================================")
		fmt.Print("Pilih Menu (A-E / X): ")
		
		fmt.Scanln(&pilihan)
		pilihan = strings.ToUpper(pilihan)

		switch pilihan {
		case "A":
			menuKelolaMahasiswa()
		case "B":
			catatStatusKehadiran()
		case "C":
			menuCariMahasiswa()
		case "D":
			menuUrutkanMahasiswa()
		case "E":
			tampilkanStatistik()
		case "X":
			fmt.Println("Terima kasih telah menggunakan SiPresensi.")
			return
		default:
			fmt.Println("Pilihan tidak tersedia! Gunakan huruf A, B, C, D, E, atau X.")
		}
	}
}

func menuKelolaMahasiswa() {
	for {
		fmt.Println("\n[SUB-MENU A] KELOLA DATA MAHASISWA")
		fmt.Println("1. Tambahkan Data Mahasiswa")
		fmt.Println("2. Mengubah Data Mahasiswa")
		fmt.Println("3. Menghapus Data Mahasiswa")
		fmt.Println("4. Lihat Semua Data Mahasiswa")
		fmt.Println("0. Kembali ke Menu Utama")
		fmt.Print("Pilih Aksi (0-4): ")
		
		var aksi int
		fmt.Scanln(&aksi)

		switch aksi {
		case 1:
			tambahMahasiswa()
		case 2:
			ubahMahasiswa()
		case 3:
			hapusMahasiswa()
		case 4:
			tampilkanSemuaMahasiswa()
		case 0:
			return
		default:
			fmt.Println("Aksi tidak valid!")
		}
	}
}

func tambahMahasiswa() {
	var m Mahasiswa
	fmt.Print("Masukkan NIM: ")
	fmt.Scanln(&m.NIM)
	fmt.Print("Masukkan Nama: ")
	fmt.Scanln(&m.Nama)
	m.JumlahAlpa = 0

	for _, mhs := range daftarMahasiswa {
		if mhs.NIM == m.NIM {
			fmt.Println("Gagal: Mahasiswa dengan NIM tersebut sudah ada.")
			return
		}
	}
	daftarMahasiswa = append(daftarMahasiswa, m)
	fmt.Println("Data mahasiswa berhasil ditambahkan!")
}

func ubahMahasiswa() {
	var nim string
	fmt.Print("Masukkan NIM mahasiswa yang ingin diubah: ")
	fmt.Scanln(&nim)

	for i := range daftarMahasiswa {
		if daftarMahasiswa[i].NIM == nim {
			fmt.Print("Masukkan Nama Baru: ")
			fmt.Scanln(&daftarMahasiswa[i].Nama)
			fmt.Println("Data mahasiswa berhasil diubah!")
			return
		}
	}
	fmt.Println("Mahasiswa tidak ditemukan.")
}

func hapusMahasiswa() {
	var nim string
	fmt.Print("Masukkan NIM mahasiswa yang ingin dihapus: ")
	fmt.Scanln(&nim)

	for i, m := range daftarMahasiswa {
		if m.NIM == nim {
			daftarMahasiswa = append(daftarMahasiswa[:i], daftarMahasiswa[i+1:]...)
			fmt.Println("Data mahasiswa berhasil dihapus!")
			return
		}
	}
	fmt.Println("Mahasiswa tidak ditemukan.")
}

func tampilkanSemuaMahasiswa() {
	if len(daftarMahasiswa) == 0 {
		fmt.Println("Data mahasiswa kosong.")
		return
	}
	fmt.Println("\n--- DAFTAR MAHASISWA ---")
	for _, m := range daftarMahasiswa {
		fmt.Printf("NIM: %s | Nama: %s | Total Alpa: %d\n", m.NIM, m.Nama, m.JumlahAlpa)
	}
}

func catatStatusKehadiran() {
	if len(daftarMahasiswa) == 0 {
		fmt.Println("Gagal: Masukkan data mahasiswa terlebih dahulu di Menu A!")
		return
	}

	var log LogKehadiran
	fmt.Println("\n[MENU B] CATAT STATUS KEHADIRAN PER PERTEMUAN")
	
	log.IDKelas = "IF-01" 
	
	fmt.Print("Masukkan NIM Mahasiswa: ")
	fmt.Scanln(&log.NIM)

	mhsIndex := -1
	for i, m := range daftarMahasiswa {
		if m.NIM == log.NIM {
			mhsIndex = i
			break
		}
	}
	if mhsIndex == -1 {
		fmt.Println("Gagal: NIM Mahasiswa tidak terdaftar.")
		return
	}

	fmt.Print("Pertemuan Ke: ")
	fmt.Scanln(&log.PertemuanKe)
	fmt.Print("Masukkan Status (hadir / izin / sakit / alpa): ")
	fmt.Scanln(&log.Status)
	log.Status = strings.ToLower(log.Status)

	if log.Status == "hadir" || log.Status == "izin" || log.Status == "sakit" || log.Status == "alpa" {
		daftarLog = append(daftarLog, log)
		
		if log.Status == "alpa" {
			daftarMahasiswa[mhsIndex].JumlahAlpa++
		}
		fmt.Println("Status kehadiran berhasil dicatat oleh sistem!")
	} else {
		fmt.Println("Gagal: Status harus diisi antara 'hadir', 'izin', 'sakit', atau 'alpa'.")
	}
}

func menuCariMahasiswa() {
	for {
		fmt.Println("\n[SUB-MENU C] PENCARIAN DATA MAHASISWA")
		fmt.Println("1. Cari Berdasarkan Status Kehadiran (Sequential Search)")
		fmt.Println("2. Cari Berdasarkan NIM (Binary Search)")
		fmt.Println("0. Kembali ke Menu Utama")
		fmt.Print("Pilih Metode Pencarian (0-2): ")

		var opsi int
		fmt.Scanln(&opsi)

		switch opsi {
		case 1:
			cariSequentialStatus()
		case 2:
			cariBinaryNIM()
		case 0:
			return
		default:
			fmt.Println("Pilihan tidak valid!")
		}
	}
}

func cariSequentialStatus() {
	if len(daftarLog) == 0 {
		fmt.Println("Belum ada data log kehadiran yang dicatat.")
		return
	}

	var statusTarget string
	fmt.Print("Masukkan Status Kehadiran yang dicari: ")
	fmt.Scanln(&statusTarget)
	statusTarget = strings.ToLower(statusTarget)

	fmt.Printf("\nHasil Pencarian (Sequential Search) Status '%s':\n", statusTarget)
	ditemukan := false
	for _, log := range daftarLog {
		if log.Status == statusTarget {
			fmt.Printf("- NIM: %s | Pertemuan Ke-%d | Status: %s\n", log.NIM, log.PertemuanKe, log.Status)
			ditemukan = true
		}
	}
	if !ditemukan {
		fmt.Println("Tidak ada data mahasiswa dengan status tersebut.")
	}
}

func cariBinaryNIM() {
	if len(daftarMahasiswa) == 0 {
		fmt.Println("Data mahasiswa kosong.")
		return
	}

	var targetNIM string
	fmt.Print("Masukkan NIM Mahasiswa yang dicari: ")
	fmt.Scanln(&targetNIM)

	for i := 0; i < len(daftarMahasiswa)-1; i++ {
		for j := i + 1; j < len(daftarMahasiswa); j++ {
			if daftarMahasiswa[i].NIM > daftarMahasiswa[j].NIM {
				daftarMahasiswa[i], daftarMahasiswa[j] = daftarMahasiswa[j], daftarMahasiswa[i]
			}
		}
	}

	low, high := 0, len(daftarMahasiswa)-1
	indexKetemu := -1

	for low <= high {
		mid := (low + high) / 2
		if daftarMahasiswa[mid].NIM == targetNIM {
			indexKetemu = mid
			break
		} else if daftarMahasiswa[mid].NIM < targetNIM {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	if indexKetemu != -1 {
		m := daftarMahasiswa[indexKetemu]
		fmt.Println("\n[Data Ditemukan dengan Binary Search]")
		fmt.Printf("NIM: %s | Nama: %s | Akumulasi Alpa: %d\n", m.NIM, m.Nama, m.JumlahAlpa)
	} else {
		fmt.Println("Mahasiswa dengan NIM tersebut tidak ditemukan.")
	}
}

func menuUrutkanMahasiswa() {
	for {
		fmt.Println("\n[SUB-MENU D] PENGURUTAN DATA MAHASISWA")
		fmt.Println("1. Urutkan Berdasarkan Jumlah Total Absensi/Alpa (Selection Sort - Terbanyak)")
		fmt.Println("2. Urutkan Berdasarkan Nama Mahasiswa (Insertion Sort - Ascending)")
		fmt.Println("0. Kembali ke Menu Utama")
		fmt.Print("Pilih Metode Pengurutan (0-2): ")

		var opsi int
		fmt.Scanln(&opsi)

		switch opsi {
		case 1:
			urutSelectionAlpa()
		case 2:
			urutInsertionNama()
		case 0:
			return
		default:
			fmt.Println("Pilihan tidak valid!")
		}
	}
}

func urutSelectionAlpa() {
	n := len(daftarMahasiswa)
	if n == 0 {
		fmt.Println("Data mahasiswa kosong.")
		return
	}

	for i := 0; i < n-1; i++ {
		maxIdx := i
		for j := i + 1; j < n; j++ {
			if daftarMahasiswa[j].JumlahAlpa > daftarMahasiswa[maxIdx].JumlahAlpa {
				maxIdx = j
			}
		}
		daftarMahasiswa[i], daftarMahasiswa[maxIdx] = daftarMahasiswa[maxIdx], daftarMahasiswa[i]
	}

	fmt.Println("\nData Berhasil Diurutkan (Selection Sort - Alpa Terbanyak):")
	tampilkanSemuaMahasiswa()
}

func urutInsertionNama() {
	n := len(daftarMahasiswa)
	if n == 0 {
		fmt.Println("Data mahasiswa kosong.")
		return
	}

	for i := 1; i < n; i++ {
		key := daftarMahasiswa[i]
		j := i - 1
		for j >= 0 && daftarMahasiswa[j].Nama > key.Nama {
			daftarMahasiswa[j+1] = daftarMahasiswa[j]
			j = j - 1
		}
		daftarMahasiswa[j+1] = key
	}

	fmt.Println("\nData Berhasil Diurutkan (Insertion Sort - Nama A-Z):")
	tampilkanSemuaMahasiswa()
}

func tampilkanStatistik() {
	if len(daftarMahasiswa) == 0 {
		fmt.Println("Belum ada data mahasiswa dalam sistem.")
		return
	}

	fmt.Println("\n[MENU E] STATISTIK PERSENTASE KEHADIRAN KELAS")
	
	totalLog := len(daftarLog)
	if totalLog == 0 {
		fmt.Println("Persentase Kehadiran Kelas: 0.00% (Belum ada log absensi).")
	} else {
		totalHadir := 0
		for _, log := range daftarLog {
			if log.Status == "hadir" {
				totalHadir++
			}
		}
		persentase := (float64(totalHadir) / float64(totalLog)) * 100
		fmt.Printf("Persentase Total Kehadiran Kelas (IF-01): %.2f%%\n", persentase)
	}

	maxAlpa := -1
	for _, m := range daftarMahasiswa {
		if m.JumlahAlpa > maxAlpa {
			maxAlpa = m.JumlahAlpa
		}
	}

	fmt.Println("\nDaftar Mahasiswa dengan Jumlah Alpa Terbanyak:")
	fmt.Println("-----------------------------------------------")
	adaAlpa := false
	for _, m := range daftarMahasiswa {
		if m.JumlahAlpa == maxAlpa && maxAlpa > 0 {
			fmt.Printf("- Nama: %s (NIM: %s) | Jumlah Alpa: %d\n", m.Nama, m.NIM, m.JumlahAlpa)
			adaAlpa = true
		}
	}
	if !adaAlpa {
		fmt.Println("Tidak ada mahasiswa yang membolos (Semua Alpa = 0).")
	}
	fmt.Println("-----------------------------------------------")
}