package piscine

import "fmt"

func QuadA(x, y int) {
	if x > 0 && y > 0 { // Pozitiflik negatiflik kontrolu
		for a := 0; a < x; a++ {
			if a == 0 || a == x-1 { // Başlangıç - Son değeri-karakteri belirlemek için yapıyoruz
				fmt.Print("o") // Eğer baş ve son karaktere gelirse 'o' karakterini eklemesini istedik
			} else {
				fmt.Print("-") // Eğer baş ve son karakter hariç bir karaktere denk gelirse yani ardaki karakterler '-' karakterini eklemesini istedik
			}
		}
		fmt.Println() // Bu komutla da  yaptığı işlemden sonra alt satıra  inmesini istedik

		for b := 0; b < y-2; b++ { // Bu işlemde de dikeylik kontrolü yaptık 'y-2'ile başa ve sona koyduğumuz 'o' harflerini eksilterek arada kalan kadar satırlara '|' karakterini koymasını istedik
			fmt.Print("|")             // BUrda direk başlangıç kısmına işareti koydurduk satırda
			for c := 0; c < x-2; c++ { // Bu işlemle  baş ve son satır dışındaki kısımlarda boşluk yazıdrmayı ve sonuncu kısıma '|' ifadesini bastırmayı kontrol ediyoruz
				fmt.Print(" ")
			}
			if x > 1 {
				fmt.Print("|") // kısımda satırdaki '|' işaretini bastırılmasını kontrol ediyoruz
			}
			fmt.Println() // Alta geçmesi durumunu kontrol ediyoruz
		}
		if y > 1 { // Burdaki işlemde satır veya sütüunda tek işlem yapmak istediğimiz zaman ikinci bir şaretin gelmesini engelliyoruz ve onu kontrol ediyoruz
			for a := 0; a < x; a++ { // Burda da onun yazıdırlması işlemi kısmı var
				if a == 0 || a == x-1 {
					fmt.Print("o")
				} else {
					fmt.Print("-")
				}
			}
		}
	}
}
