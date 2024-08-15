# go-keycloak-sample

### Things todo list

1. Clone this repository: `git clone https://github.com/hendisantika/go-keycloak-sample.git`
2. Navigate to the folder: `cd go-keycloak-sample`
3. Run Keycloak Container: `docker compose up`
4. Run the go app: `go run main.go`
5. Import reals that I already provided which is `PowerRanger.json` file into your keycloak

### Console Log

#### Generate Token

```shell
curl --location 'http://localhost:8080/realms/PowerRanger/protocol/openid-connect/token' \
--header 'Content-Type: application/x-www-form-urlencoded' \
--data-urlencode 'client_id=my-go-service' \
--data-urlencode 'client_secret=ASAjLseDVyuRvHjtTf26gvR1USoT0uO4' \
--data-urlencode 'grant_type=client_credentials' \
--data-urlencode 'scope=openid'
```

Body Response:

```json
{
  "access_token": "eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJMVFFGV1phc0JKNDdSSTV0ZVc1VEVadGR2Yi1oRF9LOGN1LXVqU0ptc2tBIn0.eyJleHAiOjE3MjM3MTIwOTAsImlhdCI6MTcyMzcxMTc5MCwianRpIjoiZGEwN2U5NDktZTI5OC00NjcxLWEwMGItYTIyYTlkZjA4NDU5IiwiaXNzIjoiaHR0cDovL2xvY2FsaG9zdDo4MDgwL3JlYWxtcy9Qb3dlclJhbmdlciIsImF1ZCI6ImFjY291bnQiLCJzdWIiOiI2ZWQwMjVhZC0wZGUyLTQ0MGQtYWIyZi0zMWMxOTA1NjcwZjMiLCJ0eXAiOiJCZWFyZXIiLCJhenAiOiJteS1nby1zZXJ2aWNlIiwiYWNyIjoiMSIsImFsbG93ZWQtb3JpZ2lucyI6WyIvKiJdLCJyZWFsbV9hY2Nlc3MiOnsicm9sZXMiOlsib2ZmbGluZV9hY2Nlc3MiLCJ1bWFfYXV0aG9yaXphdGlvbiIsImRlZmF1bHQtcm9sZXMtcG93ZXJyYW5nZXIiXX0sInJlc291cmNlX2FjY2VzcyI6eyJhY2NvdW50Ijp7InJvbGVzIjpbIm1hbmFnZS1hY2NvdW50IiwibWFuYWdlLWFjY291bnQtbGlua3MiLCJ2aWV3LXByb2ZpbGUiXX19LCJzY29wZSI6Im9wZW5pZCBwcm9maWxlIGVtYWlsIiwiZW1haWxfdmVyaWZpZWQiOmZhbHNlLCJjbGllbnRIb3N0IjoiMTkyLjE2OC42NS4xIiwicHJlZmVycmVkX3VzZXJuYW1lIjoic2VydmljZS1hY2NvdW50LW15LWdvLXNlcnZpY2UiLCJjbGllbnRBZGRyZXNzIjoiMTkyLjE2OC42NS4xIiwiY2xpZW50X2lkIjoibXktZ28tc2VydmljZSJ9.eiPOW-iTB8LMaJdnUcZxWIIGk0V9s8nC_ZxMsXCZxWFLNDWwh0UAd0kz-9oLzugy6E9-41r8-ZC0rcSZ8WTOzfSx_-TnutcJkTjuybKAWGOqlFobN9J76rM8zym5esVsI5nR3c0moLUj3nMTjPd5a8Mj72OP8gOs60FTtgEcEdJ_UZ-wVN8ZojwagucwT9t7eq3a0D2LK1BJdUfFTuIgGZwy9G2tdN3nqghjUPsqLcFYCLPodE04TsRvCAyGxeavChpEt_K-9Yar9g4WmXxeX-SWVUBQVJp06HgaW2fETyCuACh3e-0k8zqJasDpFPP6Ga_1dJBNkShnEcfsmco7qA",
  "expires_in": 300,
  "refresh_expires_in": 0,
  "token_type": "Bearer",
  "id_token": "eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJMVFFGV1phc0JKNDdSSTV0ZVc1VEVadGR2Yi1oRF9LOGN1LXVqU0ptc2tBIn0.eyJleHAiOjE3MjM3MTIwOTAsImlhdCI6MTcyMzcxMTc5MCwianRpIjoiNTI2NDhmODgtYTFjMC00MzY4LWI3ZGUtYjY2MDUyMWMyYzRkIiwiaXNzIjoiaHR0cDovL2xvY2FsaG9zdDo4MDgwL3JlYWxtcy9Qb3dlclJhbmdlciIsImF1ZCI6Im15LWdvLXNlcnZpY2UiLCJzdWIiOiI2ZWQwMjVhZC0wZGUyLTQ0MGQtYWIyZi0zMWMxOTA1NjcwZjMiLCJ0eXAiOiJJRCIsImF6cCI6Im15LWdvLXNlcnZpY2UiLCJhdF9oYXNoIjoiS04tNGVTUlBBcmd6OUROYmtxUW9FdyIsImFjciI6IjEiLCJlbWFpbF92ZXJpZmllZCI6ZmFsc2UsImNsaWVudEhvc3QiOiIxOTIuMTY4LjY1LjEiLCJwcmVmZXJyZWRfdXNlcm5hbWUiOiJzZXJ2aWNlLWFjY291bnQtbXktZ28tc2VydmljZSIsImNsaWVudEFkZHJlc3MiOiIxOTIuMTY4LjY1LjEiLCJjbGllbnRfaWQiOiJteS1nby1zZXJ2aWNlIn0.HRis2J8YFGlcmV_1b9PPmXNgheOfrOq_3cjZ6Wa7l5GkQHuAtAFsOb7pM7bGUKL-Av_mARuwHc4VZ5e3YI-5-kA4gGPgX2y34zF5j9RSuGcJ7WG9f_7RvmwF61YR_zao9jnzdTbtVArinyq3MWnyAsVffNhU6wEkLlyFi23X93yLzftYx7URtWICx-K3YazuSLqQw3A6n6Kz2cLvazRyIEDp1oaXIFixQah2mv3SR-N3CRhwSyoUzd8dF3zGhxsGMQ9OAMQiVuWqap-4I5DcuE9D_E1yCjiPYvVe-0HY-v84w18n1_Ucr0MmluiaWSykbUMDYLPMc_uD6d4_JDfX8Q",
  "not-before-policy": 0,
  "scope": "openid profile email"
}
```

#### Add new event

```shell
curl --location 'http://localhost:8081/events' \
--header 'Content-Type: application/json' \
--header 'Authorization: Bearer eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJMVFFGV1phc0JKNDdSSTV0ZVc1VEVadGR2Yi1oRF9LOGN1LXVqU0ptc2tBIn0.eyJleHAiOjE3MjM3MDcwNzksImlhdCI6MTcyMzcwNjc3OSwianRpIjoiOGY1N2M4MTktZTBlYy00YTVlLWExNjMtNzY3MWM4ZTRlMzA1IiwiaXNzIjoiaHR0cDovL2xvY2FsaG9zdDo4MDgwL3JlYWxtcy9Qb3dlclJhbmdlciIsImF1ZCI6ImFjY291bnQiLCJzdWIiOiI2ZWQwMjVhZC0wZGUyLTQ0MGQtYWIyZi0zMWMxOTA1NjcwZjMiLCJ0eXAiOiJCZWFyZXIiLCJhenAiOiJteS1nby1zZXJ2aWNlIiwiYWNyIjoiMSIsImFsbG93ZWQtb3JpZ2lucyI6WyIvKiJdLCJyZWFsbV9hY2Nlc3MiOnsicm9sZXMiOlsib2ZmbGluZV9hY2Nlc3MiLCJ1bWFfYXV0aG9yaXphdGlvbiIsImRlZmF1bHQtcm9sZXMtcG93ZXJyYW5nZXIiXX0sInJlc291cmNlX2FjY2VzcyI6eyJhY2NvdW50Ijp7InJvbGVzIjpbIm1hbmFnZS1hY2NvdW50IiwibWFuYWdlLWFjY291bnQtbGlua3MiLCJ2aWV3LXByb2ZpbGUiXX19LCJzY29wZSI6Im9wZW5pZCBwcm9maWxlIGVtYWlsIiwiZW1haWxfdmVyaWZpZWQiOmZhbHNlLCJjbGllbnRIb3N0IjoiMTkyLjE2OC42NS4xIiwicHJlZmVycmVkX3VzZXJuYW1lIjoic2VydmljZS1hY2NvdW50LW15LWdvLXNlcnZpY2UiLCJjbGllbnRBZGRyZXNzIjoiMTkyLjE2OC42NS4xIiwiY2xpZW50X2lkIjoibXktZ28tc2VydmljZSJ9.vvkTRwbVwwnSPkd3-Sd_o04OL_aoCcqQVwWQaiWj4E46mLaaQDjCOhRufeN-rXM3te_LFSRhbCP4KfwLK_6tEPEMb9DR-gNyKDI4sibHl55pmIvVtbCCYaZzLslY4lnedigcvXQUnBkTV6-EUtjXKluBkMmBJsOPoUNHzlxU2xJLg7xdqjTIy_7aUNxAdefn6d0jGvb9hv51qrARcZZkePA84AKc7J-_GRwz1Ji4KJ_45ZQtxC-jcsPWaBIu94P5UOjgUO2MeguyVnCU99Wuq8m8Z7n7NPJHVeNTwrGDCq55r3dLt86hPbWmzMPMNO1Gc3QRQA-EsLQJ_23e80UTjg' \
--data '{
    "Title": "Test 1",
    "Description": "Description"
}'
```

Body Response:

```shell
{
    "CreatedAt": "2024-08-15T15:55:27.955+07:00",
    "UpdatedAt": "2024-08-15T15:55:27.955+07:00",
    "DeletedAt": null,
    "ID": 7,
    "Title": "Test 1",
    "Description": "Description"
}
```

#### Get All Events

```shell
curl --location 'http://localhost:8081/events' \
--header 'Authorization: Bearer eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJMVFFGV1phc0JKNDdSSTV0ZVc1VEVadGR2Yi1oRF9LOGN1LXVqU0ptc2tBIn0.eyJleHAiOjE3MjM3MDcwNzksImlhdCI6MTcyMzcwNjc3OSwianRpIjoiOGY1N2M4MTktZTBlYy00YTVlLWExNjMtNzY3MWM4ZTRlMzA1IiwiaXNzIjoiaHR0cDovL2xvY2FsaG9zdDo4MDgwL3JlYWxtcy9Qb3dlclJhbmdlciIsImF1ZCI6ImFjY291bnQiLCJzdWIiOiI2ZWQwMjVhZC0wZGUyLTQ0MGQtYWIyZi0zMWMxOTA1NjcwZjMiLCJ0eXAiOiJCZWFyZXIiLCJhenAiOiJteS1nby1zZXJ2aWNlIiwiYWNyIjoiMSIsImFsbG93ZWQtb3JpZ2lucyI6WyIvKiJdLCJyZWFsbV9hY2Nlc3MiOnsicm9sZXMiOlsib2ZmbGluZV9hY2Nlc3MiLCJ1bWFfYXV0aG9yaXphdGlvbiIsImRlZmF1bHQtcm9sZXMtcG93ZXJyYW5nZXIiXX0sInJlc291cmNlX2FjY2VzcyI6eyJhY2NvdW50Ijp7InJvbGVzIjpbIm1hbmFnZS1hY2NvdW50IiwibWFuYWdlLWFjY291bnQtbGlua3MiLCJ2aWV3LXByb2ZpbGUiXX19LCJzY29wZSI6Im9wZW5pZCBwcm9maWxlIGVtYWlsIiwiZW1haWxfdmVyaWZpZWQiOmZhbHNlLCJjbGllbnRIb3N0IjoiMTkyLjE2OC42NS4xIiwicHJlZmVycmVkX3VzZXJuYW1lIjoic2VydmljZS1hY2NvdW50LW15LWdvLXNlcnZpY2UiLCJjbGllbnRBZGRyZXNzIjoiMTkyLjE2OC42NS4xIiwiY2xpZW50X2lkIjoibXktZ28tc2VydmljZSJ9.vvkTRwbVwwnSPkd3-Sd_o04OL_aoCcqQVwWQaiWj4E46mLaaQDjCOhRufeN-rXM3te_LFSRhbCP4KfwLK_6tEPEMb9DR-gNyKDI4sibHl55pmIvVtbCCYaZzLslY4lnedigcvXQUnBkTV6-EUtjXKluBkMmBJsOPoUNHzlxU2xJLg7xdqjTIy_7aUNxAdefn6d0jGvb9hv51qrARcZZkePA84AKc7J-_GRwz1Ji4KJ_45ZQtxC-jcsPWaBIu94P5UOjgUO2MeguyVnCU99Wuq8m8Z7n7NPJHVeNTwrGDCq55r3dLt86hPbWmzMPMNO1Gc3QRQA-EsLQJ_23e80UTjg'
```

Body Response:

```json
[
  {
    "CreatedAt": "2024-08-15T14:26:37.795+07:00",
    "UpdatedAt": "2024-08-15T14:26:37.795+07:00",
    "DeletedAt": null,
    "ID": 1,
    "Title": "Test 1",
    "Description": "Description"
  },
  {
    "CreatedAt": "2024-08-15T14:26:41.244+07:00",
    "UpdatedAt": "2024-08-15T14:26:41.244+07:00",
    "DeletedAt": null,
    "ID": 2,
    "Title": "Test 1",
    "Description": "Description"
  },
  {
    "CreatedAt": "2024-08-15T14:26:41.877+07:00",
    "UpdatedAt": "2024-08-15T14:26:41.877+07:00",
    "DeletedAt": null,
    "ID": 3,
    "Title": "Test 1",
    "Description": "Description"
  },
  {
    "CreatedAt": "2024-08-15T14:26:42.515+07:00",
    "UpdatedAt": "2024-08-15T14:26:42.515+07:00",
    "DeletedAt": null,
    "ID": 4,
    "Title": "Test 1",
    "Description": "Description"
  },
  {
    "CreatedAt": "2024-08-15T14:26:43.001+07:00",
    "UpdatedAt": "2024-08-15T14:26:43.001+07:00",
    "DeletedAt": null,
    "ID": 5,
    "Title": "Test 1",
    "Description": "Description"
  },
  {
    "CreatedAt": "2024-08-15T14:26:43.516+07:00",
    "UpdatedAt": "2024-08-15T14:26:43.516+07:00",
    "DeletedAt": null,
    "ID": 6,
    "Title": "Test 1",
    "Description": "Description"
  }
]
```

#### Get Event By ID

```shell
curl --location 'http://localhost:8081/events/1' \
--header 'Authorization: Bearer eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJMVFFGV1phc0JKNDdSSTV0ZVc1VEVadGR2Yi1oRF9LOGN1LXVqU0ptc2tBIn0.eyJleHAiOjE3MjM3MTI0MjIsImlhdCI6MTcyMzcxMjEyMiwianRpIjoiOGIxNjVhZWYtM2QxMC00YjEyLWI3OTYtMzJmMzg4NzBkMmQ0IiwiaXNzIjoiaHR0cDovL2xvY2FsaG9zdDo4MDgwL3JlYWxtcy9Qb3dlclJhbmdlciIsImF1ZCI6ImFjY291bnQiLCJzdWIiOiI2ZWQwMjVhZC0wZGUyLTQ0MGQtYWIyZi0zMWMxOTA1NjcwZjMiLCJ0eXAiOiJCZWFyZXIiLCJhenAiOiJteS1nby1zZXJ2aWNlIiwiYWNyIjoiMSIsImFsbG93ZWQtb3JpZ2lucyI6WyIvKiJdLCJyZWFsbV9hY2Nlc3MiOnsicm9sZXMiOlsib2ZmbGluZV9hY2Nlc3MiLCJ1bWFfYXV0aG9yaXphdGlvbiIsImRlZmF1bHQtcm9sZXMtcG93ZXJyYW5nZXIiXX0sInJlc291cmNlX2FjY2VzcyI6eyJhY2NvdW50Ijp7InJvbGVzIjpbIm1hbmFnZS1hY2NvdW50IiwibWFuYWdlLWFjY291bnQtbGlua3MiLCJ2aWV3LXByb2ZpbGUiXX19LCJzY29wZSI6Im9wZW5pZCBwcm9maWxlIGVtYWlsIiwiZW1haWxfdmVyaWZpZWQiOmZhbHNlLCJjbGllbnRIb3N0IjoiMTkyLjE2OC42NS4xIiwicHJlZmVycmVkX3VzZXJuYW1lIjoic2VydmljZS1hY2NvdW50LW15LWdvLXNlcnZpY2UiLCJjbGllbnRBZGRyZXNzIjoiMTkyLjE2OC42NS4xIiwiY2xpZW50X2lkIjoibXktZ28tc2VydmljZSJ9.CuzQNQ95vNEqFVO9PVH_Gfgref6CIKRexEQXLIqs-WqFvh4QHSR7rpOvK7gQaCaAGHTzMusHhn-wiyDQo-bka-0qehidE0U-fnfRJB_JlU3d7Y52a6rq3T3LiGMzJZ3As1ifC26lfMxWkMS5hyDM4NzM8l_85H2UKESyzaYZyqMqMzcn8ll9vvmCCQOyPTY8WqYQbZDNN-Fb7tfRk7V_s9L27IasNnjy-CnJauQ32m2WrZxTc43QgUzkqLfq-YdfGqueOTqHs8--5dBiTOsTyRf3UZFndXABvxWJEiFP-4_QI2c06w8qNI3WdveZKzDXY85mAaazsOSX_g-G9fOGfA'
```

Body Response:

```json
{
  "CreatedAt": "2024-08-15T14:26:37.795+07:00",
  "UpdatedAt": "2024-08-15T14:26:37.795+07:00",
  "DeletedAt": null,
  "ID": 1,
  "Title": "Test 1",
  "Description": "Description"
}
```