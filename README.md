# api-region-indonesia

untuk mengambil data region indonesia melalui api

## service app

PRODUCTION:

```
  ip : 192.168.88./24
  id : 139
  name : region-indonesia
  user-server : root
  pass-server : dc2013
  location app : /home/ubuntu/app/region-indonesia
```

DEVELOPMENT:

```
  ip : 192.168.88.
  id : 112
  name : dev-wec-checkout
  user-server : root
  pass-server : dc2013
  location app : ~/wec/checkout
```

### Build

```
	env GOOS=linux GOARCH=amd64 go build -o region-indonesia.linux region.go

	env GOOS=darwin GOARCH=amd64 go build -o region-indonesia.linux region.go

	go build -o region-indonesia.linux region.go
```

## runnig

> nohup ./region-indonesia.linux > region-indonesia.out &

## upload server go-build to server redis-region-indonesia

PRODUCTION

```
  scp region-indonesia.linux ubuntu@redis-region-indonesia:~/app/region-indonesia/region-indonesia.linux
  scp region-indonesia.linux ubuntu@192.168.88.220:~/app/region-indonesia/region-indonesia.linux
```

DEVELOPMENT

```
  scp region-indonesia.linux ubuntu@192.168.88.50:~/app/region-indonesia/region-indonesia.linux
```

## get provinsi

```
  {"topic":"region-indonesia","proc":"get","data":"{\"id\":\"provinces\"}"}
```

## get non-provinsi

```
  {"topic":"region-indonesia","proc":"get","data":"{\"id\":\"72\"}"}
```

## Provinces

> url : 192.168.88.108/region-indonesia/provinces

> metode : GET

Respon:

```
  [
    {
		"id": 11,
		"name": "ACEH"
	},
	{
		"id": 12,
		"name": "SUMATERA UTARA"
	},
	{
		"id": 13,
		"name": "SUMATERA BARAT"
	},
  ...
	{
		"id": 94,
		"name": "PAPUA"
	}
  ]
```

## Kota / Kecamatan / Kelurahan

> url : 192.168.88.108/region-indonesia/[id]

> metode : GET

Respon:

```
  [
    {
      "id": 1101,
      "provinceId": 11,
      "provinceName": "ACEH",
      "name": "KABUPATEN SIMEULUE"
    },
    {
      "id": 1102,
      "provinceId": 11,
      "provinceName": "ACEH",
      "name": "KABUPATEN ACEH SINGKIL"
    },
    {
      "id": 1103,
      "provinceId": 11,
      "provinceName": "ACEH",
      "name": "KABUPATEN ACEH SELATAN"
    },
    ...
    {
      "id": 1175,
      "provinceId": 11,
      "provinceName": "ACEH",
      "name": "KOTA SUBULUSSALAM"
    }
  ]
```
