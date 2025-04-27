
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

Obraz został poprawnie zbudowany i opublikowany.

## Weryfikacja manifestu obrazu

Poprawność wieloplatformowego obrazu została zweryfikowana poleceniem:

```bash
docker buildx imagetools inspect kamilz12/weather-go:latest
```

Wynik pokazał manifesty dla platform `linux/amd64` oraz `linux/arm64`, co potwierdza poprawność budowy.

## Analiza podatności (CVE)

Bezpieczeństwo obrazu zostało sprawdzone przy pomocy narzędzia Trivy:

```bash
trivy image kamilz12/weather-go:latest
```

Wynik:
- Brak zagrożeń o poziomie CRITICAL lub HIGH.

Obraz jest bezpieczny i spełnia wymagania dotyczące bezpieczeństwa.

## Podsumowanie

W ramach części dodatkowej wykonano:
- Budowę wieloplatformowego obrazu Docker,
- Publikację obrazu do DockerHub,
- Weryfikację obecności manifestów dla wymaganych platform,
- Skuteczną analizę bezpieczeństwa (bez krytycznych podatności).

Linki do repozytoriów:

