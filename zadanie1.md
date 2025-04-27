
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
![image](https://github.com/user-attachments/assets/49719049-5433-416e-a560-a99a8c408f18)
![image](https://github.com/user-attachments/assets/d8059bcc-84f3-4a56-a8b2-5bbdd13c233d)

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
![image](https://github.com/user-attachments/assets/b7fd3e7e-dd87-403f-b1d8-66374aaa852a)

## Wyświetlenie strony aplikacji

Po uruchomieniu kontenera, aplikacja była dostępna pod adresem `http://localhost:8000`. 
Umożliwiała wybór miasta z listy oraz wyświetlenie aktualnych danych pogodowych.
![image](https://github.com/user-attachments/assets/7dfd3989-20c2-430a-b966-238210ddda5f)

![image](https://github.com/user-attachments/assets/32f7a817-ead0-4518-ac40-fdd61ae5ca13)

## Odczytanie logów kontenera

Logi kontenera zostały wyświetlone za pomocą polecenia:
![image](https://github.com/user-attachments/assets/49c25883-924c-40b8-81fa-34b1ceb501ab)

```bash
docker logs pogoda
```

Log zawierał datę uruchomienia aplikacji, autora oraz numer portu.

## Sprawdzenie uruchomionych kontenerów

Lista uruchomionych kontenerów została uzyskana poleceniem:

```bash
docker ps
```

Wynik zawierał kontener `pogoda` działający na porcie 8000.
![image](https://github.com/user-attachments/assets/34e61f49-b1bf-4459-b5b0-2032d36edd2f)
![image](https://github.com/user-attachments/assets/23c5a82f-649c-45d4-995c-92373809423b)

## Sprawdzenie rozmiaru obrazu

Rozmiar utworzonego obrazu sprawdzono poleceniem:
![image](https://github.com/user-attachments/assets/8d85b8e5-52b2-43ff-85e9-2dbcb131e8b7)

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


