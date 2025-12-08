package main

import "fmt"

func main() {
	var pilihan int
	var nilai float64
	var lanjut string = "y"

	fmt.Print("\033[H\033[2J")

	fmt.Println("╔══════════════════════════════════════════════════════════╗")
	fmt.Println("                  🌡 KONVERTER SUHU CLI 🌡                  ")
	fmt.Println("╚══════════════════════════════════════════════════════════╝")
	fmt.Println()

	fmt.Println("══════════════════════════════════════════════════════════")
	fmt.Println("                     📋 PETUNJUK PENGGUNAAN")
	fmt.Println(" 1. Pilih jenis konversi dari menu")
	fmt.Println(" 2. Masukkan nilai suhu")
	fmt.Println(" 3. Lihat hasil konversi")
	fmt.Println(" 4. Pilih 'y' untuk konversi lagi atau 'n' untuk keluar")
	fmt.Println("══════════════════════════════════════════════════════════")

	for lanjut == "y" || lanjut == "Y" {
		fmt.Println("\n▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬")
		fmt.Println("                     📊 MENU UTAMA")
		fmt.Println("▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬▬")
		fmt.Println("  ┌────────────────────────────────────────────┐")
		fmt.Println("  │  1.  Celsius    →   Fahrenheit   [°C → °F] │")
		fmt.Println("  │  2.  Celsius    →   Kelvin       [°C →  K] │")
		fmt.Println("  │  3.  Fahrenheit →   Celsius      [°F → °C] │")
		fmt.Println("  │  4.  Fahrenheit →   Kelvin       [°F →  K] │")
		fmt.Println("  │  5.  Kelvin     →   Celsius      [ K → °C] │")
		fmt.Println("  │  6.  Kelvin     →   Fahrenheit   [ K → °F] │")
		fmt.Println("  │  7.  📤 Keluar dari program                │")
		fmt.Println("  └────────────────────────────────────────────┘")

		fmt.Print("\n  → Pilih menu [1-7]: ")
		fmt.Scan(&pilihan)

		if pilihan == 7 {
			fmt.Println("\n══════════════════════════════════════════════════════════")
			fmt.Println("                 👋 TERIMA KASIH")
			fmt.Println("══════════════════════════════════════════════════════════")
			break
		}

		if pilihan < 1 || pilihan > 7 {
			fmt.Println("\n⚠️  Pilihan tidak valid! Silakan pilih antara 1-7")
			continue
		}

		fmt.Print("\n  ✏️  Masukkan nilai suhu: ")
		fmt.Scan(&nilai)

		if (pilihan == 5 || pilihan == 6) && nilai < 0 {
			fmt.Println("\n❌ ERROR: Nilai Kelvin tidak boleh negatif!")
			continue
		}

		fmt.Println("\n══════════════════════════════════════════════════════════")
		fmt.Println("                     📈 HASIL KONVERSI")
		fmt.Println("══════════════════════════════════════════════════════════")

		var hasil float64
		var dari, ke string

		if pilihan == 1 {
			hasil = (nilai * 9.0 / 5.0) + 32.0
			dari = "°C"
			ke = "°F"

		} else if pilihan == 2 {
			hasil = nilai + 273.15
			dari = "°C"
			ke = "K"

		} else if pilihan == 3 {
			hasil = (nilai - 32.0) * 5.0 / 9.0
			dari = "°F"
			ke = "°C"

		} else if pilihan == 4 {
			hasil = (nilai-32.0)*5.0/9.0 + 273.15
			dari = "°F"
			ke = "K"

		} else if pilihan == 5 {
			hasil = nilai - 273.15
			dari = "K"
			ke = "°C"

		} else if pilihan == 6 {
			hasil = (nilai-273.15)*9.0/5.0 + 32.0
			dari = "K"
			ke = "°F"
		}

		fmt.Printf("\n  ┌────────────────────────────────────────────┐\n")
		fmt.Printf("  │  📊  %10.2f %3s  =  %10.2f %3s  │\n", nilai, dari, hasil, ke)
		fmt.Printf("  └────────────────────────────────────────────┘\n")

		fmt.Println("\n  📝 INFORMASI:")

		var celsius float64
		if pilihan == 1 || pilihan == 2 {
			celsius = nilai
		} else if pilihan == 3 || pilihan == 4 {
			celsius = (nilai - 32.0) * 5.0 / 9.0
		} else if pilihan == 5 || pilihan == 6 {
			celsius = nilai - 273.15
		}

		if celsius <= -273.15 {
			fmt.Println("  • ❄️  Nol absolut (suhu terdingin)")
		} else if celsius < 0 {
			fmt.Println("  • 🌨️  Di bawah titik beku air")
		} else if celsius == 0 {
			fmt.Println("  • 🧊 Titik beku air")
		} else if celsius > 0 && celsius < 20 {
			fmt.Println("  • 🌬️  Suhu sejuk")
		} else if celsius >= 20 && celsius < 30 {
			fmt.Println("  • 😊 Suhu ruangan normal")
		} else if celsius >= 30 && celsius < 37 {
			fmt.Println("  • 🌞 Suhu hangat")
		} else if celsius >= 37 && celsius < 40 {
			fmt.Println("  • 🏥 Suhu tubuh manusia")
		} else if celsius == 100 {
			fmt.Println("  • ♨️  Titik didih air")
		} else if celsius > 100 {
			fmt.Println("  • 🔥 Suhu sangat panas")
		}

		if pilihan == 1 && nilai == -40 {
			fmt.Println("  • 💡 Fakta: -40°C = -40°F (titik temu)")
		} else if pilihan == 3 && nilai == -40 {
			fmt.Println("  • 💡 Fakta: -40°F = -40°C (titik temu)")
		} else if celsius == 37 {
			fmt.Println("  • 💡 Fakta: Suhu tubuh manusia normal")
		}

		fmt.Println("══════════════════════════════════════════════════════════")

		fmt.Print("\n  🔄 Konversi suhu lain? [y/n]: ")
		fmt.Scan(&lanjut)

		for lanjut != "y" && lanjut != "Y" && lanjut != "n" && lanjut != "N" {
			fmt.Print("  Masukkan 'y' untuk Ya atau 'n' untuk Tidak: ")
			fmt.Scan(&lanjut)
		}

		if lanjut == "y" || lanjut == "Y" {
			fmt.Println("\n" + stringsRepeat("─", 60))
		}
	}

	fmt.Println("\n")
	fmt.Println("╔══════════════════════════════════════════════════════════╗")
	fmt.Println("║                 PROGRAM TELAH SELESAI                    ║")
	fmt.Println("║                                                          ║")
	fmt.Println("║          Created with Azril, Chenghoo, Ilham             ║")
	fmt.Println("╚══════════════════════════════════════════════════════════╝")
}

// Fungsi untuk mengulang string (mengganti strings.Repeat)
func stringsRepeat(char string, count int) string {
	result := ""
	for i := 0; i < count; i++ {
		result += char
	}
	return result
}
