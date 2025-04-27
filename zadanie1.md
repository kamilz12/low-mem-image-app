
# Zadanie 1 - Sprawozdanie

## Wstęp
Celem zadania było stworzenie aplikacji serwerowej działającej w kontenerze Docker, umożliwiającej wybór miasta
i wyświetlenie aktualnej pogody. Aplikacja została napisana w języku Go. W ramach zadania wykonano również
optymalizację obrazu kontenera oraz przygotowano plik `Dockerfile`.

## Budowanie obrazu Docker

Aby zbudować obraz kontenera dla aplikacji, wykorzystano polecenie:

```bash
docker buildx build --platform linux/amd64 --no-cache -t kziolkowski/weather-go:latest --load .
```

Opis:
- `--platform linux/amd64` - budowanie obrazu pod platformę amd64,
- `--no-cache` - budowanie obrazu od podstaw, bez użycia cache,
- `-t kziolkowski/weather-go:latest` - nadanie tagu dla obrazu,
- `--load` - załadowanie obrazu do lokalnego cache Dockera.

Po wykonaniu polecenia obraz został poprawnie zbudowany i załadowany.

## Uruchomienie kontenera

Kontener z aplikacją został uruchomiony poleceniem:

```bash
docker run -d --name pogoda -p 8000:8000 -e OWM_KEY=<API_KEY> kziolkowski/weather-go:latest
```

Opis:
- `-d` - uruchomienie w tle (detached mode),
- `--name pogoda` - nadanie nazwy kontenerowi,
- `-p 8000:8000` - mapowanie portu lokalnego na port kontenera,
- `-e OWM_KEY=<API_KEY>` - ustawienie zmiennej środowiskowej z kluczem API do pobierania pogody.

## Wyświetlenie strony aplikacji

Po uruchomieniu kontenera, aplikacja była dostępna pod adresem `http://localhost:8000`. 
Umożliwiała wybór miasta z listy oraz wyświetlenie aktualnych danych pogodowych.


## Odczytanie logów kontenera

Logi kontenera zostały wyświetlone za pomocą polecenia:

```bash
docker logs pogoda
```

Przykładowy log:
```
Start 2025-04-27T12:35:32Z | Autor: Kamil Ziolkowski | Port: 8000
```

Log zawierał datę uruchomienia aplikacji, autora oraz numer portu.

## Sprawdzenie uruchomionych kontenerów

Lista uruchomionych kontenerów została uzyskana poleceniem:

```bash
docker ps
```

Wynik zawierał kontener `pogoda` działający na porcie 8000.

## Sprawdzenie rozmiaru obrazu

Rozmiar utworzonego obrazu sprawdzono poleceniem:

```bash
docker images kziolkowski/weather-go:latest --format "{{.Size}}"
```

Uzyskany rozmiar obrazu:
```
2.53MB
```

## Podsumowanie

W ramach zadania:
- Stworzono aplikację serwerową w języku Go,
- Zbudowano zoptymalizowany obraz Dockera o rozmiarze 2.53MB,
- Uruchomiono aplikację w kontenerze,
- Zweryfikowano poprawne działanie aplikacji poprzez stronę WWW oraz logi systemowe.

Linki do repozytoriów:

