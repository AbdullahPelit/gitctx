# gitctx - GitHub/GitLab Hesap Yönetim Aracı

`gitctx`, birden fazla GitHub veya GitLab hesabını kolayca yönetmenizi sağlayan bir araçtır. SSH anahtarlarını dinamik olarak oluşturabilir, hesaplar arasında geçiş yapabilir ve farklı hesaplar için ayrı SSH anahtarları kullanabilirsiniz.

## Özellikler

- SSH anahtarı oluşturma ve otomatik olarak ekrana yazdırma
- Hesap ekleme, silme ve geçiş yapma
- Konfigürasyon dosyasına kayıtlı hesapları listeleme

## Kurulum

### 1. Depoyu Klonlayın

Öncelikle, `gitctx` aracını indirmeniz gerekmektedir. Bunun için aşağıdaki komutu kullanarak depoyu klonlayın:

```bash
git clone https://github.com/yourusername/gitctx.git
cd gitctx
```

### 2. Script'i Çalıştırılabilir Hale Getirin

Script'in çalıştırılabilir olması için gerekli izinleri verin:

```bash
chmod +x gitctx.sh
```

### 3. Script'i PATH'e Ekleyin

Script'i sistem PATH'ine ekleyerek her yerden kullanılabilir hale getirin. Bunun için script'i ~/bin/ dizinine taşıyabilir ve PATH'e ekleyebilirsiniz:

```bash
mv gitctx.sh ~/bin/gitctx
```

Eğer ~/bin/ dizini yoksa, oluşturabilirsiniz:

```bash
mkdir -p ~/bin
mv gitctx.sh ~/bin/gitctx
```

Ardından, bu dizini PATH'e eklemek için aşağıdaki satırı ~/.bashrc veya ~/.zshrc dosyanıza ekleyin:

```bash
export PATH=$PATH:~/bin
source ~/.bashrc
```

## Kullanım

### Hesap Ekleme

Yeni bir hesap eklemek için:

```bash
gitctx add
```

Bu komut, size yeni bir SSH anahtarı oluşturma veya mevcut bir anahtarı kullanma seçeneği sunar. Yeni bir SSH anahtarı oluşturursanız, bu anahtarı Git provider'ınıza (GitHub, GitLab, vb.) eklemeniz gerekir. SSH anahtarı ekrana yazdırılacaktır.

### Hesap Silme

Mevcut bir hesabı silmek için:

```bash
gitctx remove
```

### Hesaplar Arasında Geçiş Yapma

Eklediğiniz hesaplar arasında geçiş yapmak için:

```bash
gitctx switch <account_name>
```

### Hesapları Listeleme

Kayıtlı tüm hesapları listelemek için:

```bash
gitctx list
```

# Katkıda Bulunma
Katkıda bulunmak isterseniz, lütfen bir fork yapın ve pull request gönderin. Her türlü katkıya açığız!



