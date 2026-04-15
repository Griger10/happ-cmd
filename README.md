# happcmd

CLI-инструмент для генерации ссылок импорта профилей маршрутизации для приложения [Happ](https://happ.su).

## Установка

### Скачать бинарник

Скачайте последний релиз со страницы [Releases](https://github.com/Griger10/happcmd/releases).

### Собрать из исходников

```bash
git clone https://github.com/Griger10/happcmd
cd happcmd
go build -o happcmd .
```

## Использование

### Интерактивный режим

Запустите без аргументов — откроется интерактивное меню:

```bash
happcmd
```

### CLI режим

```bash
# Дефолтный профиль для России
happcmd generate

# С именем профиля
happcmd generate -n "Мой профиль"

# Автоактивация при импорте
happcmd generate -m onadd

# Добавить сайты в прямой доступ
happcmd generate --add-direct-site "domain:github.com" --add-direct-site "domain:notion.so"

# Добавить сайты в блокировку
happcmd generate --add-block-site "geosite:gambling"

# Добавить IP в прямой доступ
happcmd generate --add-direct-ip "1.2.3.4/32"

# Комбинированный пример
happcmd generate -n "Work" -m onadd \
  --add-direct-site "domain:github.com" \
  --add-block-site "geosite:gambling"
```

## Флаги

| Флаг | Короткий | По умолчанию | Описание |
|------|----------|--------------|----------|
| `--name` | `-n` | `DefaultProfile` | Название профиля в Happ |
| `--mode` | `-m` | `add` | Режим импорта: `add` или `onadd` |
| `--add-direct-site` | — | — | Добавить домен в прямой доступ |
| `--add-block-site` | — | — | Добавить домен в блокировку |
| `--add-direct-ip` | — | — | Добавить IP/CIDR в прямой доступ |

## Режимы импорта

| Режим | Описание |
|-------|----------|
| `add` | Добавляет профиль в список. Первый добавленный становится активным после загрузки геофайлов |
| `onadd` | Добавляет и сразу активирует профиль |

## Что включено в дефолтный профиль

**Прямой доступ (без туннеля):**
- Российские домены и IP (`geosite:ru`, `geoip:ru`)
- ВКонтакте, Яндекс, Mail.ru и связанные CDN
- Государственные сайты (`geosite:category-gov-ru`)
- Локальные сети

**Блокировка:**
- Рекламные сети (`geosite:category-ads-all`)

**DNS:**
- Remote: Google DoH (`https://dns.google/dns-query`)
- Domestic: Яндекс DoU

## Требования

- Go 1.21+
- Приложение [Happ](https://happ.su) на iOS

## Лицензия

MIT