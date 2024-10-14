# Katkı Kılavuzu

Projeye katkıda bulunduğunuz için teşekkürler! İşte katkı yapmanın yolları ve süreçleri hakkında bazı yönergeler:

## Nasıl Katkı Yapabilirsiniz?

1. Bir sorun (issue) rapor edebilir veya var olan bir sorun üzerinde çalışabilirsiniz.
2. Yeni özellikler ekleyebilir veya mevcut olanları geliştirebilirsiniz.
3. Hataları düzeltebilirsiniz.

## Kodlama Standartları

- Kod yazarken `prettier` kullanarak kodunuzu formatlayın.
- Değişiklik yapmadan önce lütfen projenin kod stili ile uyumlu olduğundan emin olun.

## Pull Request Süreci

1. Geliştirmelerinizi kendi dalınızda (branch) yapın: `git checkout -b yeni-ozellik-adi`
2. Değişiklikleri yapın ve commit'leyin: `git commit -m 'Özellik açıklaması'`
3. Deponuza (repo) push'layın: `git push origin yeni-ozellik-adi`
4. GitHub üzerinde bir pull request oluşturun ve değişikliklerinizi açıklayın.

## Testler

Lütfen katkıda bulunmadan önce tüm testlerin geçtiğinden emin olun. Yeni özellikler ekliyorsanız, yeni testler de eklemeyi unutmayın. Test komutunu çalıştırmak için:
```bash
npm test
