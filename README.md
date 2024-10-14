# gitctx - GitHub/GitLab Hesap Yönetim Aracı

`gitctx`, birden fazla GitHub veya GitLab hesabını kolayca yönetmenizi sağlayan bir araçtır. SSH anahtarlarını dinamik olarak oluşturabilir, hesaplar arasında geçiş yapabilir ve farklı hesaplar için ayrı SSH anahtarları kullanabilirsiniz.

## Özellikler

- Hesaplar için ayrı ayrı SSH anahtarı oluşturma ve yönetme
- Hesaplar arasında hızlı geçiş
- Kayıtlı hesapları listeleme
- Hesap ekleme ve silme

## Kurulum

### 1. Depoyu Klonlayın

Öncelikle, `gitctx` aracını indirmeniz gerekmektedir. Bunun için aşağıdaki komutu kullanarak depoyu klonlayın:

```bash
git clone https://github.com/AbdullahPelit/gitctx.git
cd gitctx
```

### 2. Bağımlılıkları Kurun

Proje cobra paketini kullanmaktadır. Bağımlılıkları yüklemek için terminalde şu komutu çalıştırın:

```bash
go mod tidy
```

### 3. Projeyi Derleyin (Build)

Projenin çalışabilir bir versiyonunu oluşturmak için aşağıdaki komutu kullanarak projeyi derleyebilirsiniz:

```bash
go build -o gitctx
```
### 4. gitctx'i PATH'e Ekleyin

gitctx komutunu her yerden kullanabilmek için çalıştırılabilir dosyayı sistemin PATH'ine ekleyin. Bunun için:

```bash
mv gitctx ~/bin/
```

Ardından, bu dizini PATH'e eklemek için ~/.bashrc, ~/.zshrc veya terminalde kullanılan kabuğa uygun dosyaya şu satırı ekleyin:

```bash
export PATH=$PATH:~/bin
source ~/.bashrc  # veya ~/.zshrc kullanıyorsanız bu komutu çalıştırın
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
gitctx remove <account_name>
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



