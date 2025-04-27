
# Zadanie 1 - Część dodatkowa (Sprawozdanie)

## Wstęp

Celem części dodatkowej było:
- Przygotowanie obrazu kontenera dla platform `linux/amd64` oraz `linux/arm64`,
- Wykorzystanie buildera opartego o sterownik `docker-container`,
- Weryfikacja poprawności manifestu obrazów,
- Przeprowadzenie analizy bezpieczeństwa obrazu przy użyciu narzędzia Trivy.

## Utworzenie własnego buildera

Aby stworzyć dedykowany builder `multi-builder`, wykorzystano polecenie:

```bash
docker buildx create --name multi-builder --use
docker buildx inspect --bootstrap
```
![image](https://github.com/user-attachments/assets/97d1dccc-c7b1-4dbe-85d0-2923b9812755)

Builder został poprawnie utworzony i uruchomiony.

## Budowa obrazu dla wielu platform

Następnie wykonano budowę obrazu dla dwóch platform (`linux/amd64`, `linux/arm64`) oraz wypchnięcie obrazu do DockerHub:

```bash
docker buildx build --platform linux/amd64,linux/arm64 -t kamilz12/weather-go:latest --push .
```

Opis:
- `--platform linux/amd64,linux/arm64` - budowa wieloplatformowa,
- `-t kamilz12/weather-go:latest` - oznaczenie obrazu,
- `--push` - automatyczne wypchnięcie obrazu do DockerHub.
![image](https://github.com/user-attachments/assets/68d793fa-3010-4666-bb10-9e984584c388)

Obraz został poprawnie zbudowany i opublikowany.

## Weryfikacja manifestu obrazu

Poprawność wieloplatformowego obrazu została zweryfikowana poleceniem:

```bash
docker buildx imagetools inspect kamilz12/weather-go:latest
```
![image](https://github.com/user-attachments/assets/29d94401-f9bd-4e0a-871a-40c11ff89416)

Wynik pokazał manifesty dla platform `linux/amd64` oraz `linux/arm64`, co potwierdza poprawność budowy.

## Analiza podatności (CVE)

Bezpieczeństwo obrazu zostało sprawdzone przy pomocy narzędzia Trivy:

```bash
trivy image kamilz12/weather-go:latest
```
![image](https://github.com/user-attachments/assets/9d25a418-1eb6-40d3-9ea2-f2a2447938c5)

Wynik:
- Brak zagrożeń o poziomie CRITICAL lub HIGH.

Obraz jest bezpieczny i spełnia wymagania dotyczące bezpieczeństwa.

## Podsumowanie

W ramach części dodatkowej wykonano:
- Budowę wieloplatformowego obrazu Docker,
- Publikację obrazu do DockerHub,
- Weryfikację obecności manifestów dla wymaganych platform,
- Skuteczną analizę bezpieczeństwa (bez krytycznych podatności).


